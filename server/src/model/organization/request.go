package model_organization

type CreateOrganizationRequest struct {
	Name string `json:"name" binding:"required,min=3,max=255"`
}
