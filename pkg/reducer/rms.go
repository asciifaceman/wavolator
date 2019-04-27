package reducer

import (
	"math"

	"azul3d.org/engine/audio"
)

// Ensure the reducer conforms
var _ Reducer = (*RMSReducer)(nil)

// RMSReducer is a Reducer which calculates the RMS
// of a slice of float64 audio samples, enabling the measure
// of magnitude over the  entire set
// modified from https://github.com/mdlayher/waveform/blob/master/samplereducefunc.go
type RMSReducer struct {
}

// NewRMSReducer returns a new RMSReducer
func NewRMSReducer() *RMSReducer {
	return &RMSReducer{}
}

// Reduce  performs the RMSReducer  reduction
func (r *RMSReducer) Reduce(samples audio.Float64) []float64 {
	// Square and sum all input samples
	var sumSquare float64
	//spew.Dump(len(samples))
	for i := range samples {
		//this := math.Exp(samples.At(i))
		sumSquare += math.Pow(samples.At(i), 2)
	}

	//spew.Dump(sumSquare)
	// Multiply squared sum by length of samples slice, return square root
	squared := math.Sqrt(sumSquare / float64(samples.Len()))
	//spew.Dump(fmt.Sprintf("==%06f", squared))
	//normalized := math.Exp(squared)

	// Append to a slice to conform
	var sampleSlice []float64
	sampleSlice = append(sampleSlice, squared)
	return sampleSlice
}
