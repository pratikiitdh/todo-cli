/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"task/db"

	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists all todos",
	Run: func(cmd *cobra.Command, args []string) {
		tasks, err := db.AllTask()
		if err != nil {
			fmt.Println("Something went wrongs:", err)
			os.Exit(1)
		}

		if len(tasks) == 0 {
			fmt.Println("No tasks left")
		}

		for ind, task := range tasks {
			fmt.Printf("%d: %s\n", ind+1, task.Value)
		}

	},
}

func init() {
	RootCmd.AddCommand(listCmd)
}
