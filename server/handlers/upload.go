package handlers

import (
	"OnlineStoreBackend/models"
	"OnlineStoreBackend/pkgs/utils"
	"OnlineStoreBackend/responses"
	s "OnlineStoreBackend/server"
	prodattrvalsvc "OnlineStoreBackend/services/attribute_values"
	prodattrsvc "OnlineStoreBackend/services/attributes"
	catesvc "OnlineStoreBackend/services/categories"
	prodcatesvc "OnlineStoreBackend/services/product_categories"
	prodtagsvc "OnlineStoreBackend/services/product_tags"
	prodsvc "OnlineStoreBackend/services/products"
	tagsvc "OnlineStoreBackend/services/tags"
	prodvarsvc "OnlineStoreBackend/services/variations"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
)

type HandlersUpload struct {
	server *s.Server
}

func NewHandlersUpload(server *s.Server) *HandlersUpload {
	return &HandlersUpload{server: server}
}

func readCSV(file *multipart.File, modelCsvs *[]models.CSVs) {
	reader := csv.NewReader(*file)

	header, err := reader.Read()
	if err != nil {
		return
	}

	for {
		if record, err := reader.Read(); err == io.EOF {
			break
		} else {
			mapCsv := make(map[string]string)
			for index := range record {
				mapCsv[header[index]] = record[index]
			}
			jsonCsv, _ := json.Marshal(mapCsv)
			modelCsv := models.CSVs{}
			modelCsv.ID = record[0]
			if err := json.Unmarshal(jsonCsv, &modelCsv); err == nil {
				*modelCsvs = append(*modelCsvs, modelCsv)
			}
		}
	}
}

// @Summary Upload a CSV file
// @Description Upload a CSV file to the server
// @Tags CSV Upload
// @Accept multipart/form-data
// @Produce json
// @Param store_id query int true "Store ID"
// @Param file formData file true "CSV file to upload"
// @Success 200 {string} string "File uploaded successfully"
// @Router /store/api/v1/upload/csv [post]
func (h *HandlersUpload) UploadCSV(c echo.Context) error {
	storeID, _ := strconv.ParseUint(c.QueryParam("store_id"), 10, 64)

	file, err := c.FormFile("file")
	if err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, err.Error())
	}
	src, err := file.Open()
	if err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, err.Error())
	}
	defer src.Close()

	modelCsvs := make([]models.CSVs, 0)
	readCSV(&src, &modelCsvs)

	// mapSku := make(map[string]uint64)
	// mapIDs := make(map[string]string)

	// modelProducts := make([]models.Products, 0)
	// prodService := prodsvc.NewServiceProduct(h.server.DB)
	// for _, modelCsv := range modelCsvs {
	// 	modelProduct := models.Products{}
	// 	prodService.CreateWithCSV(&modelProduct, modelCsv, storeID, &mapIDs)
	// 	if modelProduct.ID != 0 {
	// 		if mapSku[modelProduct.Sku] == 0 {
	// 			modelProducts = append(modelProducts, modelProduct)
	// 			mapSku[modelProduct.Sku] = uint64(modelProduct.ID)
	// 			mapIDs[modelCsv.ID] = modelProduct.Sku
	// 		}
	// 	}
	// }

	// linkService := linksvc.NewServiceProductLinked(h.server.DB)
	// for _, modelCsv := range modelCsvs {
	// 	upSells := strings.Split(modelCsv.Upsells, ",")
	// 	crossSells := strings.Split(modelCsv.CrossSells, ",")
	// 	sku := strings.TrimSpace(modelCsv.Sku)
	// 	if modelCsv.Parent != "" {
	// 		sku = strings.TrimSpace(modelCsv.Sku)
	// 	}
	// 	for _, upSell := range upSells {
	// 		if len(upSell) > 3 && upSell[:3] == "id:" {
	// 			id := upSell[3:]
	// 			upSell = mapIDs[id]
	// 		}
	// 		if mapSku[sku] != 0 && mapSku[upSell] != 0 {
	// 			linkService.Create(mapSku[sku], mapSku[upSell], utils.UpSell)
	// 		}
	// 	}
	// 	for _, crossSell := range crossSells {
	// 		if len(crossSell) > 3 && crossSell[:3] == "id:" {
	// 			id := crossSell[3:]
	// 			crossSell = mapIDs[id]
	// 		}
	// 		if mapSku[sku] != 0 && mapSku[crossSell] != 0 {
	// 			linkService.Create(mapSku[sku], mapSku[crossSell], utils.UpSell)
	// 		}
	// 	}
	// }

	// return responses.NewResponseProducts(c, http.StatusOK, modelProducts)

	// Category Informations
	cateNames := []string{}
	cateParents := map[string]string{}
	cateIndices := map[string]int{}
	modelCategories := []models.StoreCategories{}

	// Tag Informations
	tagNames := []string{}
	tagIndices := map[string]int{}
	modelTags := []models.StoreTags{}

	// Product Informations
	prodIndices := map[string]int{}
	prodSkus := []string{}
	prodCates := map[uint64][]uint64{}
	prodTags := map[uint64][]uint64{}
	modelProducts := []models.Products{}

	// Attribute Informations
	attrIndices := map[string]int{}
	attrMatches := []string{}
	modelAttrs := []models.ProductAttributes{}

	// Attribute Value Informations
	valIndices := map[string]int{}
	valMatches := []string{}
	modelVals := []models.ProductAttributeValues{}

	// Variation Informations
	varIndices := map[string]int{}
	varMatches := []string{}
	modelVars := []models.ProductVariations{}

	// Variation Detail Informations
	detIndices := map[string]int{}
	detMatches := []string{}
	modelDets := []models.ProductVariationDetails{}

	// Linked Product Informations

	for _, modelCsv := range modelCsvs {
		// Category
		cates := strings.Split(modelCsv.Categories, ",")
		cateCurrent := []uint64{}
		for _, cate := range cates {
			subcates := strings.Split(cate, ">")
			for index, subcate := range subcates {
				subcate = strings.TrimSpace(subcate)
				if subcate != "" && cateIndices[subcate] == 0 {
					if index != 0 {
						cateParents[subcate] = subcates[index-1]
					}
					cateNames = append(cateNames, subcate)
					modelCategories = append(modelCategories, models.StoreCategories{
						Name:    subcate,
						StoreID: storeID,
					})
					cateIndices[subcate] = len(modelCategories)
				}
				if subcate != "" {
					cateCurrent = append(cateCurrent, uint64(cateIndices[subcate]-1))
				}
			}
		}
		// Tag
		tags := strings.Split(modelCsv.Tags, ",")
		tagCurrent := []uint64{}
		for _, tag := range tags {
			tag = strings.TrimSpace(tag)
			if tag != "" && tagIndices[tag] == 0 {
				modelTags = append(modelTags, models.StoreTags{
					Name:    tag,
					StoreID: storeID,
				})
				tagNames = append(tagNames, tag)
				tagIndices[tag] = len(modelTags)
			}
			if tag != "" {
				tagCurrent = append(tagCurrent, uint64(tagIndices[tag]-1))
			}
		}
		// Product
		prodSku := modelCsv.Sku
		prodName := modelCsv.Name
		prodDesc := modelCsv.Description
		prodImages := strings.Split(modelCsv.Images, ", ")
		prodStatus := utils.Draft
		lowStockAmount, _ := strconv.ParseFloat(modelCsv.LowStockAmount, 64)
		if modelCsv.Type == "variation" {
			prodSku = modelCsv.Parent
			prodName = strings.Split(prodName, " - ")[0]
			prodDesc = ""
			prodImages = []string{}
		}
		if modelCsv.Published == "1" {
			prodStatus = utils.Approved
		}
		if prodIndices[prodSku] == 0 {
			images, _ := json.Marshal(prodImages)
			modelProducts = append(modelProducts, models.Products{
				Title:             prodName,
				ShortDescription:  modelCsv.ShortDescription,
				LongDescription:   prodDesc,
				ImageUrls:         string(images),
				ShippingClass:     modelCsv.ShippingClass,
				Sku:               prodSku,
				Type:              utils.ProductTypesFromString(modelCsv.Type),
				Status:            prodStatus,
				StoreID:           storeID,
				MinimumStockLevel: lowStockAmount,
			})
			prodSkus = append(prodSkus, prodSku)
			size := len(modelProducts)
			prodIndices[prodSku] = size
			// Product Category and Tag
			prodCates[uint64(size-1)] = append(prodCates[uint64(size-1)], cateCurrent...)
			prodTags[uint64(size-1)] = append(prodTags[uint64(size-1)], tagCurrent...)
		}
		size := prodIndices[prodSku]
		valCurrent := []uint64{}
		// First Attribute
		match := strconv.Itoa(size-1) + ":" + modelCsv.AttributeName
		if modelCsv.AttributeName != "" && attrIndices[match] == 0 {
			modelAttrs = append(modelAttrs, models.ProductAttributes{
				AttributeName: modelCsv.AttributeName,
			})
			attrMatches = append(attrMatches, match)
			subsize := len(modelAttrs)
			attrIndices[match] = subsize
		}
		// First Attribute Value
		vals := strings.Split(modelCsv.Attribute1Values, ",")
		attrIndex := attrIndices[match] - 1
		for _, val := range vals {
			val = strings.TrimSpace(val)
			match := strconv.Itoa(attrIndex) + ":" + val
			if val != "" && valIndices[match] == 0 {
				modelVals = append(modelVals, models.ProductAttributeValues{
					AttributeValue: val,
				})
				valMatches = append(valMatches, match)
				valIndices[match] = len(modelVals)
			}
			if val != "" {
				valCurrent = append(valCurrent, uint64(valIndices[match]-1))
			}
		}
		// Second Attribute
		match = strconv.Itoa(size-1) + ":" + modelCsv.Attribute2Name
		if modelCsv.Attribute2Name != "" && attrIndices[match] == 0 {
			modelAttrs = append(modelAttrs, models.ProductAttributes{
				AttributeName: modelCsv.Attribute2Name,
			})
			attrMatches = append(attrMatches, match)
			subsize := len(modelAttrs)
			attrIndices[match] = subsize
		}
		// Second Attribute Values
		vals = strings.Split(modelCsv.Attribute2Values, ",")
		for _, val := range vals {
			val = strings.TrimSpace(val)
			match := strconv.Itoa(attrIndex) + ":" + val
			if val != "" && valIndices[match] == 0 {
				modelVals = append(modelVals, models.ProductAttributeValues{
					AttributeValue: val,
				})
				valMatches = append(valMatches, match)
				valIndices[match] = len(modelVals)
			}
			if val != "" {
				valCurrent = append(valCurrent, uint64(valIndices[match]-1))
			}
		}
		// Variation
		if modelCsv.Type != "variable" {
			varSku := modelCsv.Sku
			stock, _ := strconv.ParseFloat(modelCsv.Stock, 64)
			regularPrice, _ := strconv.ParseFloat(modelCsv.RegularPrice, 64)
			salePrice, _ := strconv.ParseFloat(modelCsv.SalePrice, 64)
			backordersAllowed := utils.Disabled
			if modelCsv.BackordersAllowed == "1" {
				backordersAllowed = utils.Enabled
			}
			varDesc := modelCsv.Description
			varImages := strings.Split(modelCsv.Images, ",")
			prodIndex := prodIndices[modelCsv.Parent] - 1
			if modelCsv.Type == "simple" {
				varDesc = ""
				varImages = []string{}
				prodIndex = prodIndices[modelCsv.Sku] - 1
			}
			if varSku == "" {
				varSku = modelCsv.Parent
				if modelCsv.Attribute1Values != "" {
					varSku += "-" + modelCsv.Attribute1Values
				}
				if modelCsv.Attribute2Values != "" {
					varSku += "-" + modelCsv.Attribute2Values
				}
			}
			varImgJson, _ := json.Marshal(varImages)
			match := fmt.Sprintf("%d:%s", prodIndex, varSku)
			if varIndices[match] == 0 {
				discountAmount := 0.0
				if salePrice != 0 {
					discountAmount = regularPrice - salePrice
				}
				modelVars = append(modelVars, models.ProductVariations{
					Sku:             varSku,
					Price:           regularPrice,
					StockLevel:      stock,
					DiscountAmount:  discountAmount,
					DiscountType:    utils.FixedAmountOff,
					ImageUrls:       string(varImgJson),
					Description:     varDesc,
					Title:           modelCsv.Name,
					BackOrderStatus: backordersAllowed,
				})
				size := len(modelVars)
				varIndices[match] = size
				varMatches = append(varMatches, match)

				for _, val := range valCurrent {
					match := strconv.Itoa(size-1) + ":" + strconv.FormatUint(val, 10)
					if detIndices[match] == 0 {
						modelDets = append(modelDets, models.ProductVariationDetails{})
						size := len(modelDets)
						detIndices[match] = size
						detMatches = append(detMatches, match)
					}
				}
			}
		}
	}

	cateService := catesvc.NewServiceCategory(h.server.DB)
	cateService.CreateWithCSV(&modelCategories, cateNames, cateParents, cateIndices)

	tagService := tagsvc.NewServiceTag(h.server.DB)
	tagService.CreateWithCSV(&modelTags, tagNames, tagIndices)

	prodService := prodsvc.NewServiceProduct(h.server.DB)
	prodService.CreateWithCSV(&modelProducts, prodSkus, prodIndices)

	attrIndices = map[string]int{}
	for index, match := range attrMatches {
		prodIndex := 0
		name := ""
		fmt.Sscanf(match, "%d:%s", &prodIndex, &name)
		prodID := modelProducts[prodIndex].ID
		match = fmt.Sprintf("%d:%s", prodID, name)
		modelAttrs[index].ProductID = uint64(prodID)
		attrMatches[index] = match
		attrIndices[match] = index
	}

	attrService := prodattrsvc.NewServiceProductAttribute(h.server.DB)
	attrService.CreateWithCSV(&modelAttrs, attrMatches, attrIndices)

	newProdCates := map[uint64][]uint64{}
	for prodIndex, cates := range prodCates {
		prodID := modelProducts[prodIndex].ID
		for _, cate := range cates {
			if cate < uint64(len(modelCategories)) {
				newProdCates[uint64(prodID)] = append(newProdCates[uint64(prodID)], uint64(modelCategories[cate].ID))
			}
		}
	}

	newProdTags := map[uint64][]uint64{}
	for prodIndex, tags := range prodTags {
		prodID := modelProducts[prodIndex].ID
		for _, tag := range tags {
			if tag < uint64(len(modelTags)) {
				newProdTags[uint64(prodID)] = append(newProdTags[uint64(prodID)], uint64(modelTags[tag].ID))
			}
		}
	}

	prodCateService := prodcatesvc.NewServiceProductCategory(h.server.DB)
	prodCateService.CreateWithCSV(newProdCates)

	prodTagService := prodtagsvc.NewServiceProductTag(h.server.DB)
	prodTagService.CreateWithCSV(newProdTags)

	for index, match := range valMatches {
		name := ""
		attrIndex := uint64(0)
		fmt.Sscanf(match, "%d:%s", &attrIndex, &name)
		attrID := modelAttrs[attrIndex].ID
		match = fmt.Sprintf("%d:%s", attrID, name)
		modelVals[index].AttributeID = uint64(attrID)
		valMatches[index] = match
		valIndices[match] = index
	}

	valService := prodattrvalsvc.NewServiceProductAttributeValue(h.server.DB)
	valService.CreateWithCSV(&modelVals, valMatches, valIndices)

	for index, match := range varMatches {
		sku := ""
		prodIndex := 0
		fmt.Sscanf(match, "%d:%s", &prodIndex, &sku)
		prodID := uint(0)
		if prodIndex >= 0 {
			prodID = modelProducts[prodIndex].ID
		}
		match = fmt.Sprintf("%d:%s", prodID, sku)
		modelVars[index].ProductID = uint64(prodID)
		varMatches[index] = match
		varIndices[match] = index
	}

	varService := prodvarsvc.NewServiceProductVariation(h.server.DB)
	varService.CreateWithCSV(&modelVars, varMatches, varIndices)

	for index, match := range detMatches {
		varIndex := 0
		valIndex := 0
		fmt.Sscanf(match, "%d:%d", &varIndex, &valIndex)
		varID := modelVars[varIndex].ID
		valID := modelVals[valIndex].ID
		match = fmt.Sprintf("%d:%d", varID, valID)
		modelDets[index].VariationID = uint64(varID)
		modelDets[index].AttributeValueID = uint64(valID)
		detMatches[index] = match
		detIndices[match] = index
	}
	return responses.NewResponseProducts(c, http.StatusOK, modelProducts)
}
