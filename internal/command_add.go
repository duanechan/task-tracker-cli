package task

import (
	"fmt"
	"strings"
	"time"
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

	state.NextID++
	task := Task{
		ID:          state.NextID,
		Description: description,
		Status:      Todo,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	state.Tasks = append(state.Tasks, task)

	fmt.Printf("Task added successfully (ID: %d)\n", state.NextID)

	return saveState(state)
}
