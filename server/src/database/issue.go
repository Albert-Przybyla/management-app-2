package database

import model_operations "menagment-app-2/src/model/operations"

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
