package model_operations

type CreateIssueRequest struct {
	SourceStorageID   string   `json:"source_storage_id"`
	DestinationUserID string   `json:"destination_user_id"`
	Notes             *string  `json:"notes"`
	Items             []string `json:"items"`
	Packages          []string `json:"packages"`
}
