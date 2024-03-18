package prodcatesvc

import (
	"OnlineStoreBackend/models"
	"fmt"
	"sort"
)

func (service *Service) Create(categoryID uint64, productID uint64) {
	service.DB.Create(&models.ProductCategories{
		CategoryID: categoryID,
		ProductID:  productID,
	})
}

func (service *Service) CreateWithCSV(prodCates map[uint64][]uint64) {
	modelNewCates := []models.ProductCategories{}
	modelCurCates := []models.ProductCategories{}
	matches := []string{}
	indices := map[string]int{}
	for prodID, cateIDs := range prodCates {
		for _, cateID := range cateIDs {
			modelNewCates = append(modelNewCates, models.ProductCategories{
				ProductID:  prodID,
				CategoryID: cateID,
			})
			match := fmt.Sprintf("%d:%d", prodID, cateID)
			if indices[match] == 0 {
				matches = append(matches, match)
				indices[match] = len(matches)
			}
		}
	}
	sort.Slice(modelNewCates, func(i, j int) bool {
		if modelNewCates[i].ProductID == modelNewCates[j].ProductID {
			return modelNewCates[i].CategoryID < modelNewCates[j].CategoryID
		}
		return modelNewCates[i].ProductID < modelNewCates[j].ProductID
	})
	service.DB.Where("Concat(product_id, ':', category_id) In (?)", matches).Find(&modelCurCates)
	for _, modelCate := range modelCurCates {
		match := fmt.Sprintf("%d:%d", modelCate.ProductID, modelCate.CategoryID)
		index := indices[match] - 1
		modelNewCates[index].ID = modelCate.ID
	}
	service.DB.Save(&modelNewCates)
}

func (service *Service) CreateWithCSV1(modelCategories []models.StoreCategories, productID uint64) {
	for _, modelCategory := range modelCategories {
		service.DB.Create(&models.ProductCategories{
			CategoryID: uint64(modelCategory.ID),
			ProductID:  productID,
		})
	}
}
