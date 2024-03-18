package prodcatesvc

import (
	"OnlineStoreBackend/models"
	"fmt"
)

func (service *Service) Create(categoryID uint64, productID uint64) {
	service.DB.Create(&models.ProductCategories{
		CategoryID: categoryID,
		ProductID:  productID,
	})
}

func (service *Service) CreateWithCSV(modelNewCates *[]models.ProductCategories, cateMatches []string, cateIndices map[string]int) {
	modelCurCates := []models.ProductCategories{}
	service.DB.Where("Concat(product_id, ':', category_id) In (?)", cateMatches).Find(&modelCurCates)
	for _, modelCate := range modelCurCates {
		match := fmt.Sprintf("%d:%d", modelCate.ProductID, modelCate.CategoryID)
		index := cateIndices[match]
		(*modelNewCates)[index].ID = modelCate.ID
	}
	service.DB.Save(modelNewCates)
}
