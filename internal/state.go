package task

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

// State object of the Task Tracker CLI.
type state struct {
	NextID int    `json:"next_id"`
	Tasks  []Task `json:"tasks"`
}

var (
	cwd, _   = os.Getwd()
	filename = "/.tasktracker.json"
)

// Reads the saved state from JSON file.
func readState() (*state, error) {
	path := filepath.Join(cwd, filename)

	if _, err := os.Stat(path); os.IsNotExist(err) {
		s := state{NextID: 0, Tasks: []Task{}}
		if err := saveState(&s); err != nil {
			return nil, err
		}
		return &s, nil
	}

	file, err := os.OpenFile(path, os.O_RDONLY, 0)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var s state
	err = json.NewDecoder(file).Decode(&s)
	if err == nil {
		return &s, nil
	}

	if err == io.EOF {
		newState := state{NextID: 0, Tasks: []Task{}}
		if err := saveState(&newState); err != nil {
			return nil, err
		}
		return &newState, nil
	}

	var syntaxErr *json.SyntaxError
	if errors.As(err, &syntaxErr) {
		newState := state{NextID: 0, Tasks: []Task{}}
		if err := saveState(&newState); err != nil {
			return nil, err
		}
		return &newState, nil
	}

	var typeErr *json.UnmarshalTypeError
	if errors.As(err, &typeErr) {
		newState := state{NextID: 0, Tasks: []Task{}}
		if err := saveState(&newState); err != nil {
			return nil, err
		}
		return &newState, nil
	}

	return nil, err
}

// Writes the given application state to the JSON file.
func saveState(state *state) error {
	path := filepath.Join(cwd, filename)

	file, err := os.OpenFile(path, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	return json.NewEncoder(file).Encode(state)
}

func (s state) displayTasks(pred func(t Task) bool) {
	sb := strings.Builder{}
	idx := 1

	for _, t := range s.Tasks {
		if pred(t) {
			sb.WriteString(fmt.Sprintf("%d. %s\n", idx, t))
			idx++
		}
	}

	if sb.String() == "" {
		sb.WriteString("No tasks to display.")
	}

	fmt.Println(sb.String())
}
