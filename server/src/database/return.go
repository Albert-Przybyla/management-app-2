package database

import (
	"menagment-app-2/src/model"
	model_operations "menagment-app-2/src/model/operations"
)

func (p *Postgres) CreateReturn(req model_operations.CreateRetrunRequest, organizationID string) error {
	returnItem := model_operations.Return{
		SourceUserID:         req.SourceUserID,
		DestinationStorageID: req.DestinationStorageID,
		Notes:                req.Notes,
		OrganizationID:       organizationID,
	}

	res := p.db.Create(&returnItem)
	if res.Error != nil {
		return res.Error
	}

	var returnItems []model_operations.ReturnItem
	for _, itemID := range req.Items {
		returnItems = append(returnItems, model_operations.ReturnItem{
			ReturnID: returnItem.ID,
			ItemID:   itemID,
		})
	}
	var returnPackages []model_operations.ReturnPackage
	for _, packageID := range req.Packages {
		returnPackages = append(returnPackages, model_operations.ReturnPackage{
			ReturnID:  returnItem.ID,
			PackageID: packageID,
		})
	}

	if len(returnItems) > 0 {
		res = p.db.Create(&returnItems)
		if res.Error != nil {
			return res.Error
		}
	}

	if len(returnPackages) > 0 {
		res = p.db.Create(&returnPackages)
		if res.Error != nil {
			return res.Error
		}
	}
	return nil
}

func (p *Postgres) GetRetruns(organizationID string, pageSize, pageNumber int) (*model.PagedListResponse[model_operations.Return], error) {
	var items []model_operations.Return
	res := p.db.Preload("Items").Preload("Packages").Where("organization_id = ?", organizationID).Order("created_at desc").Find(&items)

	if res.Error != nil {
		return nil, res.Error
	}
	var totalItems int64
	p.db.Model(&model_operations.Return{}).Where("organization_id = ?", organizationID).Count(&totalItems)

	totalPages := int((totalItems + int64(pageSize) - 1) / int64(pageSize))

	response := model.PagedListResponse[model_operations.Return]{
		Items:       items,
		TotalItems:  int(totalItems),
		TotalPages:  totalPages,
		CurrentPage: pageNumber,
	}

	return &response, nil
}

func (p *Postgres) GetAllRetrurns(organizationID string) (*[]model_operations.Return, error) {
	var items []model_operations.Return
	res := p.db.Preload("Items").Preload("Packages").Where("organization_id = ?", organizationID).Order("created_at desc").Find(&items)

	if res.Error != nil {
		return nil, res.Error
	}
	p.db.Model(&model_operations.Return{}).Where("organization_id = ?", organizationID)
	return &items, nil
}
