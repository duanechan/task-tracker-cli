package task

import (
	"fmt"
	"time"
)

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

func (t Task) String() string {
	return fmt.Sprintf("(ID: %d) %s", t.ID, t.Description)
}
