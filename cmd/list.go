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
	"os"

	"github.com/ghokun/climan/cmd/tools"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run:     list,
	Aliases: []string{"l", "li", "lis"},
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

func list(cmd *cobra.Command, args []string) {


	
	if len(args) > 0 {
		listOne(args[0])
	} else {
		listAll()
	}
}

func listOne(name string) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetAutoWrapText(false)
	table.SetHeader([]string{"Tool", "Versions", "Status", "Path"})
	tool := tools.Tools[name]
	for _, version := range tool.List() {
		table.Append([]string{"", version, "-", "/home/gokhun/.."})
	}
	table.Render()
}

func listAll() {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Tool", "Latest", "Current", "Description", "Type"})
	table.SetAutoWrapText(false)

	for _, tool := range tools.Tools {
		table.Append([]string{tool.Name(), tool.Latest(), tool.Current(), tool.Description(), tool.Type()})
	}
	table.Render()
}
