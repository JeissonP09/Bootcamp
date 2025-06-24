package main

import (
	"flag"
	"fmt"
	"os"
	"todo"
)

var fileName string

func main() {
	fileName = ".todo.json"
	if env := os.Getenv("TODO_FILENAME"); env != "" {
		fileName = env
	}

	list := flag.Bool("list", false, "List incomplete tasks")
	task := flag.String("task", "", "Add a new task")
	complete := flag.Int("complete", -1, "Complete a task")
	delete := flag.Int("delete", -1, "Delete a task")
	flag.Parse()

	var l todo.List

	if err := l.Get(fileName); err != nil {
		fmt.Fprintf(os.Stderr, "Error getting tasks: %v\n", err)
		os.Exit(1)
	}

	if *list {
		fmt.Print(l)
		return
	}

	if *complete != -1 {
		if err := l.Complete(*complete - 1); err != nil {
			fmt.Fprintf(os.Stderr, "Error completing task: %v\n", err)
			os.Exit(1)
		}
		if err := l.Save(fileName); err != nil {
			fmt.Fprintf(os.Stderr, "Error saving ToDo list: %v\n", err)
			os.Exit(1)
		}
		return
	}

	if *delete != -1 {
		if err := l.Delete(*delete - 1); err != nil {
			fmt.Fprintf(os.Stderr, "Error deleting task: %v\n", err)
			os.Exit(1)
		}
		if err := l.Save(fileName); err != nil {
			fmt.Fprintf(os.Stderr, "Error saving ToDo list: %v\n", err)
			os.Exit(1)
		}
		return
	}

	if *task == "" {
		fmt.Fprintf(os.Stderr, "Task cannot be empty")
		os.Exit(1)
	}
	l.Add(*task)
	if err := l.Save(fileName); err != nil {
		fmt.Fprintf(os.Stderr, "Error saving ToDo list: %v\n", err)
		os.Exit(1)
	}
}
