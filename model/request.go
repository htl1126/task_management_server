package model

type CreateReq struct {
	Name string `json:"name" binding:"required"`
}

type UpdateReq struct {
	Name   *string     `json:"name"`
	Status *TaskStatus `json:"status"`
}
