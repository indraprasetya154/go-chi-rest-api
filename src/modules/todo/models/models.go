package models

import "time"

type Todo struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Status      bool      `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
}

var todos = []*Todo{
	{
		ID:          "1",
		Title:       "Setup Project",
		Description: "Install dependencies before we run the project",
		Status:      false,
		CreatedAt:   mustParseTime("2023-11-01T12:00:00Z"),
	},
}

func ListTodos() []*Todo {
	return todos
}

func mustParseTime(dateStr string) time.Time {
	t, _ := time.Parse(time.RFC3339, dateStr)
	return t
}
