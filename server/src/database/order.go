package database

import (
	"menagment-app-2/src/model"
	model_order "menagment-app-2/src/model/order"

	"gorm.io/gorm"
)

func (p *Postgres) CreateOrder(req model_order.CreateOrderRequest, organizationID string) error {
	pkg := model_order.Order{
		Quantity:       req.Quantity,
		OrganizationID: organizationID,
	}
	res := p.db.Create(&pkg)
	if res.Error != nil {
		return res.Error
	}
	return nil
}

func (p *Postgres) GetOrders(organizationID string, pageSize, pageNumber int) (model.PagedListResponse[model_order.Order], error) {

	var items []model_order.Order
	res := p.db.Where("organization_id = ?", organizationID).Limit(pageSize).Offset((pageNumber - 1) * pageSize).Find(&items)
	if res.Error != nil {
		return model.PagedListResponse[model_order.Order]{}, res.Error
	}
	var totalItems int64
	p.db.Model(&model_order.Order{}).Where("organization_id = ?", organizationID).Count(&totalItems)

	totalPages := int((totalItems + int64(pageSize) - 1) / int64(pageSize))

	response := model.PagedListResponse[model_order.Order]{
		Items:       items,
		TotalItems:  int(totalItems),
		TotalPages:  totalPages,
		CurrentPage: pageNumber,
	}
	return response, nil
}

func (p *Postgres) GetOrder(id string, organizationID string) (*model_order.Order, error) {
	var customer model_order.Order
	res := p.db.Where("id = ? AND organization_id = ?", id, organizationID).First(&customer)
	if res.Error != nil {
		if res.Error == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, res.Error
	}
	return &customer, nil
}

func (p *Postgres) DeleteOrder(id string, organizationID string) error {
	res := p.db.Where("id = ? AND organization_id = ?", id, organizationID).Delete(&model_order.Order{})
	if res.Error != nil {
		return res.Error
	}
	return nil
}
