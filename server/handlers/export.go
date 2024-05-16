package handlers

import (
	"OnlineStoreBackend/pkgs/constants"
	"OnlineStoreBackend/repositories"
	"OnlineStoreBackend/responses"
	s "OnlineStoreBackend/server"
	"encoding/csv"
	"net/http"
	"os"
	"strconv"

	"github.com/labstack/echo/v4"
)

type HandlersExport struct {
	server *s.Server
}

func NewHandlersExport(server *s.Server) *HandlersExport {
	return &HandlersExport{server: server}
}

// @Summary Export a CSV file
// @Description Export a CSV file to the server
// @Tags Export Actions
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param store_id query int true "Store ID"
// @Success 201 {object} responses.Data
// @Success 400 {object} responses.Error
// @Success 500 {object} responses.Error
// @Router /store/api/v1/export/csv [get]
func (h *HandlersExport) ExportCSV(c echo.Context) error {
	storeID, _ := strconv.ParseUint(c.QueryParam("store_id"), 10, 64)
	userID, err := strconv.ParseUint(c.Request().Header.Get("id"), 10, 64)
	if userID == 0 || err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, constants.InvalidData)
	}

	exports := [][]string{}
	exportRepo := repositories.NewRepositoryExport(h.server.DB)
	err = exportRepo.ReadAll(&exports, storeID)
	if err != nil {
		return responses.ErrorResponse(c, http.StatusInternalServerError, constants.InternalServerErrorMessage)
	}

	// Open export.csv
	file, err := os.Create("export.csv")
	if err != nil {
		return responses.ErrorResponse(c, http.StatusInternalServerError, constants.InternalServerErrorMessage)
	}
	defer file.Close()

	writer := csv.NewWriter(file)

	for _, export := range exports {
		err := writer.Write(export)
		if err != nil {
			return responses.ErrorResponse(c, http.StatusInternalServerError, constants.InternalServerErrorMessage)
		}
	}

	writer.Flush()

	return c.Attachment("export.csv", "export.csv")
}
