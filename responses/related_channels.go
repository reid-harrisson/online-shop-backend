package responses

import (
	"OnlineStoreBackend/models"

	"github.com/labstack/echo/v4"
)

type ResponseProductChannel struct {
	ID          uint64 `json:"id"`
	ProductID   uint64 `json:"product_id"`
	ChannelID   uint64 `json:"channel_id"`
	ChannelName string `json:"channel_name"`
}

func NewResponseProductChannels(c echo.Context, statusCode int, modelChannels []models.ProductChannelsWithName) error {
	responseChannels := make([]ResponseProductChannel, 0)
	for _, modelChannel := range modelChannels {
		responseChannels = append(responseChannels, ResponseProductChannel{
			ID:          uint64(modelChannel.ID),
			ProductID:   modelChannel.ProductID,
			ChannelID:   modelChannel.ChannelID,
			ChannelName: modelChannel.ChannelName,
		})
	}
	return Response(c, statusCode, responseChannels)
}
