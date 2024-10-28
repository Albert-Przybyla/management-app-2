package model_operations

type CreateRetrunRequest struct {
	SourceUserID         string   `json:"source_user_id"`
	DestinationStorageID string   `json:"destination_storage_id"`
	Notes                *string  `json:"notes"`
	Items                []string `json:"items"`
	Packages             []string `json:"packages"`
}
