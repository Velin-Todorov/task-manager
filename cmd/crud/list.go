package crud

import (
	"fmt"
	"task-manager/heap"
	"github.com/spf13/cobra"
)

// ListCmd represents the list command
var ListCmd = &cobra.Command{
	Use:   "list",
	Short: "Prints out all the tasks ordered by their priority",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		heap.LoadHeap() // Ensure the heap is loaded from the file
		tasks := heap.Tasks

		if len(tasks) == 0 {
			fmt.Println("No tasks available.")
			return
		}

		for _, task := range tasks {
			fmt.Printf("Task: \n\tName: %s\n\tPriority: %d\n\tID: %d\n", task.Name, task.Priority, task.ID)
		}
	},
}
