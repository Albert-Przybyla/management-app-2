package database

import model_storage "menagment-app-2/src/model/storage"

func (p *Postgres) CreateStorage(req model_storage.CreateStorageRequest, organizationId string) error {
	storage := model_storage.Storage{
		Name:           req.Name,
		OrganizationID: organizationId,
	}
	res := p.db.Create(&storage)
	if res.Error != nil {
		return res.Error
	}
	return nil
}
