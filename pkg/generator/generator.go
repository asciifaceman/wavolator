// Package generator generates a CSV training dataset
// for https://github.com/asciifaceman/wavolator-network
package generator

import (
	"encoding/csv"
	"errors"
	"fmt"
	"os"

	"github.com/asciifaceman/wavolator/pkg/logging"
	"go.uber.org/zap"
)

// Generator ...
type Generator struct {
	Filename string
	Label    string
	Logger   *zap.Logger
}

// Sample is a condensed sample for csv output for training
type Sample struct {
	Dataset []float64
}

// NewSample ...
func (s *Set) NewSample(data []float64) {
	s.Data = append(s.Data, &Sample{
		Dataset: data,
	})
}

// Set is a st of TrainingSample
type Set struct {
	Label string
	Data  []*Sample
}

// NewSet ...
func (g *Generator) NewSet() *Set {
	return &Set{
		Label: g.Label,
	}
}

func (s *Set) generateLine() *CSVData {
	csvd := &CSVData{}
	var line []string
	line = append(line, s.Label)
	for _, sample := range s.Data {
		for _, subSample := range sample.Dataset {
			line = append(line, fmt.Sprintf("%f", subSample))
		}
	}
	csvd.Data = line
	csvd.Length = len(line)
	return csvd
}

// Write  ...
func (g *Generator) Write(s *Set) error {
	dump := s.generateLine()
	f, err := os.Create(g.GetFilename())
	if err != nil {
		return err
	}
	defer f.Close()

	writer := csv.NewWriter(f)
	defer writer.Flush()

	err = writer.Write(dump.Data)
	if err != nil {
		return err
	}

	return nil
}

// CSVData is ready to write
type CSVData struct {
	Length int
	Data   []string
}

// New returns a new  generator
func New(optFnc ...Option) (*Generator, error) {
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

	if options.Label == "" {
		return nil, errors.New("label is required")
	}

	return &Generator{
		Filename: options.Filename,
		Label:    options.Label,
		Logger:   options.Logger,
	}, nil
}

// GetFilename  returns  a formatted filename
func (g *Generator) GetFilename() string {
	return fmt.Sprintf("%s-%s", g.Label, g.Filename)
}
