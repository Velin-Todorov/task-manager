package models

import (
	"strconv"
)

type Task struct {
	ID       int
	Name     string
	Priority int
}

func (t Task) String() string {
	return "[" + "Name: " + t.Name + "Priority: " + strconv.Itoa(t.Priority) + "]"
}