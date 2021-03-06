// Copyright © 2018 Chris Holmes chris@holmser.net
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

	"github.com/aws/aws-sdk-go/service/codecommit"
	"github.com/spf13/cobra"
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new Repo",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		//fmt.Println("create called")
		codeCommit := initAWS()
		if repoName != "" {
			res, err := codeCommit.CreateRepository(&codecommit.CreateRepositoryInput{RepositoryName: &repoName})
			if err != nil {
				fmt.Println()
			}
			fmt.Println(res)
		}
	},
}
var repoName string

func init() {
	rootCmd.AddCommand(createCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createCmd.PersistentFlags().String("foo", "", "A help for foo")
	createCmd.Flags().StringVarP(&repoName, "name", "n", "", "repository name")
	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
