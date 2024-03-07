package repositories

import (
	"OnlineStoreBackend/models"

	"gorm.io/gorm"
)

type RepositoryShippingZone struct {
	DB *gorm.DB
}

func NewRepositoryShippingZone(db *gorm.DB) *RepositoryShippingZone {
	return &RepositoryShippingZone{DB: db}
}

func (repository *RepositoryShippingZone) ReadDetailByID(modelZones *models.ShippingZonesWithPlace, zoneID uint64) error {
	return repository.DB.Table("store_shipping_zones As zones").
		Select("zones.*, Group_Concat(places.id) As place_ids, Group_Concat(places.name) As place_names").
		Joins("Join store_shipping_places As places On places.zone_id = zones.id").
		Group("zones.id").
		Where("zones.id = ?", zoneID).
		Scan(modelZones).Error
}
