package database

import (
	"menagment-app-2/src/model"
	model_operations "menagment-app-2/src/model/operations"
)

func (p *Postgres) CreateTransfer(req model_operations.CreateTransferRequest, organizationID string) error {

	transfer := model_operations.Transfer{
		SourceStorageID:      req.SourceStorageID,
		DestinationStorageID: req.DestinationStorageID,
		Notes:                req.Notes,
		OrganizationID:       organizationID,
	}

	res := p.db.Create(&transfer)
	if res.Error != nil {
		return res.Error
	}

	var transferItems []model_operations.TransferItem
	for _, itemID := range req.Items {
		transferItems = append(transferItems, model_operations.TransferItem{
			TransferID: transfer.ID,
			ItemID:     itemID,
		})
	}
	var transferPackages []model_operations.TransferPackage
	for _, packageID := range req.Packages {
		transferPackages = append(transferPackages, model_operations.TransferPackage{
			TransferID: transfer.ID,
			PackageID:  packageID,
		})
	}

	if len(transferItems) > 0 {
		res = p.db.Create(&transferItems)
		if res.Error != nil {
			return res.Error
		}
	}

	if len(transferPackages) > 0 {
		res = p.db.Create(&transferPackages)
		if res.Error != nil {
			return res.Error
		}
	}
	return nil
}

func (p *Postgres) GetTransfer(transferID, organizationID string) (*model_operations.Transfer, error) {
	var transfer model_operations.Transfer

	res := p.db.Preload("TransferItems").Preload("TransferPackages").First(&transfer, "id = ? AND organization_id = ?", transferID, organizationID)
	if res.Error != nil {
		return nil, res.Error
	}

	return &transfer, nil
}

func (p *Postgres) GetTransfers(organizationID string, pageSize, pageNumber int) (*model.PagedListResponse[model_operations.Transfer], error) {
	var items []model_operations.Transfer
	res := p.db.Preload("Items").Preload("Packages").Where("organization_id = ?", organizationID).Order("created_at desc").Find(&items)

	if res.Error != nil {
		return nil, res.Error
	}
	var totalItems int64
	p.db.Model(&model_operations.Transfer{}).Where("organization_id = ?", organizationID).Count(&totalItems)

	totalPages := int((totalItems + int64(pageSize) - 1) / int64(pageSize))

	response := model.PagedListResponse[model_operations.Transfer]{
		Items:       items,
		TotalItems:  int(totalItems),
		TotalPages:  totalPages,
		CurrentPage: pageNumber,
	}

	return &response, nil
}

func (p *Postgres) GetAllTransfers(organizationID string) (*[]model_operations.Transfer, error) {
	var items []model_operations.Transfer
	res := p.db.Preload("Items").Preload("Packages").Where("organization_id = ?", organizationID).Order("created_at desc").Find(&items)

	if res.Error != nil {
		return nil, res.Error
	}
	p.db.Model(&model_operations.Transfer{}).Where("organization_id = ?", organizationID)
	return &items, nil
}
