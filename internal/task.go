package task

import "time"

type Status int

const (
	Todo Status = iota
	InProgress
	Done
)

type Task struct {
	ID          int       `json:"id"`
	Description string    `json:"description"`
	Status      Status    `json:"status"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

func NewTask(state *state, description string) *Task {
	state.Count++
	t := &Task{
		ID:          state.Count,
		Description: description,
		Status:      Todo,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	state.Tasks = append(state.Tasks, *t)
	return t
}
