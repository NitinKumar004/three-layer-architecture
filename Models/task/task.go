package task

type Task struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Status string `json:"status"`
	UserID int    `json:"user_id"`
}
