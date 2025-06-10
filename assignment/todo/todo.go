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

type List []item

func (l *List) Add(title string) {
	newItem := item{Task: title, Done: false, CreatedAt: time.Now()}
	*l = append(*l, newItem)
}

func (l *List) Complete(index int) error {
	if index < 0 || index >= len(*l) {
		return fmt.Errorf("invalid index")
	}
	(*l)[index].Done = true
	(*l)[index].CompletedAt = time.Now()

	return nil
}

func (l *List) Delete(index int) error {
	if index < 0 || index >= len(*l) {
		return fmt.Errorf("invalid index")
	}
	*l = append((*l)[:index], (*l)[index+1:]...)

	return nil
}
