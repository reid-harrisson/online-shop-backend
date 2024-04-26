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
	attrService := prodattrsvc.NewServiceAttribute(service.DB)
	linkService := linksvc.NewServiceLink(service.DB)
	valService := prodattrvalsvc.NewServiceAttributeValue(service.DB)

	imageUrls, _ := json.Marshal(req.ImageUrls)
	modelProduct.ImageUrls = string(imageUrls)
	err := service.DB.Create(modelProduct).Error
	if err != nil {
		return err
	}

	productID := uint64(modelProduct.ID)

	err = chanService.Update(productID, &requests.RequestProductChannel{
		ChannelIDs: req.RelatedChannels,
	})
	if err != nil {
		return err
	}

	err = contService.Update(productID, &requests.RequestProductContent{
		ContentIDs: req.RelatedContents,
	})
	if err != nil {
		return err
	}

	for _, categoryID := range req.Categories {
		err = cateService.Create(categoryID, productID)
		if err != nil {
			return err
		}
	}

	for _, tag := range req.Tags {
		err = tagService.Create(tag, modelProduct)
		if err != nil {
			return err
		}
	}

	for _, linkID := range req.CrossSell {
		err = linkService.Create(productID, linkID, utils.CrossSell)
		if err != nil {
			return err
		}
	}

	for _, linkID := range req.UpSell {
		err = linkService.Create(productID, linkID, utils.UpSell)
		if err != nil {
			return err
		}
	}

	for name, values := range req.Attributes {
		if len(values) > 0 {
			modelAttr := models.Attributes{}
			err = attrService.Create(productID, &requests.RequestAttribute{Name: name}, &modelAttr)
			if err != nil {
				return err
			}

			attributeID := modelAttr.ID
			for _, value := range values {
				err = valService.Create(uint64(attributeID), value)
				if err != nil {
					return err
				}
			}
		}
	}

	return nil
}

func (service *Service) CreateWithCSV(modelNewProds *[]models.Products, prodSkus []string, prodIndices map[string]int) error {
	modelCurProds := []models.Products{}
	if err := service.DB.Where("sku In (?)", prodSkus).Find(&modelCurProds).Error; err != nil {
		return err
	}
	for _, modelProd := range modelCurProds {
		index := prodIndices[modelProd.Sku] - 1
		(*modelNewProds)[index].ID = modelProd.ID
	}
	return service.DB.Save(modelNewProds).Error
}
