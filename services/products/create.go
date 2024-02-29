package prodsvc

import (
	"OnlineStoreBackend/models"
	"OnlineStoreBackend/pkgs/utils"
	"OnlineStoreBackend/repositories"
	"OnlineStoreBackend/requests"
	catesvc "OnlineStoreBackend/services/categories"
	prodattrvalsvc "OnlineStoreBackend/services/product_attribute_values"
	prodattrsvc "OnlineStoreBackend/services/product_attributes"
	prodcatesvc "OnlineStoreBackend/services/product_categories"
	linksvc "OnlineStoreBackend/services/product_links"
	prodtagsvc "OnlineStoreBackend/services/product_tags"
	chansvc "OnlineStoreBackend/services/related_channels"
	contsvc "OnlineStoreBackend/services/related_contents"
	tagsvc "OnlineStoreBackend/services/tags"
	"encoding/json"
	"strconv"
	"strings"
)

func (service *Service) Create(modelProduct *models.Products, req *requests.RequestProductWithDetail) error {
	modelProduct.StoreID = req.StoreID
	modelProduct.Title = req.Title
	modelProduct.ShortDescription = req.ShortDescription
	modelProduct.LongDescription = req.LongDescirpiton

	prodRepo := repositories.NewRepositoryProduct(service.DB)
	prodRepo.ReadCurrencyID(modelProduct, req.StoreID)

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
		tagService.Create(tag, productID)
	}

	for _, linkID := range req.CrossSell {
		linkService.Create(productID, linkID, utils.CrossSell)
	}

	for _, linkID := range req.UpSell {
		linkService.Create(productID, linkID, utils.UpSell)
	}

	for name, values := range req.Attributes {
		if len(values) > 0 {
			unit := values[0]
			modelAttr := models.ProductAttributes{}
			attrService.Create(productID, &requests.RequestAttribute{Name: name, Unit: unit}, &modelAttr)
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

func (service *Service) CreateWithCSV(modelProduct *models.Products, modelCsv models.CSVs, storeID uint64) {
	imageUrls := strings.Split(modelCsv.Images, ",")
	images, _ := json.Marshal(imageUrls)
	modelProduct.StoreID = storeID
	modelProduct.Title = modelCsv.Name
	modelProduct.ShortDescription = modelCsv.ShortDescription
	modelProduct.LongDescription = modelCsv.Description
	modelProduct.ImageUrls = string(images)
	modelProduct.MinimumStockLevel, _ = strconv.ParseFloat(modelCsv.LowStockAmount, 64)
	switch modelCsv.Published {
	case "1":
		modelProduct.Status = utils.Approved
	case "0":
		modelProduct.Status = utils.Draft
	}

	modelCategories := make([]models.StoreCategories, 0)
	categories := strings.Split(modelCsv.Categories, ">")
	cateService := catesvc.NewServiceCategory(service.DB)
	cateService.CreateWithCSV(&modelCategories, categories, storeID)

	prodcateService := prodcatesvc.NewServiceProductCategory(service.DB)

	modelTags := make([]models.StoreTags, 0)
	tags := strings.Split(modelCsv.Tags, ",")
	tagService := tagsvc.NewServiceTag(service.DB)
	tagService.CreateWithCSV(&modelTags, tags, storeID)

	prodtagService := prodtagsvc.NewServiceProductTag(service.DB)

	switch modelCsv.Type {
	case "simple":
		service.DB.Create(modelProduct)
		prodcateService.CreateWithCSV(modelCategories, uint64(modelProduct.ID))
		prodtagService.CreateWithCSV(modelTags, uint64(modelProduct.ID))
	case "variable":
	case "variation":
	}
}
