package model

type CreateReq struct {
	Name string `json:"name" binding:"required"`
}
