package prodattrsvc

import (
	"OnlineStoreBackend/models"
)

func (service *Service) Delete(attributeID uint64) {
	modelDetails := make([]models.ProductVariationDetails, 0)
	service.DB.
		Table("store_product_variation_details As dets").
		Select("dets.*").
		Joins("Left Join store_product_attribute_values As vals On vals.id = dets.attribute_value_id").
		Where("vals.attribute_id = ?", attributeID).
		Where("vals.deleted_at Is Null And dets.deleted_at Is Null").
		Scan(&modelDetails)
	varIDs := make([]uint64, 0)
	detailIDs := make([]uint64, 0)
	for _, modelDetail := range modelDetails {
		varIDs = append(varIDs, modelDetail.VariationID)
		detailIDs = append(detailIDs, uint64(modelDetail.ID))
	}
	service.DB.Where("id In (?)", varIDs).Delete(models.ProductVariations{})
	service.DB.Where("id In (?)", detailIDs).Delete(models.ProductVariationDetails{})
	service.DB.Where("attribute_id = ?", attributeID).Delete(models.ProductAttributeValues{})
	service.DB.Where("id = ?", attributeID).Delete(models.ProductAttributes{})
}
