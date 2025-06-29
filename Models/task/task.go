package task

type Task struct {
	TaskID     int    `json:"taskid"`
	TaskName   string `json:"taskname"`
	TaskStatus string `json:"status"`
	AssignUser int    `json:"assigned_user_id"`
}
