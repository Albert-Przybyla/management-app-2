package database

import (
	"menagment-app-2/src/model"
	model_package "menagment-app-2/src/model/package"

	"gorm.io/gorm"
)

func (p *Postgres) CreatePackage(req model_package.CreatePackageRequest, organizationID string) error {
	pkg := model_package.Package{
		Name:           req.Name,
		OrganizationID: organizationID,
	}
	res := p.db.Create(&pkg)
	if res.Error != nil {
		return res.Error
	}
	return nil
}

func (p *Postgres) GetPackages(organizationID string, pageSize, pageNumber int) (model.PagedListResponse[model_package.Package], error) {

	var items []model_package.Package
	res := p.db.Where("organization_id = ?", organizationID).Limit(pageSize).Offset((pageNumber - 1) * pageSize).Find(&items)
	if res.Error != nil {
		return model.PagedListResponse[model_package.Package]{}, res.Error
	}
	var totalItems int64
	p.db.Model(&model_package.Package{}).Where("organization_id = ?", organizationID).Count(&totalItems)

	totalPages := int((totalItems + int64(pageSize) - 1) / int64(pageSize))

	response := model.PagedListResponse[model_package.Package]{
		Items:       items,
		TotalItems:  int(totalItems),
		TotalPages:  totalPages,
		CurrentPage: pageNumber,
	}
	return response, nil
}

func (p *Postgres) GetPackage(id string, organizationID string) (*model_package.Package, error) {
	var pkg model_package.Package
	res := p.db.Where("id = ? AND organization_id = ?", id, organizationID).First(&pkg)
	if res.Error != nil {
		if res.Error == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, res.Error
	}
	return &pkg, nil
}

func (p *Postgres) DeletePackage(id string, organizationID string) error {
	res := p.db.Where("id = ? AND organization_id = ?", id, organizationID).Delete(&model_package.Package{})
	if res.Error != nil {
		return res.Error
	}
	return nil
}

func (p *Postgres) UpdatePackage(id string, organizationID string, req model_package.UpdatePackageRequest) error {
	res := p.db.Model(&model_package.Package{}).Where("id = ? AND organization_id = ?", id, organizationID).Updates(req)
	if res.Error != nil {
		return res.Error
	}
	return nil
}
