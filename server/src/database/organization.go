package database

import (
	"menagment-app-2/src/model"
	model_organization "menagment-app-2/src/model/organization"
)

func (p *Postgres) CreateOrganization(req model_organization.CreateOrganizationRequest) (*model.CreateElementResponse, error) {
	organization := model_organization.Organization{
		Name: req.Name,
	}
	res := p.db.Create(&organization)
	if res.Error != nil {
		return nil, res.Error
	}
	return &model.CreateElementResponse{Id: organization.Id}, nil
}
