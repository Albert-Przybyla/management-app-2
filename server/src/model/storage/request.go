package model_storage

type CreateStorageRequest struct {
	Name string `json:"name" binding:"required,min=3,max=255"`
}
