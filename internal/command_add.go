package task

import (
	"fmt"
	"strings"
	"time"
)

func commandAdd(c *CLI, args []string) error {
	if len(args) < 1 {
		return ErrMissingArg
	}

	if len(args) > 1 {
		return ErrTooManyArgs
	}

	description := strings.TrimSpace(args[0])
	if description == "" {
		return ErrEmptyArgs
	}

	c.state.NextID++
	task := Task{
		ID:          c.state.NextID,
		Description: description,
		Status:      Todo,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	c.state.Tasks = append(c.state.Tasks, task)

	fmt.Printf("Task added successfully: %s\n", task)

	return saveState(c.state)
}
