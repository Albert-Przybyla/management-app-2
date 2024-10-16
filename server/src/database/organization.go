package database

import (
	"menagment-app-2/src/model"
	model_organization "menagment-app-2/src/model/organization"

	"gorm.io/gorm"
)

func (p *Postgres) CreateOrganization(req model_organization.CreateOrganizationRequest) (*model.CreateElementResponse, error) {
	organization := model_organization.Organization{
		Name: req.Name,
	}
	res := p.db.Create(&organization)
	if res.Error != nil {
		return nil, res.Error
	}
	return &model.CreateElementResponse{Id: organization.ID}, nil
}

func (p *Postgres) GetOrganizationById(id string) (*model_organization.Organization, error) {
	var organization model_organization.Organization
	res := p.db.Table("organizations").
		Preload("Users").
		Preload("Storages").
		Where("id = ?", id).
		First(&organization)
	if res.Error != nil {
		if res.Error == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, res.Error
	}
	return &organization, nil
}
