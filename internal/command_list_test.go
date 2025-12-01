package task

import (
	"bytes"
	"io"
	"os"
	"strings"
	"testing"
)

func TestCommandList(t *testing.T) {
	state := &state{
		Tasks: []Task{
			{ID: 1, Description: "Task 1", Status: Todo},
			{ID: 2, Description: "Task 2", Status: InProgress},
			{ID: 3, Description: "Task 3", Status: Done},
		},
	}

	tests := []struct {
		name    string
		args    []string
		wantErr error
		wantOut []string
	}{
		{"all tasks", []string{}, nil, []string{"Task 1", "Task 2", "Task 3"}},
		{"todo tasks", []string{"todo"}, nil, []string{"Task 1"}},
		{"done tasks", []string{"done"}, nil, []string{"Task 3"}},
		{"invalid arg", []string{"foo"}, ErrInvalidArg, nil},
		{"empty arg", []string{"   "}, ErrEmptyArgs, nil},
		{"too many args", []string{"todo", "extra"}, ErrTooManyArgs, nil},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf bytes.Buffer

			old := os.Stdout

			r, w, _ := os.Pipe()

			os.Stdout = w

			_ = commandList(state, tt.args)

			w.Close()

			os.Stdout = old

			io.Copy(&buf, r)

			out := buf.String()

			err := commandList(state, tt.args)

			if err != tt.wantErr {
				t.Errorf("expected error %v, got %v", tt.wantErr, err)
			}

			if tt.wantOut != nil {
				for _, s := range tt.wantOut {
					if !strings.Contains(out, s) {
						t.Errorf("expected output to contain %q, got %q", s, out)
					}
				}
			}
		})
	}
}
