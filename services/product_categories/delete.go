package prodcatesvc

import "OnlineStoreBackend/models"

func (service *Service) Delete(categoryID uint64) {
	service.DB.Where("category_id = ?", categoryID).Delete(models.ProductCategories{})
}
