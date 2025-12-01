package task

import (
	"fmt"
	"slices"
	"strings"
)

func commandList(state *state, args []string) error {
	if len(args) > 1 {
		return ErrTooManyArgs
	}

	if len(args) == 0 {
		fmt.Println("List of all tasks:")
		state.displayTasks(func(t Task) bool {
			return true
		})
		return nil
	}

	validArgs := []string{"todo", "in-progress", "done"}

	status := strings.ToLower(strings.TrimSpace(args[0]))
	if status == "" {
		return ErrEmptyArgs
	}

	if !slices.Contains(validArgs, status) {
		return ErrInvalidArg
	}

	fmt.Println("List of", status, "tasks:")

	switch status {
	case "todo":
		state.displayTasks(func(t Task) bool {
			return t.Status == Todo
		})
	case "in-progress":
		state.displayTasks(func(t Task) bool {
			return t.Status == InProgress
		})
	case "done":
		state.displayTasks(func(t Task) bool {
			return t.Status == Done
		})
	}

	return nil
}
