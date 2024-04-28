package prodattrsvc

import (
	"OnlineStoreBackend/models"
)

func (service *Service) Delete(attributeID uint64) error {
	modelDetails := make([]models.VariationDetails, 0)
	err := service.DB.
		Table("store_product_variation_details As dets").
		Select("dets.*").
		Joins("Left Join store_product_attribute_values As vals On vals.id = dets.attribute_value_id").
		Where("vals.attribute_id = ?", attributeID).
		Where("vals.deleted_at Is Null And dets.deleted_at Is Null").
		Scan(&modelDetails).Error
	if err != nil {
		return err
	}
	varIDs := make([]uint64, 0)
	detailIDs := make([]uint64, 0)
	for _, modelDetail := range modelDetails {
		varIDs = append(varIDs, modelDetail.VariationID)
		detailIDs = append(detailIDs, uint64(modelDetail.ID))
	}
	err = service.DB.Where("id In (?)", varIDs).Delete(&models.Variations{}).Error
	if err != nil {
		return err
	}
	err = service.DB.Where("id In (?)", detailIDs).Delete(&models.VariationDetails{}).Error
	if err != nil {
		return err
	}
	err = service.DB.Where("attribute_id = ?", attributeID).Delete(&models.AttributeValues{}).Error
	if err != nil {
		return err
	}
	err = service.DB.Where("id = ?", attributeID).Delete(&models.Attributes{}).Error
	if err != nil {
		return err
	}
	return nil
}
