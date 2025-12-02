# Task Tracker

![Test Status](https://github.com/duanechan/task-tracker/actions/workflows/ci.yml/badge.svg)
![Go Version](https://img.shields.io/badge/Go-1.25%2B-blue?logo=go&logoColor=white)
![Binary Size](https://img.shields.io/badge/Binary-3MB-brightgreen)

A simple and lightweight CLI task manager for adding, listing, updating, and deleting tasks.

This project was built as part of the **Backend Developer Roadmap** projects from [roadmap.sh](https://roadmap.sh).

**Project URL**: [https://roadmap.sh/projects/task-tracker](https://roadmap.sh/projects/task-tracker)

## Installation

1. Install [Go 1.25.3+](https://go.dev/dl/) or up and verify the installation:

```bash
go version
```

2. Install the Task Tracker CLI binary:

```bash
go install github.com/duanechan/task-tracker/cmd/task-cli@latest
```

3. Ensure `$GOPATH/bin` (usually `$HOME/go/bin`) is in your `PATH`:

```bash
export PATH=$PATH:$HOME/go/bin
```

4. Verify installation by running:

```bash
task-cli version
```

**You're all set.** 🎉

## Commands

Quick overview of main commands:

| Command               | Description                          |
|-----------------------|--------------------------------------|
| **add**               | Add a new task                       |
| **delete**            | Delete a task by ID                  |
| **list**              | List all tasks or filter by status   |
| **mark-done**         | Mark a task as done                  |
| **mark-in-progress**  | Mark a task as in-progress           |
| **update**            | Update a task's description          |

For full usage and options, run:

```bash
task-cli help
```

or for a specific command:

```bash
task-cli help <command>
```

## Usage Examples

### Add a New Task

```bash
task-cli add "Buy groceries"
# Output:
# Task added successfully: (ID: 1) Buy groceries
```

### List All Tasks

```bash
task-cli list
# Output:
# 1. [TODO] ID:1 - Buy groceries
```

### List Task by Status

```bash
task-cli list todo
# Output:
# 1. [TODO] ID:1 - Buy groceries
```
### Mark a Task as Done

```bash
task-cli mark-done 1
# Output:
# Task (ID: 1) Buy groceries status updated to: Done
```

### Update a Task Description

```bash
task-cli update 1 "Buy groceries and milk"
# Output:
# Updated Task (ID: 1) description to "Buy groceries and milk"
```

### Delete a Task

```bash
task-cli delete 1
# Output:
# Deleted Task: (ID: 1) Buy groceries and milk
```

## License

This project is licensed under the [MIT License](LICENSE).

## Acknowledgements

Built as part of the **Backend Developer Roadmap** projects: [roadmap.sh](https://roadmap.sh/projects/task-tracker)
