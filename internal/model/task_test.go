package model

import (
	"testing"
	"time"
)

func TestTask_IsOverdue(t *testing.T) {
	tests := []Task{
		{
			ID:        0,
			Title:     "due date is nil",
			Body:      "due date is nil",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			Done:      false,
			DueDate:   nil,
		},
		{
			ID:        1,
			Title:     "due date is in the past",
			Body:      "due date is in the past",
			CreatedAt: time.Now().Add(-48 * time.Hour),
			UpdatedAt: time.Now().Add(-48 * time.Hour),
			Done:      true,
			DueDate: func() *time.Time {
				t := time.Now().Add(-24 * time.Hour)
				return &t
			}(),
		},
		{
			ID:        2,
			Title:     "due date is now",
			Body:      "due date is now",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			Done:      false,
			DueDate: func() *time.Time {
				t := time.Now()
				return &t
			}(),
		},
		{
			ID:        3,
			Title:     "due date is in the future",
			Body:      "due date is in the future",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			Done:      false,
			DueDate: func() *time.Time {
				t := time.Now().Add(24 * time.Hour)
				return &t
			}(),
		},
	}

	for _, tt := range tests {
		t.Run(tt.Title, func(t *testing.T) {
			if got := tt.IsOverdue(); got != tt.Done {
				t.Errorf("IsOverdue() = %v, want %v", got, tt.Done)
			}
		})
	}
}
