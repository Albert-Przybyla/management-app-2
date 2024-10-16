package database

import (
	"menagment-app-2/src/model"
	model_item "menagment-app-2/src/model/item"

	"gorm.io/gorm"
)

func (p *Postgres) CreateItem(req model_item.CreateItemRequest, organizationID string) error {
	item := model_item.Item{
		Name:           req.Name,
		Description:    req.Description,
		Code:           req.Code,
		Quantity:       req.Quantity,
		OrganizationID: organizationID,
	}
	res := p.db.Create(&item)
	if res.Error != nil {
		return res.Error
	}
	return nil
}

func (p *Postgres) GetItems(organizationID string, pageSize, pageNumber int) (model.PagedListResponse[model_item.Item], error) {

	var items []model_item.Item
	res := p.db.Where("organization_id = ?", organizationID).Limit(pageSize).Offset((pageNumber - 1) * pageSize).Find(&items)
	if res.Error != nil {
		return model.PagedListResponse[model_item.Item]{}, res.Error
	}
	var totalItems int64
	p.db.Model(&model_item.Item{}).Where("organization_id = ?", organizationID).Count(&totalItems)

	totalPages := int((totalItems + int64(pageSize) - 1) / int64(pageSize))

	response := model.PagedListResponse[model_item.Item]{
		Items:       items,
		TotalItems:  int(totalItems),
		TotalPages:  totalPages,
		CurrentPage: pageNumber,
	}

	return response, nil
}

func (p *Postgres) GetItemById(id string, organizationID string) (*model_item.Item, error) {
	var item model_item.Item
	res := p.db.Where("id = ? AND organization_id = ?", id, organizationID).First(&item)
	if res.Error != nil {
		if res.Error == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, res.Error
	}
	return &item, nil
}

func (p *Postgres) UpdateItem(id string, organizationID string, req model_item.UpdateItemRequest) error {
	res := p.db.Model(&model_item.Item{}).Where("id = ? AND organization_id = ?", id, organizationID).Updates(req)
	if res.Error != nil {
		return res.Error
	}
	return nil
}

func (p *Postgres) DeleteItem(id string, organizationID string) error {
	res := p.db.Where("id = ? AND organization_id = ?", id, organizationID).Delete(&model_item.Item{})
	if res.Error != nil {
		return res.Error
	}
	return nil
}

func (p *Postgres) AssignItemToPackage(id string, organizationID string, req model_item.AssignItemToPackageRequest) error {
	res := p.db.Model(&model_item.Item{}).Where("id = ? AND organization_id = ?", id, organizationID).Updates(req)
	if res.Error != nil {
		return res.Error
	}
	return nil
}
