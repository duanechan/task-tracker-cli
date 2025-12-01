package task

import (
	"errors"
	"os"
	"path/filepath"
	"strconv"
	"testing"
	"time"
)

func TestCommandMarkDone(t *testing.T) {
	t.Cleanup(func() {
		path := filepath.Join(cwd, filename)
		os.Remove(path)
	})

	tests := []struct {
		name       string
		args       []string
		wantErr    error
		wantStatus map[int]Status
	}{
		{"mark existing", []string{"1"}, nil, map[int]Status{1: Done, 2: InProgress}},
		{"task not found", []string{"99"}, ErrTaskNotFound, map[int]Status{1: Todo, 2: InProgress}},
		{"missing arg", []string{}, ErrMissingArg, map[int]Status{1: Todo, 2: InProgress}},
		{"empty arg", []string{"   "}, ErrEmptyArgs, map[int]Status{1: Todo, 2: InProgress}},
		{"invalid id", []string{"abc"}, strconv.ErrSyntax, map[int]Status{1: Todo, 2: InProgress}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			now := time.Now()
			state := &state{
				Tasks: []Task{
					{ID: 1, Description: "Task 1", Status: Todo, UpdatedAt: now},
					{ID: 2, Description: "Task 2", Status: InProgress, UpdatedAt: now},
				},
			}

			err := commandMarkDone(state, tt.args)

			if tt.wantErr != nil {
				if tt.wantErr == strconv.ErrSyntax {
					var syntaxErr *strconv.NumError
					if !errors.As(err, &syntaxErr) {
						t.Errorf("expected strconv.NumError, got %v", err)
					}
				} else if err != tt.wantErr {
					t.Errorf("expected error %v, got %v", tt.wantErr, err)
				}
			} else if err != nil {
				t.Errorf("unexpected error: %v", err)
			}

			for id, status := range tt.wantStatus {
				for _, task := range state.Tasks {
					if task.ID == id && task.Status != status {
						t.Errorf("task ID %d: expected status %v, got %v", id, status, task.Status)
					}
				}
			}
		})
	}
}
