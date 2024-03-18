package prodsvc

import (
	"OnlineStoreBackend/models"
	"OnlineStoreBackend/pkgs/utils"
	"OnlineStoreBackend/repositories"
	"OnlineStoreBackend/requests"
	prodattrvalsvc "OnlineStoreBackend/services/attribute_values"
	prodattrsvc "OnlineStoreBackend/services/attributes"
	catesvc "OnlineStoreBackend/services/categories"
	linksvc "OnlineStoreBackend/services/links"
	prodcatesvc "OnlineStoreBackend/services/product_categories"
	prodtagsvc "OnlineStoreBackend/services/product_tags"
	chansvc "OnlineStoreBackend/services/related_channels"
	contsvc "OnlineStoreBackend/services/related_contents"
	tagsvc "OnlineStoreBackend/services/tags"
	prodvarsvc "OnlineStoreBackend/services/variations"
	"encoding/json"
	"strconv"
	"strings"
)

func (service *Service) Create(modelProduct *models.Products, req *requests.RequestProductWithDetail) error {
	modelProduct.StoreID = req.StoreID
	modelProduct.Title = req.Title
	modelProduct.ShortDescription = req.ShortDescription
	modelProduct.LongDescription = req.LongDescirpiton
	modelProduct.Sku = utils.CleanSpecialLetters(modelProduct.Title)

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

func (service *Service) CreateWithCSV1(modelProduct *models.Products, modelCsv models.CSVs, storeID uint64, mapIDs *map[string]string) {
	modelCategories := make([]models.StoreCategories, 0)
	categories := strings.Split(modelCsv.Categories, ", ")
	cateService := catesvc.NewServiceCategory(service.DB)
	cateService.CreateWithCSV(&modelCategories, categories, map[string]string{}, map[string]int{})

	prodcateService := prodcatesvc.NewServiceProductCategory(service.DB)

	modelTags := make([]models.StoreTags, 0)
	tags := strings.Split(modelCsv.Tags, ",")
	tagService := tagsvc.NewServiceTag(service.DB)
	tagService.CreateWithCSV(&modelTags, tags, map[string]int{})

	prodtagService := prodtagsvc.NewServiceProductTag(service.DB)

	varService := prodvarsvc.NewServiceProductVariation(service.DB)
	attrService := prodattrsvc.NewServiceProductAttribute(service.DB)

	switch modelCsv.Type {
	case "simple":
		service.DB.Where("sku = ?", modelCsv.Sku).First(&modelProduct)
		if modelProduct.ID == 0 {
			imageUrls := strings.Split(modelCsv.Images, ", ")
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
			modelProduct.ShippingClass = modelCsv.ShippingClass
			modelProduct.Sku = modelCsv.Sku
			modelProduct.Type = utils.Simple
			service.DB.Create(modelProduct)

			productID := uint64(modelProduct.ID)
			prodcateService.CreateWithCSV1(modelCategories, productID)
			prodtagService.CreateWithCSV1(modelTags, productID)
			modelVar := models.ProductVariations{}
			varService.CreateSimpleWithCSV(&modelVar, &modelCsv, productID)
		}
	case "variable":
		service.DB.Where("sku = ?", modelCsv.Sku).First(&modelProduct)
		if modelProduct.ID == 0 {
			imageUrls := strings.Split(modelCsv.Images, ", ")
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
			modelProduct.Sku = modelCsv.Sku
			modelProduct.ShippingClass = modelCsv.ShippingClass
			modelProduct.Type = utils.Variable
			service.DB.Create(modelProduct)

			productID := uint64(modelProduct.ID)
			prodcateService.CreateWithCSV1(modelCategories, productID)
			prodtagService.CreateWithCSV1(modelTags, productID)
			attrService.CreateWithCSV1(&modelCsv, productID)
		}
	case "variation":
		parentSku := modelCsv.Parent
		if parentSku[:3] == "id:" {
			id := parentSku[3:]
			parentSku = (*mapIDs)[id]
			if parentSku == "" {
				parentSku = utils.CleanSpecialLetters(strings.Split(modelCsv.Name, " - ")[0])
				(*mapIDs)[id] = parentSku
			}
		}
		if modelCsv.Sku == "" {
			modelCsv.Sku = utils.CleanSpecialLetters(modelCsv.Parent + modelCsv.AttributeValues + modelCsv.OptionalAttribute2Values)
		}
		service.DB.Where("sku = ?", parentSku).First(&modelProduct)
		modelVals := make([]models.ProductAttributeValues, 0)
		if modelProduct.ID == 0 {
			modelProduct.StoreID = storeID
			modelProduct.Title = strings.Split(modelCsv.Name, " - ")[0]
			modelProduct.MinimumStockLevel, _ = strconv.ParseFloat(modelCsv.LowStockAmount, 64)
			switch modelCsv.Published {
			case "1":
				modelProduct.Status = utils.Approved
			case "0":
				modelProduct.Status = utils.Draft
			}
			modelProduct.Sku = parentSku
			modelProduct.Type = utils.Variable
			modelProduct.ShippingClass = modelCsv.ShippingClass
			service.DB.Create(modelProduct)

			productID := uint64(modelProduct.ID)
			prodcateService.CreateWithCSV1(modelCategories, productID)
			prodtagService.CreateWithCSV1(modelTags, productID)
		}
		attrService.UpdateWithCSV(&modelVals, &modelCsv, uint64(modelProduct.ID))
		modelVar := models.ProductVariations{}
		varService.DB.Where("sku = ?", modelCsv.Sku).First(&modelVar)
		if modelVar.ID == 0 {
			varService.CreateVariableWithCSV(&modelVar, &modelCsv, uint64(modelProduct.ID), &modelVals)
		}
	}
}
