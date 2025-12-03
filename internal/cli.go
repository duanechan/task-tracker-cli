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
	params      map[string]string
	callback    func(*CLI, []string) error
}

type commands map[string]command

func initializeCommands() *commands {
	return &commands{
		"add": {
			name:        "add",
			description: "Adds a task to the list.",
			usage:       "add <description>",
			params: map[string]string{
				"description": "the task description",
			},
			callback: commandAdd,
		},
		"update": {
			name:        "update",
			description: "Updates the task of a given ID with an updated description.",
			usage:       "update <id> <description>",
			params: map[string]string{
				"id":          "the id of the task to be updated",
				"description": "the updated task description",
			},
			callback: commandUpdate,
		},
		"delete": {
			name:        "delete",
			description: "Deletes a task of a given ID.",
			usage:       "delete <id>",
			params: map[string]string{
				"id": "the id of the task to be deleted",
			},
			callback: commandDelete,
		},
		"list": {
			name:        "list",
			description: "Lists tasks by status or all.",
			usage:       "list [done|todo|progress]",
			params: map[string]string{
				"done":        "(optional) list all done tasks",
				"todo":        "(optional) list todo done tasks",
				"in-progress": "(optional) list in-progress done tasks",
			},
			callback: commandList,
		},
		"mark-in-progress": {
			name:        "mark-in-progress",
			description: "Marks the task status of a given ID as 'in-progress'.",
			usage:       "mark-as-in-progress <id>",
			params: map[string]string{
				"id": "the id of the task to be marked as in-progress",
			},
			callback: commandMarkInProgress,
		},
		"mark-done": {
			name:        "mark-done",
			description: "Marks the task status of a given ID as 'done'.",
			usage:       "mark-done <id>",
			params: map[string]string{
				"id": "the id of the task to be marked as done",
			},
			callback: commandMarkDone,
		},
		"help": {
			name:        "help",
			description: "Display list of commands",
			usage:       "help [command]",
			params: map[string]string{
				"command": "(optional) the name of the command in question",
			},
			callback: commandHelp,
		},
		"version": {
			name:        "version",
			description: "Check Task Tracker version",
			usage:       "version",
			params: map[string]string{
				"": "",
			},
			callback: commandVersion,
		},
	}
}

type CLI struct {
	version  string
	state    *state
	commands map[string]command
}

var (
	ErrTaskNotFound = errors.New("task not found")
	ErrInvalidArg   = errors.New("invalid argument")
	ErrEmptyArgs    = errors.New("argument must be non-empty")
	ErrTooManyArgs  = errors.New("too many arguments")
	ErrMissingArg   = errors.New("not enough arguments")
)

// Load application state or return an error.
func LoadCLI() (*CLI, error) {
	state, err := readState()
	if err != nil {
		return nil, err
	}

	return &CLI{
		version:  "v1.0.1",
		state:    state,
		commands: *initializeCommands(),
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

	if err := cmd.callback(c, commandArgs); err != nil {
		fmt.Println("Usage:", cmd.usage)
		return err
	}

	return nil
}

// Display Task Tracker CLI commands.
func (c *CLI) DisplayCommands() {
	keys := make([]string, 0, len(c.commands))
	for k := range c.commands {
		keys = append(keys, k)
	}
	slices.Sort(keys)

	fmt.Println(Bold+Blue+"Task Tracker"+Reset, Bold+c.version+Reset)
	fmt.Printf("%sUsage%s: task-cli <command> [<args>]\n", Bold, Reset)
	fmt.Println()
	fmt.Println(Bold + "Available commands:" + Reset)
	fmt.Println()

	maxLen := 0
	for _, k := range keys {
		coloredKey := Bold + Blue + k + Reset
		if len(coloredKey) > maxLen {
			maxLen = len(coloredKey)
		}
	}

	for _, k := range keys {
		cmd := c.commands[k]
		fmt.Printf("   %-*s   %s\n", maxLen, Bold+Blue+cmd.name+Reset, cmd.description)
	}

	fmt.Println()
	fmt.Println("See 'task-cli help <command>' for more information on a specific command.")
}
