package task

import (
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
