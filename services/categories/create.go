package catesvc

import (
	"OnlineStoreBackend/models"
	"OnlineStoreBackend/requests"
)

func (service *Service) Create(modelCategory *models.Categories, req *requests.RequestCategory, storeID uint64) error {
	modelCategory.Name = req.Name
	modelCategory.StoreID = storeID
	modelCategory.ParentID = req.ParentID
	return service.DB.Create(modelCategory).Error
}

func (service *Service) CreateWithCSV(modelNewCates *[]models.Categories, cateNames []string, cateParents map[string]string, cateIndices map[string]int) error {
	modelCurCates := []models.Categories{}
	if err := service.DB.Where("name In (?)", cateNames).Find(&modelCurCates).Error; err != nil {
		return err
	}
	for _, modelCate := range modelCurCates {
		index := cateIndices[modelCate.Name] - 1
		(*modelNewCates)[index].ID = modelCate.ID
	}
	if err := service.DB.Save(modelNewCates).Error; err != nil {
		return err
	}
	for index, modelCate := range *modelNewCates {
		parentID := cateIndices[cateParents[modelCate.Name]]
		if parentID > 0 {
			(*modelNewCates)[index].ParentID = uint64((*modelNewCates)[parentID-1].ID)
		}
	}
	return service.DB.Save(modelNewCates).Error
}
