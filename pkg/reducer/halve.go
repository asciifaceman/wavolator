package reducer

import (
	"azul3d.org/engine/audio"
)

var _ Reducer = (*HalfReducer)(nil)

// HalfReducer sums parallel segments to
// reduce the overall dataset by 1/2
// while losing little fidelity
type HalfReducer struct {
}

// NewHalfReducer returns a new HalfReducer
func NewHalfReducer() *HalfReducer {
	return &HalfReducer{}
}

// Reduce perorms a HalfReduce
func (c *HalfReducer) Reduce(samples audio.Float64) []float64 {
	var sampleSlice []float64
	const modulo = 2

	for i := range samples {
		if (i % modulo) == 0 {
			sum := samples.At(i)
			if len(samples)-1 < i {
				sum += samples.At(i + 1)
			}
			//normalized := math.Exp(sum)
			//if sum < 0 {
			//	sum = -sum
			//}
			sampleSlice = append(sampleSlice, sum/modulo)
		}
	}
	return sampleSlice

}
