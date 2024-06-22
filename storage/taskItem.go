package storage

import "task-manage-api/model"

func CreateTaskItem(id int, name string, status model.TaskStatus) model.TaskItem {
	return model.TaskItem{
		ID:     id,
		Name:   name,
		Status: status,
	}
}
