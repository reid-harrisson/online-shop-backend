package prodtagsvc

import (
	"OnlineStoreBackend/models"
	"OnlineStoreBackend/requests"
	tagsvc "OnlineStoreBackend/services/tags"
)

func (service *Service) Create(productID uint64, req *requests.RequestTag, modelProdTags *[]models.ProductTags) {
	service.DB.Where("product_id = ?", productID).Delete(models.ProductTags{})
	tagService := tagsvc.CreateService(service.DB)
	for _, tag := range req.Tags {
		modelTag := models.BaseTags{}
		tagService.Create(tag, &modelTag)

		modelProdTag := models.ProductTags{
			ProductID: productID,
			TagID:     uint64(modelTag.ID),
		}
		service.DB.Where("product_id = ? And tag_id = ?", productID, modelProdTag.TagID).First(&modelProdTag)
		service.DB.Save(&modelProdTag)

		*modelProdTags = append(*modelProdTags, modelProdTag)
	}
}
