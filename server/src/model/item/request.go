package model_item

type CreateItemRequest struct {
	Name        string  `json:"name" binding:"required"`
	Description *string `json:"description"`
	Code        *string `json:"code"`
	Quantity    int     `json:"quantity" binding:"required"`
}

type CreateItemMultipleRequest struct {
	Name         string  `json:"name" binding:"required"`
	Description  *string `json:"description"`
	Code         *string `json:"code"`
	Quantity     int     `json:"quantity" binding:"required"`
	Multiplation int     `json:"multiple" binding:"required"`
}

type UpdateItemRequest struct {
	Name        string  `json:"name" binding:"required"`
	Description *string `json:"description"`
	Code        *string `json:"code"`
	Quantity    int     `json:"quantity" binding:"required"`
}

type AssignItemToPackageRequest struct {
	PackageID string `json:"package_id" binding:"required"`
}
