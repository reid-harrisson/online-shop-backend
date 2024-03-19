package chansvc

import (
	"OnlineStoreBackend/models"
	"OnlineStoreBackend/requests"
	"fmt"
)

func (service *Service) Update(productID uint64, req *requests.RequestProductChannel) {
	modelNewChans := []models.ProductChannels{}
	modelCurChans := []models.ProductChannels{}
	chanIndices := map[string]int{}
	chanMatches := []string{}
	for index, channelID := range req.ChannelIDs {
		match := fmt.Sprintf("%d:%d", productID, channelID)
		chanMatches = append(chanMatches, match)
		chanIndices[match] = index
		modelNewChans = append(modelNewChans, models.ProductChannels{
			ProductID: productID,
			ChannelID: channelID,
		})
	}
	service.DB.Where("Concat(product_id, ':', channel_id) In (?)", chanMatches).Find(&modelCurChans)
	service.DB.Where("Concat(product_id, ':', channel_id) Not In (?) And product_id = ?", chanMatches, productID).Delete(&models.ProductChannels{})
	for _, modelChan := range modelCurChans {
		match := fmt.Sprintf("%d:%d", modelChan.ProductID, modelChan.ChannelID)
		index := chanIndices[match]
		modelNewChans[index].ID = modelChan.ID
	}
	service.DB.Save(&modelNewChans)
}
