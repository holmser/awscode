// Copyright © 2018 NAME HERE <EMAIL ADDRESS>
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

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/codecommit"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List codecommit repos",
	Long:  `This will list all CodeCommit repos your credentials have acess to`,
	Run: func(cmd *cobra.Command, args []string) {
		//fmt.Println("list called")
		sess := session.Must(session.NewSession(&aws.Config{
			MaxRetries: aws.Int(3),
		}))

		codeCommit := codecommit.New(sess, &aws.Config{
			Region: aws.String("us-east-1"),
		})
		repos, err := codeCommit.ListRepositories(nil)
		if err != nil {
			fmt.Println(err)
		}

		for _, repo := range repos.Repositories {
			fmt.Println("○", *repo.RepositoryName)

			out, err := codeCommit.GetRepository(&codecommit.GetRepositoryInput{
				RepositoryName: repo.RepositoryName,
			})

			if err != nil {
				fmt.Println()
			}
			fmt.Println("\t", *out.RepositoryMetadata.CloneUrlSsh)
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
