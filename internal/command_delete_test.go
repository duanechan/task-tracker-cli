package task

import (
	"bytes"
	"errors"
	"os"
	"strconv"
	"strings"
	"testing"
)

func captureStdout(f func()) string {
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	f()

	w.Close()
	os.Stdout = old
	var buf bytes.Buffer
	buf.ReadFrom(r)
	return buf.String()
}

func TestCommandDelete(t *testing.T) {
	s := &state{
		Tasks: []Task{
			{ID: 1, Description: "Task One"},
			{ID: 2, Description: "Task Two"},
		},
	}

	err := commandDelete(s, []string{"1"})
	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}

	if len(s.Tasks) != 1 {
		t.Fatalf("expected 1 task left, got %d", len(s.Tasks))
	}

	if s.Tasks[0].ID != 2 {
		t.Fatalf("expected remaining task ID 2, got %d", s.Tasks[0].ID)
	}
}

func TestCommandDeleteError(t *testing.T) {
	tests := []struct {
		name string
		args []string
		want error
	}{
		{
			name: "missing arg",
			args: []string{},
			want: ErrMissingArg,
		},
		{
			name: "too many args",
			args: []string{"1", "2"},
			want: ErrTooManyArgs,
		},
		{
			name: "empty id",
			args: []string{" "},
			want: ErrEmptyArgs,
		},
		{
			name: "invalid id",
			args: []string{"abc"},
			want: strconv.ErrSyntax,
		},
		{
			name: "not found",
			args: []string{"99"},
			want: ErrTaskNotFound,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			s := &state{
				Tasks: []Task{
					{ID: 1, Description: "Hello"},
				},
			}

			err := commandDelete(s, tc.args)

			if err == nil {
				t.Fatalf("expected error %v, got nil", tc.want)
			}

			if errors.Is(tc.want, strconv.ErrSyntax) {
				if !strings.Contains(err.Error(), "invalid syntax") {
					t.Fatalf("expected Atoi syntax error, got %v", err)
				}
				return
			}

			if !errors.Is(err, tc.want) {
				t.Fatalf("expected error %v, got %v", tc.want, err)
			}
		})
	}
}

func TestCommandDeleteStdout(t *testing.T) {
	s := &state{
		Tasks: []Task{
			{ID: 10, Description: "Clean room"},
			{ID: 20, Description: "Study"},
		},
	}

	output := captureStdout(func() {
		err := commandDelete(s, []string{"10"})
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
	})

	expected := "Deleted Task (ID: 10) Clean room"
	if !strings.Contains(output, expected) {
		t.Fatalf("expected stdout to contain %q, got %q", expected, output)
	}
}
