package main

import (
	"flag"
	"fmt"
	"os"
	"todo"
)

const fileName = ".todo.json"

func main() {
	listFlag := flag.Bool("list", false, "List incomplete tasks")
	taskFlag := flag.String("task", "", "Add a new task")
	completeFlag := flag.Int("complete", -1, "Complete a task")
	deleteFlag := flag.Int("delete", -1, "Delete a task")
	flag.Parse()

	var l todo.List

	if err := l.Get(fileName); err != nil {
		fmt.Fprintf(os.Stderr, "Error getting tasks: %v\n", err)
		os.Exit(1)
	}

	commandExecuted := false

	if *listFlag {
		for _, item := range l {
			if !item.Done {
				fmt.Printf("Title: %s, Done: %t, CreatedAt: %s, CompletedAt: %s", item.Task, item.Done, item.CreatedAt, item.CompletedAt)
			}
		}
		commandExecuted = true
	} else if *completeFlag != -1 {
		if err := l.Complete(*completeFlag - 1); err != nil {
			fmt.Fprintf(os.Stderr, "Error completing task: %v\n", err)
			os.Exit(1)
		}
		if err := l.Save(fileName); err != nil {
			fmt.Fprintf(os.Stderr, "Error saving ToDo list: %v\n", err)
			os.Exit(1)
		}
		commandExecuted = true
	} else if *deleteFlag != -1 {
		if err := l.Delete(*deleteFlag - 1); err != nil {
			fmt.Fprintf(os.Stderr, "Error deleting task: %v\n", err)
			os.Exit(1)
		}
		if err := l.Save(fileName); err != nil {
			fmt.Fprintf(os.Stderr, "Error saving ToDo list: %v\n", err)
			os.Exit(1)
		}
		commandExecuted = true
	} else if *taskFlag != "" {
		if *taskFlag == "" {
			fmt.Fprintf(os.Stderr, "Task cannot be empty")
			os.Exit(1)
		}
		l.Add(*taskFlag)
		if err := l.Save(fileName); err != nil {
			fmt.Fprintf(os.Stderr, "Error saving ToDo list: %v\n", err)
			os.Exit(1)
		}
		commandExecuted = true
	}

	if !commandExecuted {
		fmt.Fprintf(os.Stderr, "You must provide a valid command")
		os.Exit(1)
	}
}
