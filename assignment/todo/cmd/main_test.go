package main_test

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"testing"
)

var (
	binName  = "todo"
	fileName = ".todo.json"
)

func TestMain(m *testing.M) {
	fmt.Println("Building tool...")

	if runtime.GOOS == "windows" {
		binName += ".exe"
	}

	build := exec.Command("go", "build", "-o", binName)
	err := build.Run()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Cannot build tool %s: %s", binName, err)
		os.Exit(1)
	}

	if err = os.WriteFile(fileName, []byte{}, 0644); err != nil {
		fmt.Fprintf(os.Stderr, "Cannot create file %s: %v\n", fileName, err)
		os.Exit(1)
	}

	fmt.Println("Running tests....")
	result := m.Run()

	fmt.Println("Cleaning up....")
	os.Remove(binName)
	os.Remove(fileName)

	os.Exit(result)
}

func TestTodoCLI(t *testing.T) {

	dir, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}

	cmdPath := filepath.Join(dir, binName)

	t.Run("AddNewTask", func(t *testing.T) {
		if err := os.WriteFile(fileName, []byte{}, 0644); err != nil {
			t.Fatal(err)
		}
		task := "New Task"
		cmd := exec.Command(cmdPath, "-task", task)
		if err := cmd.Run(); err != nil {
			t.Fatalf("Error add task: %v", err)
		}
	})

	t.Run("ListTasks", func(t *testing.T) {
		if err := os.WriteFile(fileName, []byte{}, 0644); err != nil {
			t.Fatal(err)
		}
		task := "New Task"
		cmdAdd := exec.Command(cmdPath, "-task", task)
		if err := cmdAdd.Run(); err != nil {
			t.Fatalf("Error add task: %v", err)
		}
		cmdList := exec.Command(cmdPath, "-list")
		out, err := cmdList.CombinedOutput()
		if err != nil {
			t.Fatalf("Error listing tasks: %v", err)
		}
		output := string(out)
		want := fmt.Sprintf("Incomplete task: - [ ] 0: %s\n", task)
		if output != want {
			t.Errorf("List output unexpected: \ngot: %q\nwant: %q", output, want)
		}
	})

	t.Run("CompleteTask", func(t *testing.T) {
		if err := os.WriteFile(fileName, []byte{}, 0644); err != nil {
			t.Fatal(err)
		}
		task := "New Task"
		cmdAdd := exec.Command(cmdPath, "-task", task)
		if err := cmdAdd.Run(); err != nil {
			t.Fatalf("Eror add task: %v", err)
		}
		cmdComplete := exec.Command(cmdPath, "-complete", "1")
		if err := cmdComplete.Run(); err != nil {
			t.Fatalf("Error completing task: %v", err)
		}
		cmdList := exec.Command(cmdPath, "-list")
		out, err := cmdList.CombinedOutput()
		if err != nil {
			t.Fatalf("Error listing taks: %v", err)
		}
		output := string(out)
		want := fmt.Sprintf("Complete task: - [X] 0: %s\n", task)
		if output != want {
			t.Errorf("List output unexpected: \ngot: %q\nwant: %q", output, want)
		}
	})

	t.Run("DeleteTask", func(t *testing.T) {
		if err := os.WriteFile(fileName, []byte{}, 0644); err != nil {
			t.Fatal(err)
		}
		task := "New Task"
		cmdAdd := exec.Command(cmdPath, "-task", task)
		if err := cmdAdd.Run(); err != nil {
			t.Fatalf("Error add taks: %v", err)
		}
		cmdDelete := exec.Command(cmdPath, "-delete", "1")
		if err := cmdDelete.Run(); err != nil {
			t.Fatalf("Error deleting task: %v", err)
		}
		cmdList := exec.Command(cmdPath, "-list")
		out, err := cmdList.CombinedOutput()
		if err != nil {
			t.Fatalf("Error lsiting taks: %v", err)
		}
		if len(strings.TrimSpace(string(out))) != 0 {
			t.Errorf("Expected no taks deletion, but got: %s", string(out))
		}
	})
}
