package models

type Task struct {
	ID          string
	Title       string
	Description string
	IsComplete  bool
	UserID      string
}
