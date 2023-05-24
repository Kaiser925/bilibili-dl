// Developed by Kaiser925 on 2021/2/14.
// Lasted modified 2021/2/14.
// Copyright (c) 2021.  All rights reserved
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package cmd

import (
	"github.com/Kaiser925/bilibili-dl/pkg/fetch"
	"github.com/spf13/cobra"
)

var output string

// coverCmd represents the cover command
var coverCmd = &cobra.Command{
	Use:   "cover <url>",
	Short: "Get cover of video",
	Long:  "Get cover of video",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fetch.FetchCover(args[0], output)
	},
}

func init() {
	rootCmd.AddCommand(coverCmd)
	coverCmd.Flags().StringVarP(&output, "output", "o", "",
		"output of saved cover, default BV name")
}
