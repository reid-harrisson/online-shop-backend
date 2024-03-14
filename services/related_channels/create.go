package chansvc

import (
	"OnlineStoreBackend/models"
)

func (service *Service) Create(channelID uint64, productID uint64) {
	modelChannel := models.BaseChannels{}
	service.DB.Table("channels").Select("id").Where("id = ? And deleted_at Is Null", channelID).Scan(&modelChannel)
	if modelChannel.ID != 0 {
		service.DB.Create(&models.ProductChannels{
			ChannelID: channelID,
			ProductID: productID,
		})
	}
}
