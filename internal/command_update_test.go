package task

import (
	"errors"
	"os"
	"path/filepath"
	"strconv"
	"testing"
)

func TestCommandUpdate(t *testing.T) {
	t.Cleanup(func() {
		path := filepath.Join(cwd, filename)
		os.Remove(path)
	})

	mockCLI := func() *CLI {
		return &CLI{
			state: &state{
				NextID: 3,
				Tasks: []Task{
					{ID: 1, Description: "Task 1", Status: Todo},
					{ID: 2, Description: "Task 2", Status: Todo},
					{ID: 3, Description: "Task 3", Status: Todo},
				},
			},
		}
	}

	tests := []struct {
		name           string
		args           []string
		wantErr        error
		wantDesc       map[int]string
		checkUpdatedAt bool
	}{
		{
			name:    "normal update",
			args:    []string{"2", "New Task 2"},
			wantErr: nil,
			wantDesc: map[int]string{
				1: "Task 1",
				2: "New Task 2",
				3: "Task 3",
			},
			checkUpdatedAt: true,
		},
		{
			name:    "task not found",
			args:    []string{"99", "Does not exist"},
			wantErr: ErrTaskNotFound,
			wantDesc: map[int]string{
				1: "Task 1",
				2: "Task 2",
				3: "Task 3",
			},
		},
		{
			name:    "missing argument",
			args:    []string{"1"},
			wantErr: ErrMissingArg,
		},
		{
			name:    "too many arguments",
			args:    []string{"1", "desc", "extra"},
			wantErr: ErrTooManyArgs,
		},
		{
			name:    "empty id",
			args:    []string{"   ", "desc"},
			wantErr: ErrEmptyArgs,
		},
		{
			name:    "empty description",
			args:    []string{"1", "   "},
			wantErr: ErrEmptyArgs,
		},
		{
			name:    "invalid id",
			args:    []string{"abc", "desc"},
			wantErr: strconv.ErrSyntax,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := mockCLI()
			oldUpdatedAt := c.state.Tasks[1].UpdatedAt

			err := commandUpdate(c, tt.args)

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

			if tt.wantDesc != nil {
				for _, task := range c.state.Tasks {
					wantDesc, ok := tt.wantDesc[task.ID]
					if !ok {
						continue
					}
					if task.Description != wantDesc {
						t.Errorf("task ID %d: expected description %q, got %q", task.ID, wantDesc, task.Description)
					}

					if tt.checkUpdatedAt && task.ID == 2 && !task.UpdatedAt.After(oldUpdatedAt) {
						t.Errorf("task ID %d: UpdatedAt not updated", task.ID)
					}
				}
			}
		})
	}
}
