package model

import (
	"time"
)

type Task struct {
	ID        int        `json:"id"`
	Title     string     `json:"title"`
	Body      string     `json:"body"`
	Tags      []string   `json:"tags"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt"`
	Done      bool       `json:"done"`
	DueDate   *time.Time `json:"dueDate"`
}

func (t Task) IsOverdue() bool {
	if t.DueDate == nil {
		return false
	}

	return time.Now().After(*t.DueDate)
}
