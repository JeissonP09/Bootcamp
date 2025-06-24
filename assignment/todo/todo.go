package todo

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"
)

type item struct {
	Task        string
	Done        bool
	CreatedAt   time.Time
	CompletedAt time.Time
}

type List []item

func (l List) String() string {
	var b strings.Builder
	for i, it := range l {
		status := " "
		prefix := "Incomplete task: "
		if it.Done {
			status = "X"
			prefix = "Complete task: "
		}
		fmt.Fprintf(&b, "%s- [%s] %d: %s\n", prefix, status, i, it.Task)
	}
	return b.String()
}

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

func (l *List) Save(filename string) error {
	data, err := json.Marshal(l)
	if err != nil {
		return err
	}
	err = os.WriteFile(filename, data, 0644)
	if err != nil {
		return err
	}
	return nil
}

func (l *List) Get(filename string) error {
	data, err := os.ReadFile(filename)
	if err != nil {
		if os.IsNotExist(err) {
			return nil
		}
		return err
	}
	if len(data) == 0 {
		return nil
	}
	if err = json.Unmarshal(data, l); err != nil {
		return err
	}
	return nil
}
