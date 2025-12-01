package task

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
)

func commandDelete(state *state, args []string) error {
	if len(args) < 1 {
		return ErrMissingArg
	}

	if len(args) > 1 {
		return ErrTooManyArgs
	}

	input := strings.TrimSpace(args[0])
	if input == "" {
		return ErrEmptyArgs
	}

	id, err := strconv.Atoi(input)
	if err != nil {
		return err
	}

	deleted := false
	state.Tasks = slices.DeleteFunc(state.Tasks, func(t Task) bool {
		if t.ID == id {
			fmt.Printf("Deleted Task (ID: %d) %s\n", t.ID, t.Description)
			deleted = true
			return true
		}
		return false
	})

	if !deleted {
		return ErrTaskNotFound
	}

	return saveState(state)
}
