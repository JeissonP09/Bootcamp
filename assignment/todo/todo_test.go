package todo

import (
	"os"
	"testing"
)

func TestAdd(t *testing.T) {
	var l List
	task := "Create ToDo test"
	l.Add(task)

	if len(l) != 1 {
		t.Errorf("Expected: one task, but there is %d", len(l))
	}

	if l[0].Task != task {
		t.Errorf("Expected task: %q, but there is %q", task, l[0].Task)
	}
}

func TestComplete(t *testing.T) {
	var l List
	task := "Create PR"
	l.Add(task)

	err := l.Complete(0)
	if err != nil {
		t.Errorf("Not complete the task: %v", err)
	}

	if !l[0].Done {
		t.Errorf("Expected: Task marked completed")
	}
}

func TestDelete(t *testing.T) {
	var l List
	t_1 := "Task 1"
	t_2 := "Task 2"
	l.Add(t_1)
	l.Add(t_2)

	err := l.Delete(1)
	if err != nil {
		t.Errorf("The task could not be deleted %v", err)
	}

	if len(l) != 1 {
		t.Errorf("Expected [1] task, but there is [%d]", len(l))
	}

	if l[0].Task != t_1 {
		t.Errorf("Was expected %q as remaining task, but was found %q", t_1, l[0].Task)
	}
}

func TestSaveAndGet(t *testing.T) {
	tf, err := os.CreateTemp("", "todo_test_*.json")
	if err != nil {
		t.Fatalf("The temporary file could not be created: %v", err)
	}
	defer os.Remove(tf.Name())

	var l1 List
	l1.Add("Save task")
	l1.Add("Recover task")

	if err := l1.Save(tf.Name()); err != nil {
		t.Fatalf("Error saving file: %v", err)
	}

	var l2 List
	if err := l2.Get(tf.Name()); err != nil {
		t.Fatalf("Error reading file: %v", err)
	}

	if len(l2) != len(l1) {
		t.Errorf("Task was expected [%d], but obtained [%d]", len(l1), len(l2))
	}

	for i := range l1 {
		if l1[i].Task != l2[i].Task {
			t.Errorf("Task [%d]: expected %q, but obtained %q", i, l1[i].Task, l2[i].Task)
		}
	}
}
