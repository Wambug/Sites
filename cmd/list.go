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
	"fmt"

	"github.com/spf13/cobra"
	"github.com/wambug/Sites/db"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists all the webpages and time remaining to be launched on the browser",

	Run: func(cmd *cobra.Command, args []string) {
		sites, err := db.AllSites()
		if err != nil {
			fmt.Println("Something went wrong:", err.Error())
			return
		}
		for k, v := range sites {
			fmt.Printf("%d. Site:%s Duration: %s\n", k, v.Url, v.Duration)
		}
		for {
			sites, err = db.AllSites()
			if err != nil {
				fmt.Println("Something went wrong:", err.Error())
				return
			}
			if len(sites) == 0 {
				break
			}
			for _, v := range sites {
				err = db.DeleteSite(&v)
				if err != nil {
					fmt.Println("something went wrong", err.Error())
				}
			}
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
