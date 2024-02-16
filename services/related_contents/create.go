package contsvc

import (
	"OnlineStoreBackend/models"
)

func (service *Service) Create(contentID uint64, productID uint64) {
	modelContent := models.BaseContents{}
	service.DB.Table("contents").Select("id").Where("id = ?", contentID).Scan(&modelContent)
	if modelContent.ID != 0 {
		service.DB.Create(&models.ProductContents{
			ContentID: contentID,
			ProductID: productID,
		})
	}
}
