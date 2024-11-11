package model_customer

type CreateCustomerRequest struct {
	Name           string `json:"name"`
	OrganizationID string `json:"organization_id"`
}
