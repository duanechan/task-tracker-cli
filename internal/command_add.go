package task

import "errors"

func commandAdd(state *state, args []string) error {
	if len(args) < 1 || len(args) > 1 {
		return errors.New("only one argument, task description, is required")
	}
	return nil
}
