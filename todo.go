package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/aquasecurity/table"
)

type Todo struct {
	Title       string
	Completed   bool
	CreatedAt   time.Time
	CompletedAt *time.Time
}

type Todos []Todo

func (todos *Todos) add(title string) {
	fmt.Println("adding todos...")
	todo := Todo{
		Title:       title,
		Completed:   false,
		CreatedAt:   time.Now(),
		CompletedAt: nil,
	}

	*todos = append(*todos, todo)
}

func (todos *Todos) validateIndex(index int) error {
	if index < 0 || index >= len(*todos) {
		err := errors.New("Invalid index")
		fmt.Println(err)
		return err
	}
	return nil
}

func (todos *Todos) delete(index int) error {
	t := *todos
	if err := t.validateIndex(index); err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println(index)
	*todos = append(t[:index], t[index+1:]...)
	return nil
}

func (todos *Todos) toggle(index int) error {
	t := *todos
	if err := t.validateIndex(index); err != nil {
		return err
	}
	isCompleted := t[index].Completed

	if !isCompleted {
		completionTime := time.Now()
		t[index].CompletedAt = &completionTime
	}
	t[index].Completed = !isCompleted
	return nil
}

func (todos *Todos) edit(index int, title string) error {
	t := *todos
	if err := t.validateIndex(index); err != nil {
		return err
	}

	t[index].Title = title
	return nil
}

func (todos *Todos) print() {
	t := table.New(os.Stdout)
	t.SetRowLines(false)
	t.SetHeaders("#", "Title", "Completed", "Created At", "Completed At")
	for index, todo := range *todos {
		completed := "❌"
		completedAt := ""

		if todo.Completed {
			completed = "✅"
			if todo.CompletedAt != nil {
				completedAt = todo.CompletedAt.Format(time.RFC3339)
			}
		}

		t.AddRow(strconv.Itoa(index), todo.Title, completed, todo.CreatedAt.Format(time.RFC3339), completedAt)
	}
	t.Render()
}
