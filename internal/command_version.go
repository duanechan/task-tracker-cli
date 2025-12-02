package task

import "fmt"

func commandVersion(c *CLI, args []string) error {
	if len(args) > 0 {
		return ErrTooManyArgs
	}
	fmt.Println(c.version)
	return nil
}