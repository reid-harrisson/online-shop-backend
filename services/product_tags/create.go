package ptagsvc

import (
	"OnlineStoreBackend/models"
	"OnlineStoreBackend/requests"
	basetag "OnlineStoreBackend/services/tags"
)

func (service *Service) Create(productID uint64, req *requests.RequestTag, modelPTags *[]models.ProductTags) {
	service.DB.Where("product_id = ?", productID).Delete(models.ProductTags{})
	tagService := basetag.CreateService(service.DB)
	for _, tag := range req.Tags {
		modelTag := models.BaseTags{}
		tagService.Create(tag, &modelTag)

		modelPTag := models.ProductTags{
			ProductID: productID,
			TagID:     uint64(modelTag.ID),
		}
		service.DB.Where("produdct_id = ? And tag_id = ?", productID, modelPTag.TagID).First(modelPTag)
		service.DB.Save(modelPTag)

		*modelPTags = append(*modelPTags, modelPTag)
	}
}
