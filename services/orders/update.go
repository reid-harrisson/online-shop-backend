package ordsvc

import (
	"OnlineStoreBackend/models"
	"OnlineStoreBackend/pkgs/utils"
	stocksvc "OnlineStoreBackend/services/stock_trails"
)

func (service *Service) UpdateStatus(modelItems *[]models.OrderItems, storeID uint64, orderID uint64, orderStatus string) error {
	status := utils.OrderStatusFromString(orderStatus)
	err := service.DB.Where("order_id = ?", orderID).Find(&modelItems).Error
	if err != nil {
		return err
	}

	modelVars := []models.Variations{}
	varIDs := []uint64{}
	varIndices := map[uint64]int{}

	for _, modelItem := range *modelItems {
		if varIndices[modelItem.VariationID] == 0 {
			varIDs = append(varIDs, modelItem.VariationID)
			varIndices[modelItem.VariationID] = len(varIDs)
		}
	}

	err = service.DB.Where("id In (?)", varIDs).Find(&modelVars).Error
	if err != nil {
		return err
	}

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

	err = service.DB.Save(modelItems).Error
	if err != nil {
		return err
	}

	err = service.DB.Save(&modelVars).Error
	if err != nil {
		return err
	}

	err = stockService.CreateStocks(&modelStocks)
	if err != nil {
		return err
	}

	if flagCompleted {
		status = utils.StatusOrderCompleted
	} else if flagPending {
		status = utils.StatusOrderPending
	} else {
		status = utils.StatusOrderProcessing
	}

	return service.DB.Model(models.Orders{}).
		Where("id = ?", orderID).
		Update("status", status).
		Error
}

func (service *Service) UpdateOrderItemStatus(orderID uint64, status string) error {
	return service.DB.
		Model(models.OrderItems{}).
		Where("order_id = ?", orderID).
		Update("status", utils.OrderStatusFromString(status)).
		Error
}
