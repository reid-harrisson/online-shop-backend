package chansvc

import (
	"OnlineStoreBackend/models"
	"OnlineStoreBackend/repositories"
	"OnlineStoreBackend/requests"
)

func (service *Service) Update(modelChannels *[]models.ProductChannelsWithName, req *requests.RequestProductChannel, productID uint64) {
	filterKeys := make(map[uint64]int)
	for _, modelChannel := range *modelChannels {
		filterKeys[modelChannel.ChannelID] = 1
	}
	for _, channelID := range req.ChannelIDs {
		if filterKeys[channelID] == 1 {
			filterKeys[channelID] = 3
		} else {
			filterKeys[channelID] = 2
		}
	}

	for channelID, key := range filterKeys {
		if key == 1 {
			service.Delete(channelID)
		} else if key == 2 {
			service.Create(channelID, productID)
		}
	}

	chanRepo := repositories.NewRepositoryProductChannel(service.DB)
	chanRepo.ReadByProductID(modelChannels, productID)
}
