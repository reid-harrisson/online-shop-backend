package chansvc

import (
	"OnlineStoreBackend/models"
)

func (service *Service) Delete(channelID uint64) {
	service.DB.Where("channel_id = ?", channelID).Delete(models.ProductChannels{})
}
