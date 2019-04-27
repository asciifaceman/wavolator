package reducer

import (
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
	const modulo = 10

	for i := range samples {
		if (i % modulo) == 0 {
			sum := samples.At(i)
			for merge := 0; merge <= modulo-1; merge++ {
				if len(samples)-1 < i+merge {
					sum += samples.At(i + merge)
				}
			}

			//normalized := math.Exp(sum)
			//if sum < 0 {
			//	normalized = -normalized
			//}
			sampleSlice = append(sampleSlice, sum/modulo)
		}
	}
	return sampleSlice

}
