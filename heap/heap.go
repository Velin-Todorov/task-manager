package heap

import (
	"container/heap"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"task-manager/models"
)

// TaskHeap is a min-heap of tasks.
type TaskHeap []models.Task

func (h TaskHeap) Len() int           { return len(h) }
func (h TaskHeap) Less(i, j int) bool { return h[i].Priority < h[j].Priority }
func (h TaskHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *TaskHeap) Push(x interface{}) {
	*h = append(*h, x.(models.Task))
}

func (h *TaskHeap) Pop() interface{} {
	old := *h
	n := len(old)
	item := old[n-1]
	*h = old[0 : n-1]
	return item
}

// Global task heap
var Tasks TaskHeap

func InitHeap() {
	LoadHeap()
	heap.Init(&Tasks)
}

const heapFile = "tasks.json"

// SaveHeap saves the heap to a file
func SaveHeap() {
	data, err := json.Marshal(Tasks)
	if err != nil {
		fmt.Println("Error saving heap:", err)
		return
	}
	err = ioutil.WriteFile(heapFile, data, 0644)
	if err != nil {
		fmt.Println("Error writing heap to file:", err)
	}
}

// LoadHeap loads the heap from a file
func LoadHeap() {
	if _, err := os.Stat(heapFile); err == nil {
		data, err := ioutil.ReadFile(heapFile)
		if err != nil {
			fmt.Println("Error reading heap file:", err)
			return
		}
		err = json.Unmarshal(data, &Tasks)
		if err != nil {
			fmt.Println("Error unmarshalling heap data:", err)
		}
	}
}
