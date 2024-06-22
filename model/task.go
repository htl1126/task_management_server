package model

type TaskStatus int

const (
	Completed   TaskStatus = 0
	Incompleted TaskStatus = 1
)

type TaskItem struct {
	ID     int        `json:"id"`
	Name   string     `json:"name"`
	Status TaskStatus `json:"status"`
}
