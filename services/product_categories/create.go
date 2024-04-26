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

func (service *Service) CreateWithCSV(modelNewCates *[]models.ProductCategories, cateMatches []string, cateIndices map[string]int) error {
	modelCurCates := []models.ProductCategories{}
	if err := service.DB.Where("Concat(product_id, ':', category_id) In (?)", cateMatches).Find(&modelCurCates).Error; err != nil {
		return err
	}
	for _, modelCate := range modelCurCates {
		match := fmt.Sprintf("%d:%d", modelCate.ProductID, modelCate.CategoryID)
		index := cateIndices[match]
		(*modelNewCates)[index].ID = modelCate.ID
	}
	return service.DB.Save(modelNewCates).Error
}
