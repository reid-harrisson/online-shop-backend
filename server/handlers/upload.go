package handlers

import (
	"OnlineStoreBackend/models"
	"OnlineStoreBackend/pkgs/constants"
	errhandle "OnlineStoreBackend/pkgs/error"
	"OnlineStoreBackend/pkgs/utils"
	"OnlineStoreBackend/responses"
	s "OnlineStoreBackend/server"
	prodattrvalsvc "OnlineStoreBackend/services/attribute_values"
	prodattrsvc "OnlineStoreBackend/services/attributes"
	catesvc "OnlineStoreBackend/services/categories"
	linksvc "OnlineStoreBackend/services/links"
	prodcatesvc "OnlineStoreBackend/services/product_categories"
	prodtagsvc "OnlineStoreBackend/services/product_tags"
	prodsvc "OnlineStoreBackend/services/products"
	shipsvc "OnlineStoreBackend/services/shipping_data"
	tagsvc "OnlineStoreBackend/services/tags"
	prodvardetsvc "OnlineStoreBackend/services/variation_details"
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

func readCSV(file *multipart.File, modelCsvs *[]models.CSVs) error {
	reader := csv.NewReader(*file)

	header, err := reader.Read()
	if err != nil {
		return err
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
	return nil
}

// @Summary Upload a CSV file
// @Description Upload a CSV file to the server
// @Tags Upload Actions
// @Accept multipart/form-data
// @Produce json
// @Security ApiKeyAuth
// @Param store_id query int true "Store ID"
// @Param file formData file true "CSV file to upload"
// @Success 201 {object} []responses.ResponseProduct
// @Success 400 {object} responses.Error
// @Success 500 {object} responses.Error
// @Router /store/api/v1/upload/csv [post]
func (h *HandlersUpload) UploadCSV(c echo.Context) error {
	storeID, _ := strconv.ParseUint(c.QueryParam("store_id"), 10, 64)
	userID, err := strconv.ParseUint(c.Request().Header.Get("id"), 10, 64)
	if userID == 0 || err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, constants.InvalidData)
	}

	file, err := c.FormFile("file")
	if err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, constants.InvalidData)
	}
	src, err := file.Open()
	if err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, constants.InvalidData)
	}
	defer src.Close()

	modelCsvs := make([]models.CSVs, 0)
	if err := readCSV(&src, &modelCsvs); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, constants.InvalidData)
	}

	// Category Informations
	cateNames := []string{}
	cateParents := map[string]string{}
	cateIndices := map[string]int{}
	modelCategories := []models.Categories{}

	// Tag Informations
	tagNames := []string{}
	tagIndices := map[string]int{}
	modelTags := []models.Tags{}

	// Product Informations
	prodIndices := map[string]int{}
	prodSkus := []string{}
	modelProducts := []models.Products{}

	// Shipping Information
	modelShips := []models.ShippingData{}
	shipVarIDs := []uint64{}
	shipIndices := map[uint64]int{}

	// Product Category Information
	prodCateIndices := map[string]int{}
	prodCateMatches := []string{}
	modelProdCates := []models.ProductCategories{}

	// Product Tag Information
	prodTagIndices := map[string]int{}
	prodTagMatches := []string{}
	modelProdTags := []models.ProductTags{}

	// Attribute Informations
	attrIndices := map[string]int{}
	attrMatches := []string{}
	modelAttrs := []models.Attributes{}

	// Attribute Value Informations
	valIndices := map[string]int{}
	valMatches := []string{}
	modelVals := []models.AttributeValues{}

	// Variation Informations
	varIndices := map[string]int{}
	varMatches := []string{}
	modelVars := []models.Variations{}

	// Variation Detail Informations
	detIndices := map[string]int{}
	detMatches := []string{}
	modelDets := []models.VariationDetails{}

	// Linked Product Informations
	linkIndices := map[string]int{}
	linkMatches := []string{}
	modelLinks := []models.Links{}

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
						cateParents[subcate] = strings.TrimSpace(subcates[index-1])
					}
					cateNames = append(cateNames, subcate)
					modelCategories = append(modelCategories, models.Categories{
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
				modelTags = append(modelTags, models.Tags{
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

			// Product Category
			for _, cate := range cateCurrent {
				match := fmt.Sprintf("%d:%d", size-1, cate)
				if prodCateIndices[match] == 0 {
					modelProdCates = append(modelProdCates, models.ProductCategories{})
					prodCateMatches = append(prodCateMatches, match)
					prodCateIndices[match] = len(modelProdCates)
				}
			}

			// Product Tag
			for _, tag := range tagCurrent {
				match := fmt.Sprintf("%d:%d", size-1, tag)
				if prodTagIndices[match] == 0 {
					modelProdTags = append(modelProdTags, models.ProductTags{})
					prodTagMatches = append(prodTagMatches, match)
					prodTagIndices[match] = len(modelProdTags)
				}
			}
		}
		size := prodIndices[prodSku]
		valCurrent := []int{}

		// First Attribute
		match := strconv.Itoa(size-1) + ":" + modelCsv.AttributeName
		if modelCsv.AttributeName != "" && attrIndices[match] == 0 {
			modelAttrs = append(modelAttrs, models.Attributes{
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
				modelVals = append(modelVals, models.AttributeValues{
					AttributeValue: val,
				})
				valMatches = append(valMatches, match)
				valIndices[match] = len(modelVals)
			}
			if val != "" {
				valCurrent = append(valCurrent, valIndices[match]-1)
			}
		}
		// Second Attribute
		match = strconv.Itoa(size-1) + ":" + modelCsv.Attribute2Name
		if modelCsv.Attribute2Name != "" && attrIndices[match] == 0 {
			modelAttrs = append(modelAttrs, models.Attributes{
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
				modelVals = append(modelVals, models.AttributeValues{
					AttributeValue: val,
				})
				valMatches = append(valMatches, match)
				valIndices[match] = len(modelVals)
			}
			if val != "" {
				valCurrent = append(valCurrent, valIndices[match]-1)
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
					varSku += "-" + utils.CleanSpecialLetters(modelCsv.Attribute1Values)
				}
				if modelCsv.Attribute2Values != "" {
					varSku += "-" + utils.CleanSpecialLetters(modelCsv.Attribute2Values)
				}
			}
			varImgJson, _ := json.Marshal(varImages)
			match := fmt.Sprintf("%d:%s", prodIndex, varSku)
			if varIndices[match] == 0 {
				discountAmount := 0.0
				if salePrice != 0 {
					discountAmount = regularPrice - salePrice
				}
				modelVars = append(modelVars, models.Variations{
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

				// Shipping Data
				if modelCsv.Weight != "" {
					weight, _ := strconv.ParseFloat(modelCsv.Weight, 64)
					width, _ := strconv.ParseFloat(modelCsv.Width, 64)
					length, _ := strconv.ParseFloat(modelCsv.Length, 64)
					height, _ := strconv.ParseFloat(modelCsv.Height, 64)
					modelShips = append(modelShips, models.ShippingData{
						VariationID: uint64(size - 1),
						Weight:      weight,
						Width:       width,
						Length:      length,
						Height:      height,
					})
				}

				// Variation Detail
				for _, val := range valCurrent {
					match := strconv.Itoa(size-1) + ":" + strconv.Itoa(val)
					if detIndices[match] == 0 {
						modelDets = append(modelDets, models.VariationDetails{})
						size := len(modelDets)
						detIndices[match] = size
						detMatches = append(detMatches, match)
					}
				}
			}
		}

		// Linked Product (Up Sell)
		upSells := strings.Split(modelCsv.Upsells, ",")
		prodIndex := prodIndices[prodSku] - 1
		for _, upSell := range upSells {
			upSell = strings.TrimSpace(upSell)
			linkIndex := prodIndices[upSell] - 1
			if linkIndex >= 0 {
				match := fmt.Sprintf("%d:%d:0", prodIndex, linkIndex)
				if linkIndices[match] == 0 {
					modelLinks = append(modelLinks, models.Links{
						IsUpCross: utils.UpSell,
					})
					linkMatches = append(linkMatches, match)
					linkIndices[match] = len(linkMatches)
				}
			}
		}

		// Linked Product (Cross Sell)
		crossSells := strings.Split(modelCsv.CrossSells, ",")
		for _, crossSell := range crossSells {
			crossSell = strings.TrimSpace(crossSell)
			linkIndex := prodIndices[crossSell] - 1
			if linkIndex >= 0 {
				match := fmt.Sprintf("%d:%d:1", prodIndex, linkIndex)
				if linkIndices[match] == 0 {
					modelLinks = append(modelLinks, models.Links{
						IsUpCross: utils.CrossSell,
					})
					linkMatches = append(linkMatches, match)
					linkIndices[match] = len(linkMatches)
				}
			}
		}
	}

	cateService := catesvc.NewServiceCategory(h.server.DB)
	err = cateService.CreateWithCSV(&modelCategories, cateNames, cateParents, cateIndices)
	if statusCode, message := errhandle.SqlErrorHandler(err); statusCode != 0 {
		return responses.ErrorResponse(c, statusCode, message)
	}

	tagService := tagsvc.NewServiceTag(h.server.DB)
	err = tagService.CreateWithCSV(&modelTags, tagNames, tagIndices)
	if statusCode, message := errhandle.SqlErrorHandler(err); statusCode != 0 {
		return responses.ErrorResponse(c, statusCode, message)
	}

	for index := range modelProducts {
		sku := modelProducts[index].Sku
		if sku[:3] == "id:" {
			sku = utils.CleanSpecialLetters(modelProducts[index].Title)
			modelProducts[index].Sku = sku
			prodSkus[index] = sku
			prodIndices[sku] = index + 1
		}
	}

	prodService := prodsvc.NewServiceProduct(h.server.DB)
	err = prodService.CreateWithCSV(&modelProducts, prodSkus, prodIndices)
	if statusCode, message := errhandle.SqlErrorHandler(err); statusCode != 0 {
		return responses.ErrorResponse(c, statusCode, message)
	}

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

	attrService := prodattrsvc.NewServiceAttribute(h.server.DB)
	err = attrService.CreateWithCSV(&modelAttrs, attrMatches, attrIndices)
	if statusCode, message := errhandle.SqlErrorHandler(err); statusCode != 0 {
		return responses.ErrorResponse(c, statusCode, message)
	}

	for index, match := range prodCateMatches {
		prodIndex := 0
		cateIndex := 0
		fmt.Sscanf(match, "%d:%d", &prodIndex, &cateIndex)
		prodID := modelProducts[prodIndex].ID
		cateID := modelCategories[cateIndex].ID
		match = fmt.Sprintf("%d:%d", prodID, cateID)
		modelProdCates[index].ProductID = uint64(prodID)
		modelProdCates[index].CategoryID = uint64(cateID)
		prodCateMatches[index] = match
		prodCateIndices[match] = index
	}

	prodCateService := prodcatesvc.NewServiceProductCategory(h.server.DB)
	err = prodCateService.CreateWithCSV(&modelProdCates, prodCateMatches, prodCateIndices)
	if statusCode, message := errhandle.SqlErrorHandler(err); statusCode != 0 {
		return responses.ErrorResponse(c, statusCode, message)
	}

	for index, match := range prodTagMatches {
		prodIndex := 0
		tagIndex := 0
		fmt.Sscanf(match, "%d:%d", &prodIndex, &tagIndex)
		prodID := modelProducts[prodIndex].ID
		tagID := modelTags[tagIndex].ID
		match = fmt.Sprintf("%d:%d", prodID, tagID)
		modelProdTags[index].ProductID = uint64(prodID)
		modelProdTags[index].TagID = uint64(tagID)
		prodTagMatches[index] = match
		prodTagIndices[match] = index
	}

	prodTagService := prodtagsvc.NewServiceProductTag(h.server.DB)
	err = prodTagService.CreateWithCSV(&modelProdTags, prodTagMatches, prodTagIndices)
	if statusCode, message := errhandle.SqlErrorHandler(err); statusCode != 0 {
		return responses.ErrorResponse(c, statusCode, message)
	}

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

	valService := prodattrvalsvc.NewServiceAttributeValue(h.server.DB)
	err = valService.CreateWithCSV(&modelVals, valMatches, valIndices)
	if statusCode, message := errhandle.SqlErrorHandler(err); statusCode != 0 {
		return responses.ErrorResponse(c, statusCode, message)
	}

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

	varService := prodvarsvc.NewServiceVariation(h.server.DB)
	err = varService.CreateWithCSV(&modelVars, varMatches, varIndices)
	if statusCode, message := errhandle.SqlErrorHandler(err); statusCode != 0 {
		return responses.ErrorResponse(c, statusCode, message)
	}

	for index := range modelShips {
		varIndex := modelShips[index].VariationID
		varID := modelVars[varIndex].ID
		modelShips[index].VariationID = uint64(varID)
		shipVarIDs = append(shipVarIDs, uint64(varID))
		shipIndices[uint64(varID)] = index
	}

	shipService := shipsvc.NewServiceShippingData(h.server.DB)
	err = shipService.CreateWithCSV(&modelShips, shipVarIDs, shipIndices)
	if statusCode, message := errhandle.SqlErrorHandler(err); statusCode != 0 {
		return responses.ErrorResponse(c, statusCode, message)
	}

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

	detService := prodvardetsvc.NewServiceVariationDetail(h.server.DB)
	err = detService.CreateWithCSV(&modelDets, detMatches, detIndices)
	if statusCode, message := errhandle.SqlErrorHandler(err); statusCode != 0 {
		return responses.ErrorResponse(c, statusCode, message)
	}

	for index, match := range linkMatches {
		prodIndex := 0
		linkIndex := 0
		sellType := 0
		fmt.Sscanf(match, "%d:%d:%d", &prodIndex, &linkIndex, &sellType)
		prodID := modelProducts[prodIndex].ID
		linkID := modelProducts[linkIndex].ID
		modelLinks[index].ProductID = uint64(prodID)
		modelLinks[index].LinkID = uint64(linkID)
		match = fmt.Sprintf("%d:%d:%d", prodID, linkID, sellType)
		linkMatches[index] = match
		linkIndices[match] = index
	}

	linkService := linksvc.NewServiceLink(h.server.DB)
	err = linkService.CreateWithCSV(&modelLinks, linkMatches, linkIndices)
	if statusCode, message := errhandle.SqlErrorHandler(err); statusCode != 0 {
		return responses.ErrorResponse(c, statusCode, message)
	}

	return responses.NewResponseProducts(c, http.StatusCreated, modelProducts)
}

// @Summary Get template of CSV file
// @Description Get template of CSV file
// @Tags Upload Actions
// @Accept multipart/form-data
// @Produce json
// @Security ApiKeyAuth
// @Success 200 {object} responses.Data
// @Router /store/api/v1/upload/csv [get]
func (h *HandlersUpload) GetTemplate(c echo.Context) error {
	return responses.Response(c, http.StatusOK, []models.CSVs{{}})
}
