// Package reducer contains reducers for processing down
// data sets
package reducer

import (
	"fmt"
	"math"
	"strings"

	"azul3d.org/engine/audio"
)

// NewReducerFunc is an abstract of a NewReducer
type NewReducerFunc func()

// Reducer encapsulates a reduction process
type Reducer interface {
	Reduce(audio.Float64) []float64
}

// Ensure the base reducer conforms
var _ Reducer = (*BaseReducer)(nil)

// BaseReducer implements Reducer without reduction
type BaseReducer struct {
}

// NewBaseReducer returns a new Default
func NewBaseReducer() *BaseReducer {
	return &BaseReducer{}
}

// Reduce returns the given dataset in native float64
func (b *BaseReducer) Reduce(samples audio.Float64) []float64 {
	var generic []float64
	for _, sample := range samples {
		generic = append(generic, math.Exp(sample))
	}
	return generic
}

// NewReducer returns a new reducer func
func NewReducer(reducer string) Reducer {
	reducer = strings.ToLower(reducer)
	switch reducer {
	case "rms":
		return NewRMSReducer()
	case "half", "halve":
		return NewHalfReducer()
	case "decimate":
		return NewDecimateReducer()
	default:
		return NewBaseReducer()
	}
}

// Normalize converts a float >1/<-1 to being <1/>-1
func Normalize(n float64) float64 {
	if n > 1.00000000 || n < -1.00000000 {
		fmt.Println("normalizing...")
		return Normalize(n / 100)
	}
	return n
}
