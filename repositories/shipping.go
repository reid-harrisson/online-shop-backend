package repositories

import (
	"OnlineStoreBackend/models"

	"github.com/jinzhu/gorm"
)

type RepositoryShipping struct {
	DB *gorm.DB
}

func NewRepositoryShipping(db *gorm.DB) *RepositoryShipping {
	return &RepositoryShipping{DB: db}
}

func (repository *RepositoryShipping) ReadByProductID(modelShipData *models.ShippingData, productID uint64) {
	repository.DB.Where("product_id = ?", productID).First(modelShipData)
}

func (repository *RepositoryShipping) ReadOptionsByStoreID(modelOptions *[]models.ShippingMethods, storeID uint64) {
	repository.DB.Where("store_id = ?", storeID).Find(modelOptions)
}

func (repository *RepositoryShipping) GetDefaultMethodID(storeID uint64) uint64 {
	id := uint64(0)
	repository.DB.Model(models.ShippingMethods{}).Select("id").Where("store_id = ?", id).Limit(1).Scan(&id)
	return id
}

func (repository *RepositoryShipping) GetDefaultShippingPrice(variationID uint64, methodID uint64, total float64, quantity float64) float64 {
	price := float64(0)
	repository.DB.Table("store_shipping_methods As mets").
		Select("(mets.flat_rate + mets.base_rate + mets.rate_per_item * ? + mets.rate_per_weight * ships.weight + mets.rate_per_total * ? / 100) As ", quantity, total).
		Joins("Left Join store_shipping_data As ships On ships.variation_id = ?", variationID).
		Where("mets.id = ?", methodID).
		Limit(1).Scan(&price)
	return price
}
