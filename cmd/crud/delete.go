/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>

*/
package crud

import (
	"github.com/spf13/cobra"
)

var (
	id string
)

// deleteCmd represents the delete command
var DeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Deletes a task with a given id",
	Long: ``,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
	DeleteCmd.Flags().StringVar(&id, "id", "", "The id of the task to be deleted")
}
