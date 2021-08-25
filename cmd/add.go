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
	"time"

	"github.com/spf13/cobra"
	"github.com/wambug/Sites/db"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Adds the site",
	Long: `Add the site you want to view later

`,
	Run: func(cmd *cobra.Command, args []string) {
		site := args[0]
		_, err := db.AddSite(site, duration)
		if err != nil {
			panic(err)
		}
		fmt.Printf("added : %s and will open %s later\n", site, duration)

	},
}

var duration time.Duration

func init() {
	rootCmd.AddCommand(addCmd)
	u, _ := time.ParseDuration("1s")
	addCmd.Flags().DurationVarP(&duration, "duration", "d", u, "how much time it will take to launch webpage")
	addCmd.MarkFlagRequired("duration")
	//fmt.Println(toggle)
	//rootCmd.PersistentFlags().StringVar(&author, "author", "YOUR NAME", "Author name for copyright attribution")
	//viper.BindPFlag("author", addCmd.Flags())

}
