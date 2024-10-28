package database

import model_operations "menagment-app-2/src/model/operations"

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
