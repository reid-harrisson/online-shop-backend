package prodvarsvc

import "OnlineStoreBackend/models"

func (service *Service) Delete(variant string, productID uint64) {
	service.DB.Where("variant = ? And product_id = ?", variant, productID).Delete(models.ProductVariations{})
}
