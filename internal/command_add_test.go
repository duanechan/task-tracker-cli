package task

import (
	"os"
	"path/filepath"
	"testing"
)

func TestCommandAdd(t *testing.T) {
	t.Cleanup(func() {
		path := filepath.Join(cwd, filename)
		os.Remove(path)
	})

	tests := []struct {
		name      string
		argsList  [][]string
		wantTasks []Task
		wantErr   []error
	}{
		{
			name: "single add",
			argsList: [][]string{
				{"Learn Go"},
			},
			wantTasks: []Task{
				{ID: 1, Description: "Learn Go", Status: Todo},
			},
			wantErr: []error{nil},
		},
		{
			name: "multiple sequential adds",
			argsList: [][]string{
				{"eat "},
				{"  sleep  "},
				{"repeat"},
			},
			wantTasks: []Task{
				{ID: 1, Description: "eat", Status: Todo},
				{ID: 2, Description: "sleep", Status: Todo},
				{ID: 3, Description: "repeat", Status: Todo},
			},
			wantErr: []error{nil, nil, nil},
		},
		{
			name: "invalid adds",
			argsList: [][]string{
				{},
				{"   "},
				{"Hello", "world"},
			},
			wantTasks: []Task{},
			wantErr:   []error{ErrMissingArg, ErrEmptyArgs, ErrTooManyArgs},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			state := &state{NextID: 0, Tasks: []Task{}}

			for i, args := range tt.argsList {
				err := commandAdd(state, args)
				if err != tt.wantErr[i] {
					t.Errorf("expected error %v, got %v", tt.wantErr, err)
				}
			}

			if len(state.Tasks) != len(tt.wantTasks) {
				t.Fatalf("expected %d tasks, got %d", len(tt.wantTasks), len(state.Tasks))
			}

			for i, want := range tt.wantTasks {
				actual := state.Tasks[i]
				checkTaskIfEqual(t, want, actual)
			}
		})
	}

}

func checkTaskIfEqual(t *testing.T, expected Task, actual Task) {
	if expected.ID != actual.ID {
		t.Errorf("expected id to be %d, got %d", expected.ID, actual.ID)
	}

	if expected.Description != actual.Description {
		t.Errorf("expected description to be %s, got %s", expected.Description, actual.Description)
	}

	if expected.Status != actual.Status {
		t.Errorf("expected status to be %d, got %d", expected.Status, actual.Status)
	}
}
