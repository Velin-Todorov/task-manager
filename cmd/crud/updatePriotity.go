/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package crud

import (
	"fmt"

	"github.com/spf13/cobra"
)


// var (
// 	id string
// )

// updatePriotityCmd represents the updatePriotity command
var UpdatePriotityCmd = &cobra.Command{
	Use:   "updatePriority",
	Short: "updates the priority of a task by id",
	Long: ``,
	Run: func(cmd *cobra.Command, args []string) {
		
	},
}

func init() {
	UpdatePriotityCmd.Flags().StringVarP(&prio, "prio", "p", "", "The id of the task")
	// UpdatePriotityCmd.Flags().StringVar(&id, "id", "", "The id of the task")


	if err := AddCmd.MarkFlagRequired("prio"); err != nil {
		fmt.Println(err)
	}

	// if err := AddCmd.MarkFlagRequired("id"); err != nil {
	// 	fmt.Println("HELOO")

	// 	fmt.Println(err)
	// }

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// updatePriotityCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// updatePriotityCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
