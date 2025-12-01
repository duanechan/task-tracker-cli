package task

import (
	"encoding/json"
	"os"
	"path/filepath"
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
		if err := saveState(&state{NextID: 0, Tasks: []Task{}}); err != nil {
			return nil, err
		}
	}

	file, err := os.OpenFile(path, os.O_RDONLY, 06444)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var state state
	if err = json.NewDecoder(file).Decode(&state); err != nil {
		return nil, err
	}

	return &state, nil
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

// Truncates the JSON file to an empty state.
func resetState() {
	new := &state{NextID: 0, Tasks: []Task{}}
	saveState(new)
}

// Load application state or return an error.
func LoadCLI() (*CLI, error) {
	s, err := readState()
	if err != nil {
		return nil, err
	}

	return &CLI{
		state: s,
		commands: map[string]command{
			"add": {
				name:        "add",
				description: "Adds a task to the list.",
				usage:       "add <description>",
				callback:    commandAdd,
			},
			"update": {
				name:        "update",
				description: "Updates the task of a given ID with an updated description.",
				usage:       "update <id> <updated_description>",
				callback:    commandUpdate,
			},
			"delete": {
				name:        "delete",
				description: "Deletes a task of a given ID.",
				usage:       "delete <id>",
				callback:    commandDelete,
			},
			"list": {
				name:        "list",
				description: "Lists tasks by status or all.",
				usage:       "list [done|todo|in-progress]",
				callback:    commandList,
			},
		},
	}, nil
}
