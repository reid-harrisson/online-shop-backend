package repositories

import (
	"OnlineStoreBackend/models"

	"gorm.io/gorm"
)

type RepositoryProductChannel struct {
	DB *gorm.DB
}

func NewRepositoryProductChannel(db *gorm.DB) *RepositoryProductChannel {
	return &RepositoryProductChannel{DB: db}
}

func (repository *RepositoryProductChannel) ReadByProductID(modelChannels *[]models.ProductChannelsWithName, productID uint64) {
	repository.DB.Table("store_product_related_channels As prodchans").
		Select("prodchans.*, chans.name As channel_name").
		Joins("Join channels As chans On chans.id = prodchans.channel_id").
		Where("chans.deleted_at Is Null And prodchans.deleted_at Is Null").
		Where("prodchans.product_id = ?", productID).
		Scan(modelChannels)
}
