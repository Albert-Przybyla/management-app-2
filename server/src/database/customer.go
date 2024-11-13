package database

import (
	"menagment-app-2/src/model"
	model_customer "menagment-app-2/src/model/customer"

	"gorm.io/gorm"
)

func (p *Postgres) CreateCustomer(req model_customer.CreateCustomerRequest, organizationID string) error {
	pkg := model_customer.Customer{
		Name:           req.Name,
		OrganizationID: organizationID,
	}
	res := p.db.Create(&pkg)
	if res.Error != nil {
		return res.Error
	}
	return nil
}

func (p *Postgres) GetCustomers(organizationID string, pageSize, pageNumber int) (model.PagedListResponse[model_customer.Customer], error) {

	var items []model_customer.Customer
	res := p.db.Where("organization_id = ?", organizationID).Limit(pageSize).Offset((pageNumber - 1) * pageSize).Find(&items)
	if res.Error != nil {
		return model.PagedListResponse[model_customer.Customer]{}, res.Error
	}
	var totalItems int64
	p.db.Model(&model_customer.Customer{}).Where("organization_id = ?", organizationID).Count(&totalItems)

	totalPages := int((totalItems + int64(pageSize) - 1) / int64(pageSize))

	response := model.PagedListResponse[model_customer.Customer]{
		Items:       items,
		TotalItems:  int(totalItems),
		TotalPages:  totalPages,
		CurrentPage: pageNumber,
	}
	return response, nil
}

func (p *Postgres) GetCustomer(id string, organizationID string) (*model_customer.Customer, error) {
	var customer model_customer.Customer
	res := p.db.Where("id = ? AND organization_id = ?", id, organizationID).First(&customer)
	if res.Error != nil {
		if res.Error == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, res.Error
	}
	return &customer, nil
}

func (p *Postgres) DeleteCustomer(id string, organizationID string) error {
	res := p.db.Where("id = ? AND organization_id = ?", id, organizationID).Delete(&model_customer.Customer{})
	if res.Error != nil {
		return res.Error
	}
	return nil
}
