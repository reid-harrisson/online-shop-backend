package prodcatesvc

import (
	"OnlineStoreBackend/models"
)

func (service *Service) Create(categoryID uint64, productID uint64) {
	service.DB.Create(models.ProductCategories{
		CategoryID: categoryID,
		ProductID:  productID,
	})
}
