package prodsvc

import (
	"OnlineStoreBackend/models"
	"OnlineStoreBackend/pkgs/utils"
	"OnlineStoreBackend/requests"
	prodattrvalsvc "OnlineStoreBackend/services/attribute_values"
	prodattrsvc "OnlineStoreBackend/services/attributes"
	linksvc "OnlineStoreBackend/services/links"
	prodcatesvc "OnlineStoreBackend/services/product_categories"
	prodtagsvc "OnlineStoreBackend/services/product_tags"
	chansvc "OnlineStoreBackend/services/related_channels"
	contsvc "OnlineStoreBackend/services/related_contents"
	"encoding/json"
)

func (service *Service) Create(modelProduct *models.Products, req *requests.RequestProductWithDetail) error {
	modelProduct.StoreID = req.StoreID
	modelProduct.Title = req.Title
	modelProduct.ShortDescription = req.ShortDescription
	modelProduct.LongDescription = req.LongDescirpiton
	modelProduct.Sku = utils.CleanSpecialLetters(modelProduct.Title)

	contService := contsvc.NewServiceProductContent(service.DB)
	chanService := chansvc.NewServiceProductChannel(service.DB)
	cateService := prodcatesvc.NewServiceProductCategory(service.DB)
	tagService := prodtagsvc.NewServiceProductTag(service.DB)
	attrService := prodattrsvc.NewServiceProductAttribute(service.DB)
	linkService := linksvc.NewServiceProductLinked(service.DB)
	valService := prodattrvalsvc.NewServiceProductAttributeValue(service.DB)

	imageUrls, _ := json.Marshal(req.ImageUrls)
	modelProduct.ImageUrls = string(imageUrls)
	service.DB.Create(modelProduct)

	productID := uint64(modelProduct.ID)

	for _, channelID := range req.RelatedChannels {
		chanService.Create(channelID, productID)
	}

	for _, contentID := range req.RelatedContents {
		contService.Create(contentID, productID)
	}

	for _, categoryID := range req.Categories {
		cateService.Create(categoryID, productID)
	}

	for _, tag := range req.Tags {
		tagService.Create(tag, modelProduct)
	}

	for _, linkID := range req.CrossSell {
		linkService.Create(productID, linkID, utils.CrossSell)
	}

	for _, linkID := range req.UpSell {
		linkService.Create(productID, linkID, utils.UpSell)
	}

	for name, values := range req.Attributes {
		if len(values) > 0 {
			modelAttr := models.ProductAttributes{}
			attrService.Create(productID, &requests.RequestAttribute{Name: name}, &modelAttr)
			attributeID := modelAttr.ID
			for index, value := range values {
				if index != 0 {
					valService.Create(uint64(attributeID), value)
				}
			}
		}
	}

	return nil
}

func (service *Service) CreateWithCSV(modelNewProds *[]models.Products, prodSkus []string, prodIndices map[string]int) {
	modelCurProds := []models.Products{}
	service.DB.Where("sku In (?)", prodSkus).Find(&modelCurProds)
	for _, modelProd := range modelCurProds {
		index := prodIndices[modelProd.Sku] - 1
		(*modelNewProds)[index].ID = modelProd.ID
	}
	service.DB.Save(modelNewProds)
}
