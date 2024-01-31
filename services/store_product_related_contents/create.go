package prodCont

import (
	"OnlineStoreBackend/models"
	"OnlineStoreBackend/requests"
)

func (service *Service) Create(productID uint64, req *requests.RequestProductContent, modelContents *[]models.ProductContents) error {
	service.DB.Where("store_product_id = ?", productID).Delete(models.ProductContents{})
	for _, content := range req.Contents {
		modelContent := models.ProductContents{}
		if err := service.DB.Table("contents").Select("id As content_id, ? As store_product_id", productID).
			Where("title = ? And deleted_at Is Null", content).Limit(1).Scan(&modelContent).Error; err == nil {
			service.DB.Create(&modelContent)
			*modelContents = append(*modelContents, modelContent)
		}
	}
	return nil
}
