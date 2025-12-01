package task

import (
	"errors"
	"os"
	"path/filepath"
	"strconv"
	"testing"
)

func TestCommandDelete(t *testing.T) {
	t.Cleanup(func() {
		path := filepath.Join(cwd, filename)
		os.Remove(path)
	})

	tests := []struct {
		name      string
		args      []string
		wantErr   error
		wantTasks []Task
	}{
		{
			name:    "delete existing task",
			args:    []string{"2"},
			wantErr: nil,
			wantTasks: []Task{
				{ID: 1, Description: "Task 1", Status: Todo},
				{ID: 3, Description: "Task 3", Status: Todo},
			},
		},
		{
			name:      "task not found",
			args:      []string{"99"},
			wantErr:   ErrTaskNotFound,
			wantTasks: []Task{{ID: 1, Description: "Task 1", Status: Todo}, {ID: 2, Description: "Task 2", Status: Todo}, {ID: 3, Description: "Task 3", Status: Todo}},
		},
		{
			name:      "missing argument",
			args:      []string{},
			wantErr:   ErrMissingArg,
			wantTasks: nil,
		},
		{
			name:      "too many arguments",
			args:      []string{"1", "2"},
			wantErr:   ErrTooManyArgs,
			wantTasks: nil,
		},
		{
			name:      "empty argument",
			args:      []string{"   "},
			wantErr:   ErrEmptyArgs,
			wantTasks: nil,
		},
		{
			name:      "invalid ID",
			args:      []string{"abc"},
			wantErr:   strconv.ErrSyntax,
			wantTasks: nil,
		},
	}

	makeState := func() *state {
		return &state{
			NextID: 3,
			Tasks: []Task{
				{ID: 1, Description: "Task 1", Status: Todo},
				{ID: 2, Description: "Task 2", Status: Todo},
				{ID: 3, Description: "Task 3", Status: Todo},
			},
		}
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			state := makeState()

			err := commandDelete(state, tt.args)

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

			if tt.wantTasks != nil {
				if len(state.Tasks) != len(tt.wantTasks) {
					t.Fatalf("expected %d tasks, got %d", len(tt.wantTasks), len(state.Tasks))
				}

				for i, want := range tt.wantTasks {
					actual := state.Tasks[i]
					checkTaskIfEqual(t, want, actual)
				}
			}
		})
	}
}
