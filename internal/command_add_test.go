package task

import (
	"bytes"
	"errors"
	"os"
	"testing"
)

func TestCommandAdd(t *testing.T) {
	testState := state{Tasks: []Task{}}
	testArgs := [][]string{
		{"Make breakfast"},
		{"Exercise for 30-mins"},
		{"Study data structures & algorithms"},
	}

	for i, args := range testArgs {
		t.Run(args[0], func(t *testing.T) {
			if err := commandAdd(&testState, args); err != nil {
				t.Errorf("expected err to be nil, got %s", err)
			}

			task := testState.Tasks[i]

			if task.ID != i+1 {
				t.Errorf("expected id to be %d, got %d", i+1, task.ID)
			}

			if task.Description != testArgs[i][0] {
				t.Errorf("expected description to be %s, got %s", testArgs[i][0], task.Description)
			}

			if task.Status != Todo {
				t.Errorf("expected status to be %d, got %d", Todo, task.Status)
			}
		})
	}

	if testState.Count != len(testArgs) {
		t.Errorf("expected count to be %d, got %d", len(testArgs), testState.Count)
	}

}

func TestCommandAddError(t *testing.T) {
	tests := []struct {
		name string
		args []string
		want error
	}{
		{"empty string", []string{"   "}, ErrEmptyDescription},
		{"too many args", []string{"This", "is", "my", "task"}, ErrTooManyArgs},
		{"missing arg", []string{}, ErrMissingArg},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			testState := state{Tasks: []Task{}}

			err := commandAdd(&testState, tc.args)
			if err == nil {
				t.Fatalf("expected error %q, got nil", tc.want)
			}

			if !errors.Is(err, tc.want) {
				t.Fatalf("expected error %q, got %q", tc.want, err)
			}
		})
	}
}

func TestCommandAddStdout(t *testing.T) {
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	t.Cleanup(func() {
		os.Stdout = old
	})

	testState := state{Tasks: []Task{}}
	testArgs := [][]string{
		{"Make breakfast"},
		{"Exercise for 30-mins"},
		{"Study data structures & algorithms"},
	}

	for _, args := range testArgs {
		if err := commandAdd(&testState, args); err != nil {
			t.Fatalf("expected no error, got %v", err)
		}
	}

	w.Close()

	var buf bytes.Buffer
	if _, err := buf.ReadFrom(r); err != nil {
		t.Fatalf("failed to read stdout: %v", err)
	}
	output := buf.String()

	expected := "Task added successfully (ID: 1)\nTask added successfully (ID: 2)\nTask added successfully (ID: 3)\n"

	if output != expected {
		t.Errorf("expected stdout %q, got %q", expected, output)
	}
}
