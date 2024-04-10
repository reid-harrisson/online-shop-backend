package linksvc

import (
	"OnlineStoreBackend/models"
	"OnlineStoreBackend/pkgs/utils"
	"fmt"
)

func (service *Service) Create(productID uint64, linkID uint64, isUpCross utils.SellTypes) error {
	return service.DB.Where("link_id = ? And product_id = ?", linkID, productID).
		FirstOrCreate(&models.Links{
			ProductID: productID,
			LinkID:    linkID,
			IsUpCross: isUpCross,
		}).Error
}

func (service *Service) CreateWithCSV(modelNewLinks *[]models.Links, linkMatches []string, linkIndices map[string]int) {
	modelCurLinks := []models.Links{}
	service.DB.Where("Concat(product_id,':',link_id,':',is_up_cross) In (?)", linkMatches).Find(&modelCurLinks)
	for _, modelLink := range modelCurLinks {
		match := fmt.Sprintf("%d:%d:%d", modelLink.ProductID, modelLink.LinkID, modelLink.IsUpCross)
		index := linkIndices[match]
		(*modelNewLinks)[index].ID = modelLink.ID
	}
	service.DB.Save(modelNewLinks)
}
