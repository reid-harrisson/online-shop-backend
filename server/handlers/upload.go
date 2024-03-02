package handlers

import (
	"OnlineStoreBackend/models"
	"OnlineStoreBackend/pkgs/utils"
	"OnlineStoreBackend/responses"
	s "OnlineStoreBackend/server"
	linksvc "OnlineStoreBackend/services/links"
	prodsvc "OnlineStoreBackend/services/products"
	"encoding/csv"
	"encoding/json"
	"io"
	"net/http"
	"os"
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

func readCSV(filename string, modelCsvs *[]models.CSVs) {
	file, err := os.Open(filename)
	if err != nil {
		return
	}
	defer file.Close()

	reader := csv.NewReader(file)

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
	file, _ := c.FormFile("file")

	src, err := file.Open()
	if err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, err.Error())
	}
	defer src.Close()

	dst, err := os.Create("./uploads/" + file.Filename)
	if err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, err.Error())
	}
	defer dst.Close()

	if _, err = io.Copy(dst, src); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	modelCsvs := make([]models.CSVs, 0)
	readCSV("./uploads/"+file.Filename, &modelCsvs)

	mapSku := make(map[string]uint64)
	mapIDs := make(map[string]string)

	modelProducts := make([]models.Products, 0)
	prodService := prodsvc.NewServiceProduct(h.server.DB)
	for _, modelCsv := range modelCsvs {
		modelProduct := models.Products{}
		prodService.CreateWithCSV(&modelProduct, modelCsv, storeID, &mapIDs)
		if modelProduct.ID != 0 {
			modelProducts = append(modelProducts, modelProduct)
			mapSku[modelProduct.Sku] = uint64(modelProduct.ID)
			mapIDs[modelCsv.ID] = modelProduct.Sku
		}
	}

	linkService := linksvc.NewServiceProductLinked(h.server.DB)
	for _, modelCsv := range modelCsvs {
		upSells := strings.Split(modelCsv.Upsells, ",")
		crossSells := strings.Split(modelCsv.CrossSells, ",")
		sku := strings.TrimSpace(modelCsv.Sku)
		if modelCsv.Parent != "" {
			sku = strings.TrimSpace(modelCsv.Sku)
		}
		for _, upSell := range upSells {
			if len(upSell) > 3 && upSell[:3] == "id:" {
				id := upSell[3:]
				upSell = mapIDs[id]
			}
			if mapSku[sku] != 0 && mapSku[upSell] != 0 {
				linkService.Create(mapSku[sku], mapSku[upSell], utils.UpSell)
			}
		}
		for _, crossSell := range crossSells {
			if len(crossSell) > 3 && crossSell[:3] == "id:" {
				id := crossSell[3:]
				crossSell = mapIDs[id]
			}
			if mapSku[sku] != 0 && mapSku[crossSell] != 0 {
				linkService.Create(mapSku[sku], mapSku[crossSell], utils.UpSell)
			}
		}
	}

	return responses.NewResponseProducts(c, http.StatusBadRequest, modelProducts)
}
