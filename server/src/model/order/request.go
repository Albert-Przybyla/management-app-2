package model_order

import "time"

type CreateOrderRequest struct {
	Quantity       int       `json:"quantity"`
	OrganizationID string    `json:"organization_id"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}
