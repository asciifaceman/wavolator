// Copyright © 2019 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"

	"go.uber.org/zap"

	"github.com/asciifaceman/wavolator/pkg/graph"
	"github.com/asciifaceman/wavolator/pkg/logging"
	"github.com/asciifaceman/wavolator/pkg/wavolate"

	"github.com/spf13/cobra"
)

// graphCmd represents the graph command
var graphCmd = &cobra.Command{
	Use:   "graph",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("graph called")
		logger := logging.Logger()
		w, err := wavolate.New(
			wavolate.WithFile(filename),
			wavolate.WithReducer(reducer),
			wavolate.WithResolution(uint(resolution)),
			wavolate.WithLogger(),
		)
		if err != nil {
			fmt.Println("[ERRO] - Failed to start up: ", err)
			return
		}

		set, err := w.Sample()
		if err != nil {
			w.Logger.Fatal("Failed to sample file.",
				zap.String("error", err.Error()),
			)
		}

		reduced := w.Reduce(set)

		grapher, err := graph.New(
			graph.WithFile(filename),
			graph.WithLogger(w.Logger),
		)
		if err != nil {
			logger.Fatal("Failed to create grapher.",
				zap.String("error", err.Error()),
			)
		}
		err = grapher.DrawAndWrite(reduced)
		if err != nil {
			logger.Fatal("Failed to write graph.",
				zap.String("error", err.Error()),
			)
		}

	},
}

func init() {
	processCmd.AddCommand(graphCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// graphCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// graphCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
