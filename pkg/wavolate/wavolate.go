// Package wavolate processes a wav file into a
// graphable dataset of magnitude or other
// insights
package wavolate

import (
	"errors"
	"fmt"
	"io"
	"os"
	"time"

	"github.com/asciifaceman/wavolator/pkg/logging"

	"azul3d.org/engine/audio"
	// WAV variant
	_ "azul3d.org/engine/audio/wav"

	"github.com/asciifaceman/wavolator/pkg/reducer"

	"go.uber.org/zap"
)

// Wavolate encapsulates our wav sampling and processing
type Wavolate struct {
	r          io.Reader
	filename   string
	resolution uint
	reducer    reducer.Reducer
	Logger     *zap.Logger
}

// New returns a new Wavolator
func New(optFnc ...Option) (*Wavolate, error) {
	options := &Options{}
	for _, opt := range optFnc {
		opt(options)
	}

	if options.Filename == "" {
		return nil, errors.New("filename is required")
	}

	if options.Logger == nil {
		options.Logger = logging.Logger()
	}

	reader, err := os.Open(options.Filename)
	if err != nil {
		return nil, err
	}

	w := &Wavolate{
		r:          reader,
		resolution: options.Resolution,
		filename:   options.Filename,
		Logger:     options.Logger,
		reducer:    options.Reducer,
	}

	return w, nil
}

// Sample is a full raw sample
type Sample struct {
	// Timecode is the point in the file (second*resolution)
	Timecode int
	// Samples is our full size sample set
	Sample audio.Float64
}

// SampleSet represents a file of samples
type SampleSet struct {
	Length  int
	Samples []*Sample
}

// ReducedSample is a reduced sample
type ReducedSample struct {
	// Timecode is the point in the file (second*resolution)
	Timecode int
	// Samples is our full size sample set
	Sample []float64
}

// ReducedSampleSet is a reduced set of samples
type ReducedSampleSet struct {
	Length  int
	Samples []*ReducedSample
}

// Sample samples  the file and  returns either
// a collection of samples  or  an error
func (w *Wavolate) Sample() (*SampleSet, error) {
	// Timelock
	start := time.Now()
	w.Logger.Info("Beginning sampling...",
		zap.String("filename", w.filename),
		zap.String("resolution", fmt.Sprintf("%vs", w.resolution)),
	)

	// Create new decoder
	decoder, _, err := audio.NewDecoder(w.r)
	if err != nil {
		return nil, err
	}

	config := decoder.Config()

	var seconds int
	samples := make(audio.Float64, uint(config.SampleRate*config.Channels)/w.resolution)
	sampleSet := &SampleSet{}
	for {
		_, err := decoder.Read(samples)
		if err != nil && err != audio.EOS {
			return nil, err
		}

		thisSample := &Sample{
			Timecode: seconds,
			Sample:   samples,
		}
		sampleSet.Samples = append(sampleSet.Samples, thisSample)

		if err == audio.EOS {
			break
		}
		seconds++
	}
	sampleSet.Length = seconds
	end := time.Now()
	duration := end.Sub(start)
	w.Logger.Info("Sampling complete...",
		zap.String("filename", w.filename),
		zap.String("resolution", fmt.Sprintf("%vs", w.resolution)),
		zap.String("process duration", fmt.Sprintf("%v", duration)),
		zap.String("file length", fmt.Sprintf("%d", seconds)),
	)
	return sampleSet, nil
}

// Reduce reduces the sample set with the configured reducer
func (w *Wavolate) Reduce(samples *SampleSet) *ReducedSampleSet {
	s := &ReducedSampleSet{}

	for _, sample := range samples.Samples {
		ret := w.reducer.Reduce(sample.Sample)
		thisSample := &ReducedSample{
			Timecode: sample.Timecode,
			Sample:   ret,
		}
		s.Samples = append(s.Samples, thisSample)
	}

	return s
}
