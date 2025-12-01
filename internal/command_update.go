package task

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

func commandUpdate(state *state, args []string) error {
	if len(args) < 2 {
		return ErrMissingArg
	}

	if len(args) > 2 {
		return ErrTooManyArgs
	}

	idString, updatedDescription := strings.TrimSpace(args[0]), strings.TrimSpace(args[1])
	if idString == "" {
		return ErrEmptyArgs
	}

	id, err := strconv.Atoi(idString)
	if err != nil {
		return err
	}

	if updatedDescription == "" {
		return ErrEmptyArgs
	}

	updated := false
	for i, t := range state.Tasks {
		if t.ID == id {
			fmt.Printf("Updated Task (ID: %d) description to %s\n", t.ID, updatedDescription)
			state.Tasks[i].Description = updatedDescription
			state.Tasks[i].UpdatedAt = time.Now()
			updated = true
			break
		}
	}

	if !updated {
		return ErrTaskNotFound
	}

	return saveState(state)
}
