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
	"github.com/spf13/cobra"
)

var (
	resolution = 1
	reducer    = "none"
)
var filename string

// processCmd represents the process command
var processCmd = &cobra.Command{
	Use:   "process",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	//	Run: func(cmd *cobra.Command, args []string) {
	//		fmt.Println("process called")
	//	},
}

func init() {
	rootCmd.AddCommand(processCmd)
	processCmd.PersistentFlags().StringVarP(&filename, "filename", "f", filename, "[f]ilename - full path to the file you wish to process")
	processCmd.PersistentFlags().StringVarP(&reducer, "reducer", "d", reducer, "re[d]ucer - to use on the dataset")
	processCmd.PersistentFlags().IntVarP(&resolution, "resolution", "r", resolution, "[r]esolution - whole int resolution for samples to be taken, in seconds")
	processCmd.MarkPersistentFlagRequired("filename")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// processCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// processCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
