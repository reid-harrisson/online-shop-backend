package prodsvc

import (
	"OnlineStoreBackend/models"
)

func (service *Service) Delete(productID uint64) error {
	query := service.DB.Where("product_id = ?", productID)
	query.Delete(&models.ProductChannels{})
	query.Delete(&models.ProductContents{})
	modelAttrs := make([]models.ProductAttributes, 0)
	query.Find(&modelAttrs)
	attributeIDs := make([]uint64, 0)
	for _, modelAttr := range modelAttrs {
		attributeIDs = append(attributeIDs, uint64(modelAttr.ID))
	}
	service.DB.Where("attribute_id In (?)", attributeIDs).Delete(&models.ProductAttributeValues{})
	query.Delete(&models.ProductAttributes{})
	query.Delete(&models.ProductTags{})
	query.Delete(&models.ProductCategories{})
	query.Delete(&models.ProductReviews{})
	query.Delete(&models.ProductVariations{})
	return service.DB.Delete(&models.Products{}, productID).Error
}
