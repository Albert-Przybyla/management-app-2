package model_operations

type CreateTransferRequest struct {
	SourceStorageID      *string  `json:"source_storage_id"`
	DestinationStorageID *string  `json:"destination_storage_id"`
	Notes                *string  `json:"notes"`
	Items                []string `json:"items"`
	Packages             []string `json:"packages"`
}
