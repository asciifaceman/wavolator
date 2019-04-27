// Package graph outputs a graph of the produced data
package graph

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"time"

	"go.uber.org/zap"

	"github.com/asciifaceman/wavolator/pkg/logging"
	"github.com/asciifaceman/wavolator/pkg/wavolate"
	chart "github.com/wcharczuk/go-chart"
)

// Graph wraps our charting functions
type Graph struct {
	Logger *zap.Logger
	Output string // filename
}

// New returns a new grapher
func New(optFnc ...Option) (*Graph, error) {
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

	g := &Graph{
		Logger: options.Logger,
		Output: options.Filename,
	}

	return g, nil
}

// Draw draws the input graph to a buffer
func (g *Graph) Draw(set *wavolate.ReducedSampleSet) (*bytes.Buffer, error) {
	g.Logger.Info("Preparing chart. This can take awhile if the dataset is not reduced...",
		zap.String("filename", g.Output),
	)
	x, y := g.process(set)

	start := time.Now()
	g.Logger.Info("... Generating chart...")
	graph := chart.Chart{
		XAxis: chart.XAxis{
			Style:     chart.StyleShow(), //enables / displays the x-axis
			Name:      "Seconds",
			NameStyle: chart.StyleShow(),
		},
		YAxis: chart.YAxis{
			Name:      "Pulse",
			NameStyle: chart.StyleShow(),
			Style:     chart.StyleShow(), //enables / displays the y-axis
			ValueFormatter: func(v interface{}) string {
				if vf, isFloat := v.(float64); isFloat {
					return fmt.Sprintf("%0.6f", vf)
				}
				return ""
			},
			//			Range: &chart.ContinuousRange{
			//				Min: -1.0,
			//				Max: 1.0,
			//			},
		},
		Series: []chart.Series{
			chart.ContinuousSeries{
				Style: chart.Style{
					Show:        true,
					StrokeColor: chart.GetDefaultColor(0).WithAlpha(64),
					//FillColor:   chart.GetDefaultColor(0).WithAlpha(64),
					StrokeWidth: 1,
					// StrokeWidth: 0.01, //for heatmap
				},
				XValues: x,
				YValues: y,
			},
		},
	}
	complete := time.Now()
	sub := complete.Sub(start)
	g.Logger.Info("... Finished chart...",
		zap.Duration("duration", sub),
	)

	g.Logger.Info("... Preparing file...",
		zap.String("filename", g.Output),
		zap.Int("segments", len(x)),
	)
	start = time.Now()
	// using my own buffer is 1s faster than using collector
	buffer := bytes.NewBuffer([]byte{})
	err := graph.Render(chart.PNG, buffer)
	if err != nil {
		return nil, err
	}
	complete = time.Now()
	sub = complete.Sub(start)
	g.Logger.Info("Finished preparing file...",
		zap.Duration("duration", sub),
	)

	return buffer, nil
}

// Write outputs the file
func (g *Graph) Write(buff *bytes.Buffer) error {
	g.Logger.Info("Writing chart...",
		zap.String("filename", g.Output),
	)
	err := ioutil.WriteFile(g.Output, buff.Bytes(), 0644)

	if err != nil {
		return err
	}
	g.Logger.Info("Done.",
		zap.String("filename", g.Output),
	)

	return nil
}

// DrawAndWrite draws then writes the result, equivalent of Write(draw)
func (g *Graph) DrawAndWrite(set *wavolate.ReducedSampleSet) error {
	buff, err := g.Draw(set)
	if err != nil {
		return err
	}
	err = g.Write(buff)
	if err != nil {
		return err
	}
	return nil
}

// process processes the data and returns the x and y values
func (g *Graph) process(set *wavolate.ReducedSampleSet) ([]float64, []float64) {
	g.Logger.Info("... Processing sample data for charting...")
	var x, y []float64
	//var timeObserved float64

	for _, sample := range set.Samples {
		if len(sample.Sample) == 1 {
			x = append(x, float64(sample.Timecode))
			y = append(y, sample.Sample[0])
			continue
		}
		for inner, sub := range sample.Sample {
			// sample.Timecode may represent thousands
			// of files, which suchs for graphing
			totalSegmentsInSecond := float64(len(sample.Sample))
			segmentsPerSecond := float64(60 / totalSegmentsInSecond)
			floatyInner := float64(inner)
			stamp := float64(float64(sample.Timecode) + (segmentsPerSecond * floatyInner))

			x = append(x, stamp)
			y = append(y, sub)
		}
	}

	g.Logger.Info("... Done processing.")
	return x, y
}

/*
// DrawGraph outputs a graph from input data
func DrawGraph(set *wavolate.ReducedSampleSet) error {
	fmt.Println("Drawing graph...")
	var xVal []float64
	var yVal []float64

	for parentIter, sample := range set.Samples {
		fmt.Println("Parent: ", parentIter)
		for iter, subSample := range sample.Sample {
			var subTime float64
			var err error
			if iter > 0 {
				subTime, err = strconv.ParseFloat(fmt.Sprintf("%d.%d", sample.Timecode, iter), 64)
				if err != nil {
					subTime = float64(sample.Timecode + (iter / 60))
				}
			} else {
				subTime = float64(sample.Timecode)
			}
			fmt.Println("Sub: ", subTime)
			thisX := float64(subTime)
			xVal = append(xVal, thisX)
			yVal = append(yVal, subSample*10000)
		}
	}

	spew.Dump(xVal)
	spew.Dump(yVal)

	graph := chart.Chart{
		XAxis: chart.XAxis{
			Style: chart.StyleShow(), //enables / displays the x-axis
		},
		YAxis: chart.YAxis{
			Style: chart.StyleShow(), //enables / displays the y-axis
		},
		Series: []chart.Series{
			chart.ContinuousSeries{
				Style: chart.Style{
					Show:        true,
					StrokeColor: chart.GetDefaultColor(0).WithAlpha(64),
					FillColor:   chart.GetDefaultColor(0).WithAlpha(64),
					StrokeWidth: 3,
				},
				XValues: xVal,
				YValues: yVal,
			},
		},
	}

	fmt.Println("Writing graph...")
	buffer := bytes.NewBuffer([]byte{})
	err := graph.Render(chart.PNG, buffer)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile("chart.png", buffer.Bytes(), 0644)

	if err != nil {
		return err
	}

	return nil
}
*/
