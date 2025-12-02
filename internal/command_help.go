package task

import (
	"fmt"
	"strings"
)

func commandHelp(c *CLI, args []string) error {
	if len(args) > 1 {
		return ErrTooManyArgs
	}

	if len(args) == 0 {
		c.DisplayCommands()
		return nil
	}

	name := strings.TrimSpace(args[0])
	command, exists := c.commands[name]
	if !exists {
		return fmt.Errorf("command '%s' does not exist", name)
	}

	fmt.Println("Task Tracker", c.version)
	fmt.Println()
	fmt.Printf("%s: %s\n", name, command.description)
	fmt.Println("Usage: task-cli", command.usage)

	return nil
}
