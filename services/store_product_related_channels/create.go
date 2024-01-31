package prodChan

import (
	"OnlineStoreBackend/models"
	"OnlineStoreBackend/requests"
)

func (service *Service) Create(productID uint64, req *requests.RequestProductChannel, modelChannels *[]models.ProductChannels) error {
	service.DB.Where("store_product_id = ?", productID).Delete(models.ProductChannels{})
	for _, channel := range req.Channels {
		modelChannel := models.ProductChannels{}
		if err := service.DB.Table("channels").Select("id As channel_id, ? As store_product_id", productID).
			Where("name = ? And deleted_at Is Null", channel).Limit(1).Scan(&modelChannel).Error; err == nil {
			service.DB.Create(&modelChannel)
			*modelChannels = append(*modelChannels, modelChannel)
		}
	}
	return nil
}
