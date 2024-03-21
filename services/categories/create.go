package catesvc

import (
	"OnlineStoreBackend/models"
	"OnlineStoreBackend/requests"
)

func (service *Service) Create(modelCategory *models.StoreCategories, req *requests.RequestCategory, storeID uint64) {
	modelCategory.Name = req.Name
	modelCategory.StoreID = storeID
	modelCategory.ParentID = req.ParentID
	service.DB.Create(modelCategory)
}

func (service *Service) CreateWithCSV(modelNewCates *[]models.StoreCategories, cateNames []string, cateParents map[string]string, cateIndices map[string]int) {
	modelCurCates := []models.StoreCategories{}
	service.DB.Where("name In (?)", cateNames).Find(&modelCurCates)
	for _, modelCate := range modelCurCates {
		index := cateIndices[modelCate.Name] - 1
		(*modelNewCates)[index].ID = modelCate.ID
	}
	service.DB.Save(modelNewCates)
	for index, modelCate := range *modelNewCates {
		parentID := cateIndices[cateParents[modelCate.Name]]
		if parentID > 0 {
			(*modelNewCates)[index].ParentID = uint64((*modelNewCates)[parentID-1].ID)
		}
	}
	service.DB.Save(modelNewCates)
}
