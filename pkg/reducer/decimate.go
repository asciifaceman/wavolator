package reducer

import (
	"math"

	"azul3d.org/engine/audio"
)

var _ Reducer = (*DecimateReducer)(nil)

// DecimateReducer sums sets of 10 segments to
// reduce the overall dataset by 1/10
// while losing some fidelity
type DecimateReducer struct {
}

// NewDecimateReducer returns a new DecimateReducer
func NewDecimateReducer() *DecimateReducer {
	return &DecimateReducer{}
}

// Reduce perorms a DecimateReduce
func (c *DecimateReducer) Reduce(samples audio.Float64) []float64 {
	var sampleSlice []float64

	for i := range samples {
		if (i % 100) == 0 {
			sum := samples.At(i)
			for merge := 0; merge <= 99; merge++ {
				if len(samples)-1 < i+merge {
					sum += samples.At(i + merge)
				}
			}

			normalized := math.Exp(sum)
			if sum < 0 {
				normalized = -normalized
			}
			sampleSlice = append(sampleSlice, normalized/10)
		}
	}
	return sampleSlice

}
