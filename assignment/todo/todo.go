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
	newItem := item{Title: title, Done: false}
	*l = append(*l, newItem)
}

func (l *List) Complete(index int) error {
	if index < 0 || index >= len(*l) {
		return fmt.Errorf("Invalid index")
	}
	(*l)[index].Done = true
	(*l)[index].CompletedAt = time.Now()

	return nil
}

func (l *List) Delete(index int) error {
	if index < 0 || index >= len(*l) {
		return fmt.Errorf("Invalid index")
	}
	*l = append((*l)[:index], (*l)[index+1:]...)

	return nil
}
