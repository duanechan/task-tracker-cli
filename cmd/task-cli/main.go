package main

import (
	"fmt"
	"os"

	task "github.com/duanechan/task-tracker/internal"
)

func main() {
	cli, err := task.LoadCLI()
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	if len(os.Args) < 2 {
		cli.DisplayCommands()
		os.Exit(1)
	}

	if err = cli.Run(os.Args[1:]); err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}
