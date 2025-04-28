package task

import (
	"fmt"
	"time"
)

type Task struct {
	ID          uint      `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Priority    string    `json:"priority"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

var ValidPriorities = []string{"low", "medium", "high"}

func (t *Task) Validate() error {
	if t.Title == "" {
		return fmt.Errorf("title is required")
	}

	validPriority := false
	for _, p := range ValidPriorities {
		if t.Priority == p {
			validPriority = true
			break
		}
	}
	if !validPriority {
		return fmt.Errorf("invalid priority")
	}

	if t.CreatedAt.IsZero() {
		return fmt.Errorf("created at is required")
	}

	if t.UpdatedAt.IsZero() {
		return fmt.Errorf("updated at is required")
	}

	if t.UpdatedAt.Before(t.CreatedAt) {
		return fmt.Errorf("updated at cannot be before created at")
	}

	now := time.Now()
	if t.CreatedAt.After(now) {
		return fmt.Errorf("created at cannot be in the future")
	}
	if t.UpdatedAt.After(now) {
		return fmt.Errorf("updated at cannot be in the future")
	}

	return nil
}
