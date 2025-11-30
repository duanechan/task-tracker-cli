package task

import (
	"errors"
	"fmt"
	"strings"
	"time"
)

var (
	ErrEmptyDescription = errors.New("task description is required")
	ErrTooManyArgs      = errors.New("too many arguments")
	ErrMissingArg       = errors.New("missing one argument")
)

func commandAdd(state *state, args []string) error {
	if len(args) < 1 {
		return ErrMissingArg
	}

	if len(args) > 1 {
		return ErrTooManyArgs
	}

	description := strings.TrimSpace(args[0])
	if description == "" {
		return ErrEmptyDescription
	}

	state.Count++
	task := Task{
		ID:          state.Count,
		Description: description,
		Status:      Todo,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	state.Tasks = append(state.Tasks, task)

	fmt.Printf("Task added successfully (ID: %d)\n", state.Count)

	return saveState(state)
}
