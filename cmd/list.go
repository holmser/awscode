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

		sess := session.Must(session.NewSession(&aws.Config{
			MaxRetries: aws.Int(3),
		}))

		codeCommit := codecommit.New(sess, &aws.Config{
			Region: aws.String(Region),
		})
		repos, err := codeCommit.ListRepositories(nil)
		if err != nil {
			fmt.Println(err)
		}

		// declare channel for threading description API calls
		ch := make(chan *codecommit.GetRepositoryOutput)

		for _, repo := range repos.Repositories {
			// currently doesn't do anything important, but will be more important later
			go func(ch chan<- *codecommit.GetRepositoryOutput, rname *string) {
				out, err := codeCommit.GetRepository(&codecommit.GetRepositoryInput{
					RepositoryName: rname,
				})
				if err != nil {
					fmt.Println(err)
				}
				ch <- out
			}(ch, repo.RepositoryName)
		}
		// iterate over channel to print repos
		for range repos.Repositories {
			repo := <-ch
			fmt.Println(*repo.RepositoryMetadata.RepositoryName)
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
