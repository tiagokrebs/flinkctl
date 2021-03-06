/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/spf13/cobra"
	j "github.com/tiagokrebs/flinkctl/internal/platform/json"
)

type Jobs struct {
	Jobs []struct {
		Id     string `json:"id"`
		Status string `json:"status"`
	} `json:"jobs"`
}

type Job struct {
	Jid   string `json:"jid"`
	Name  string `json:"name"`
	State string `json:"state"`
}

var Verbose bool

// jobsCmd represents the jobs command
var jobsCmd = &cobra.Command{
	Use:   "jobs",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		response, err := http.Get("xxx:8081/v1/jobs")

		if err != nil {
			fmt.Print(err.Error())
			os.Exit(1)
		}

		responseData, err := ioutil.ReadAll(response.Body)
		if err != nil {
			log.Fatal(err)
		}

		data := []byte(responseData)

		jobs := Jobs{}
		data = []byte(responseData)
		json.Unmarshal(data, &jobs)

		// The Job.Id is not so helpful, we need at least the name of it
		for _, job := range jobs.Jobs {
			uri := fmt.Sprintf("xxx:8081/v1/jobs/%s", job.Id)
			response, err := http.Get(uri)

			if err != nil {
				fmt.Print(err.Error())
				os.Exit(1)
			}

			responseData, err := ioutil.ReadAll(response.Body)
			if err != nil {
				log.Fatal(err)
			}

			if Verbose {
				data := []byte(responseData)
				prettyJSON, err := j.FormatJSON(data)
				if err != nil {
					log.Fatal(err)
				}

				fmt.Println(string(prettyJSON))
			} else {
				job := Job{}
				data = []byte(responseData)
				json.Unmarshal(data, &job)
				fmt.Printf("Job: %s\tState: %s\n", job.Name, job.State)
			}
		}

	},
}

func init() {
	getCmd.AddCommand(jobsCmd)

	// Here you will define your flags and jobsuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// jobsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Verbose flag will show all raw response from REST API
	rootCmd.PersistentFlags().BoolVarP(&Verbose, "verbose", "v", false, "verbose output")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// jobsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
