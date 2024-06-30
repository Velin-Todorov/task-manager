package crud

import (
	"container/heap"
	"fmt"
	"strconv"
	taskHeap "task-manager/heap"
	"task-manager/models"
	"github.com/spf13/cobra"
)

var (
	prio string
	name string
)

// AddCmd represents the add command
var AddCmd = &cobra.Command{
	Use:   "add",
	Short: "Adds a task",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		prioAsInt, err := strconv.Atoi(prio)
		if err != nil {
			fmt.Println("Invalid priority")
			return
		}

		if prioAsInt < 0 || prioAsInt > 10 {
			fmt.Println("Priority can be between 0 and 10 only.")
			return
		}

		task := models.Task{
			ID:       len(taskHeap.Tasks) + 1,
			Name:     name,
			Priority: prioAsInt,
		}

		heap.Push(&taskHeap.Tasks, task)
		taskHeap.SaveHeap()
		fmt.Printf("Added task '%s' with priority %d\n", name, prioAsInt)
	},
}

func init() {
	AddCmd.Flags().StringVarP(&name, "name", "n", "", "The name of the task")
	AddCmd.Flags().StringVarP(&prio, "prio", "p", "", "The priority of the task")

	if err := AddCmd.MarkFlagRequired("prio"); err != nil {
		fmt.Println(err)
	}

	if err := AddCmd.MarkFlagRequired("name"); err != nil {
		fmt.Println(err)
	}
}
