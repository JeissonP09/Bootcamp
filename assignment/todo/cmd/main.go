package main

import (
	"flag"
	"fmt"
	"os"
	"todo"
)

const fileName = ".todo.json"

func main() {
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
		for _, item := range l {
			if !item.Done {
				fmt.Printf("Title: %s, Done: %t, CreatedAt: %s, CompletedAt: %s", item.Task, item.Done, item.CreatedAt, item.CompletedAt)
			}
		}
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

	if *task != "" {
		if *task == "" {
			fmt.Fprintf(os.Stderr, "Task cannot be empty")
			os.Exit(1)
		}
		l.Add(*task)
		if err := l.Save(fileName); err != nil {
			fmt.Fprintf(os.Stderr, "Error saving ToDo list: %v\n", err)
			os.Exit(1)
		}
		return
	}

	fmt.Fprintf(os.Stderr, "You must provide a valid command")
	os.Exit(1)
}
