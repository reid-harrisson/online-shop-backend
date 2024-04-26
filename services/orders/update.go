package ordsvc

import (
	"OnlineStoreBackend/models"
	"OnlineStoreBackend/pkgs/utils"
	stocksvc "OnlineStoreBackend/services/stock_trails"
)

func (service *Service) UpdateStatus(modelItems *[]models.OrderItems, storeID uint64, orderID uint64, orderStatus string) {
	status := utils.OrderStatusFromString(orderStatus)
	service.DB.Where("order_id = ?", orderID).Find(&modelItems)

	modelVars := []models.Variations{}
	varIDs := []uint64{}
	varIndices := map[uint64]int{}

	for _, modelItem := range *modelItems {
		if varIndices[modelItem.VariationID] == 0 {
			varIDs = append(varIDs, modelItem.VariationID)
			varIndices[modelItem.VariationID] = len(varIDs)
		}
	}
	service.DB.Where("id In (?)", varIDs).Find(&modelVars)

	modelStocks := []models.StockTrails{}
	stockService := stocksvc.NewServiceStockTrail(service.DB)

	flagCompleted := true
	flagPending := true
	for index := range *modelItems {
		modelItem := &(*modelItems)[index]
		varIndex := varIndices[modelItem.VariationID] - 1
		if modelItem.StoreID == storeID && modelItem.Status != status {
			modelItem.Status = status
			if status == utils.StatusOrderProcessing {
				modelVars[varIndex].StockLevel -= modelItem.Quantity
				modelStocks = append(modelStocks, models.StockTrails{
					VariationID: uint64(modelVars[varIndex].ID),
					ProductID:   modelVars[varIndex].ProductID,
					Change:      -modelItem.Quantity,
					Event:       utils.OrderPlaced,
				})
			} else if status == utils.StatusOrderCancelled {
				modelVars[varIndex].StockLevel += modelItem.Quantity
				modelStocks = append(modelStocks, models.StockTrails{
					VariationID: uint64(modelVars[varIndex].ID),
					ProductID:   modelVars[varIndex].ProductID,
					Change:      modelItem.Quantity,
					Event:       utils.OrderCancelled,
				})
			}
		}
		if modelItem.Status != utils.StatusOrderCompleted {
			flagCompleted = false
		}
		if modelItem.Status != utils.StatusOrderPending {
			flagPending = false
		}
	}

	service.DB.Save(modelItems)
	service.DB.Save(&modelVars)
	stockService.CreateStocks(&modelStocks)

	if flagCompleted {
		status = utils.StatusOrderCompleted
	} else if flagPending {
		status = utils.StatusOrderPending
	} else {
		status = utils.StatusOrderProcessing
	}
	service.DB.Model(models.Orders{}).
		Where("id = ?", orderID).
		Update("status", status)
}

func (service *Service) UpdateOrderItemStatus(orderID uint64, status string) {
	service.DB.
		Model(models.OrderItems{}).
		Where("order_id = ?", orderID).
		Update("status", utils.OrderStatusFromString(status))
}
