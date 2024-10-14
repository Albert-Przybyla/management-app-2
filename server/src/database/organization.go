package database

import model_organization "menagment-app-2/src/model/organization"

func (p *Postgres) CreateOrganization(req model_organization.CreateOrganizationRequest) error {
	organization := model_organization.Organization{
		Name: req.Name,
	}
	res := p.db.Create(&organization)
	if res.Error != nil {
		return res.Error
	}
	return nil
}
