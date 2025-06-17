package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
	"todo"
)

const (
	fileName = ".todo.json"
	NoAction = -1
)

func main() {
	add := flag.Bool("add", false, "Add a new task")
	complete := flag.Int("complete", -1, "Complete a task")
	delete := flag.Int("delete", -1, "Delete a task")

	flag.Parse()

	var list todo.List

	if err := list.Get(fileName); err != nil {
		fmt.Fprintf(os.Stderr, "Error getting tasks: %v\n", err)
		os.Exit(1)
	}

	args := flag.Args()

	if *add || (len(args) > 0 && *complete == NoAction && *delete == NoAction) {
		task := strings.Join(args, " ")
		if task == "" {
			fmt.Fprintf(os.Stderr, "Task cannot be empty")
			os.Exit(1)
		}
		list.Add(task)
	} else if *complete >= 0 {
		err := list.Complete(*complete - 1)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error completing task: %v\n", err)
			os.Exit(1)
		}
	} else if *delete >= 0 {
		err := list.Delete(*delete - 1)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error deleting task: %v\n", err)
			os.Exit(1)
		}
	} else {
		for _, item := range list {
			fmt.Println(item.Task)
		}
	}

	if err := list.Save(fileName); err != nil {
		fmt.Fprintf(os.Stderr, "Error saving ToDo list: %v\n", err)
		os.Exit(1)
	}
}
