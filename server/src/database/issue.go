package database

import (
	"menagment-app-2/src/model"
	model_operations "menagment-app-2/src/model/operations"
)

func (p *Postgres) CreateIssue(req model_operations.CreateIssueRequest, organizationID string) error {
	issue := model_operations.Issue{
		SourceStorageID:   req.SourceStorageID,
		DestinationUserID: req.DestinationUserID,
		Notes:             req.Notes,
		OrganizationID:    organizationID,
	}

	res := p.db.Create(&issue)
	if res.Error != nil {
		return res.Error
	}

	var issueItems []model_operations.IssueItem
	for _, itemID := range req.Items {
		issueItems = append(issueItems, model_operations.IssueItem{
			IssueID: issue.ID,
			ItemID:  itemID,
		})
	}
	var issuePackages []model_operations.IssuePackage
	for _, packageID := range req.Packages {
		issuePackages = append(issuePackages, model_operations.IssuePackage{
			IssueID:   issue.ID,
			PackageID: packageID,
		})
	}

	if len(issueItems) > 0 {
		res = p.db.Create(&issueItems)
		if res.Error != nil {
			return res.Error
		}
	}

	if len(issuePackages) > 0 {
		res = p.db.Create(&issuePackages)
		if res.Error != nil {
			return res.Error
		}
	}
	return nil
}

func (p *Postgres) GetIssues(organizationID string, pageSize, pageNumber int) (*model.PagedListResponse[model_operations.Issue], error) {
	var items []model_operations.Issue
	res := p.db.Preload("Items").Preload("Packages").Where("organization_id = ?", organizationID).Order("created_at desc").Find(&items)

	if res.Error != nil {
		return nil, res.Error
	}
	var totalItems int64
	p.db.Model(&model_operations.Issue{}).Where("organization_id = ?", organizationID).Count(&totalItems)

	totalPages := int((totalItems + int64(pageSize) - 1) / int64(pageSize))

	response := model.PagedListResponse[model_operations.Issue]{
		Items:       items,
		TotalItems:  int(totalItems),
		TotalPages:  totalPages,
		CurrentPage: pageNumber,
	}

	return &response, nil
}

func (p *Postgres) GetAllIssues(organizationID string) (*[]model_operations.Issue, error) {
	var items []model_operations.Issue
	res := p.db.Preload("Items").Preload("Packages").Where("organization_id = ?", organizationID).Order("created_at desc").Find(&items)

	if res.Error != nil {
		return nil, res.Error
	}
	p.db.Model(&model_operations.Issue{}).Where("organization_id = ?", organizationID)
	return &items, nil
}
