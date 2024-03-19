package contsvc

import (
	"OnlineStoreBackend/models"
	"OnlineStoreBackend/requests"
	"fmt"
)

func (service *Service) Update(productID uint64, req *requests.RequestProductContent) {
	modelNewConts := []models.ProductContents{}
	modelCurConts := []models.ProductContents{}
	contIndices := map[string]int{}
	contMatches := []string{}
	for index, contentID := range req.ContentIDs {
		match := fmt.Sprintf("%d:%d", productID, contentID)
		contMatches = append(contMatches, match)
		contIndices[match] = index
		modelNewConts = append(modelNewConts, models.ProductContents{
			ProductID: productID,
			ContentID: contentID,
		})
	}
	service.DB.Where("Concat(product_id, ':', content_id) In (?)", contMatches).Find(&modelCurConts)
	service.DB.Where("Concat(product_id, ':', content_id) Not In (?) And product_id = ?", contMatches, productID).Delete(&models.ProductContents{})
	for _, modelCont := range modelCurConts {
		match := fmt.Sprintf("%d:%d", modelCont.ProductID, modelCont.ContentID)
		index := contIndices[match]
		modelNewConts[index].ID = modelCont.ID
	}
	service.DB.Save(&modelNewConts)
}
