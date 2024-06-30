/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"container/heap"
	taskHeap "task-manager/heap"
	"task-manager/cmd"
)

func main() {
	heap.Init(&taskHeap.Tasks)
	cmd.Execute()
}
