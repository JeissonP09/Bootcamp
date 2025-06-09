package todo

import (
	"fmt"
	"time"
)

type item struct {
	Task        string
	Done        bool
	CreatedAt   time.Time
	CompletedAt time.Time
}

type List struct {
	items []item
}

func (l *List) Add(title string) {
	newItem := item{Task: title, Done: false}
	l.items = append(l.items, newItem)
}

func (l *List) Complete(index int) error {
	if index < 0 || index >= len(l.items) {
		return fmt.Errorf("invalid index")
	}
	l.items[index].Done = true
	l.items[index].CompletedAt = time.Now()

	return nil
}

func (l *List) Delete(index int) error {
	if index < 0 || index >= len(l.items) {
		return fmt.Errorf("invalid index")
	}
	l.items = append(l.items[:index], l.items[index+1:]...)

	return nil
}
