package task

import (
	"errors"
	"fmt"
	"slices"
)

// Command definition of the Task Tracker CLI.
type command struct {
	name        string
	description string
	usage       string
	callback    func(*state, []string) error
}

type CLI struct {
	state    *state
	commands map[string]command
}

var (
    ErrTaskNotFound = errors.New("task not found")
    ErrInvalidArg       = errors.New("invalid argument")
    ErrEmptyArgs        = errors.New("argument must be non-empty")
    ErrTooManyArgs      = errors.New("too many arguments")
    ErrMissingArg       = errors.New("not enough arguments")
)

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
			"mark-in-progress": {
				name: "mark-in-progress",
				description: "Marks the task status of a given ID as 'in-progress'.",
				usage: "mark-as-in-progress <id>",
				callback: commandMarkInProgress,
			},
			"mark-done": {
				name: "mark-done",
				description: "Marks the task status of a given ID as 'done'.",
				usage: "mark-done <id>",
				callback: commandMarkDone,
			},
		},
	}, nil
}

// Run the CLI state with the given arguments.
func (c *CLI) Run(args []string) error {
	name := args[0]
	commandArgs := args[1:]

	cmd, exists := c.commands[name]
	if !exists {
		return fmt.Errorf("command '%s' does not exist", name)
	}

	if err := cmd.callback(c.state, commandArgs); err != nil {
		fmt.Println("Usage:", cmd.usage)
		return err
	}

	return nil
}

// Display Task Tracker CLI commands.
func (c *CLI) DisplayCommands() {
	keys := []string{}
	for k := range c.commands {
		keys = append(keys, k)
	}
	slices.Sort(keys)

	fmt.Println("Task Tracker CLI v1.0")
	fmt.Println("Commands:")
	fmt.Println()

	for _, k := range keys {
		cmd := c.commands[k]
		fmt.Printf("* %s - %s\n  Usage: %s\n\n", cmd.name, cmd.description, cmd.usage)
	}
}
