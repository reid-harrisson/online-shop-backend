package ordsvc

import (
	"OnlineStoreBackend/models"
	"OnlineStoreBackend/repositories"
)

func (service *Service) Create(modelOrders *[]models.ProductOrders, modelCarts []models.CartItemsWithDetail, modelTaxSet models.TaxSettings, customerID uint64) error {
	id := uint(0)
	prodRepo := repositories.NewRepositoryProduct(service.DB)
	modelProduct := models.Products{}
	for i, modelCart := range modelCarts {
		modelOrder := models.ProductOrders{
			StoreID:       modelCart.StoreID,
			CustomerID:    modelCart.CustomerID,
			ProductID:     modelCart.ProductID,
			UnitPriceSale: modelCart.UnitPrice,
			Quantity:      modelCart.Quantity,
			SubTotal:      modelCart.UnitPrice,
			TaxRate:       modelTaxSet.TaxRate,
			TaxAmount:     modelTaxSet.TaxRate * modelCart.UnitPrice / 100,
			Total:         modelTaxSet.TaxRate*modelCart.UnitPrice/100 + modelCart.UnitPrice,
			Status:        "Pending",
		}
		if i == 0 {
			service.DB.Create(&modelOrder)
			id = modelOrder.ID
		} else {
			modelOrder.ID = id
			service.DB.Create(&modelOrder)
		}
		*modelOrders = append(*modelOrders, modelOrder)
		prodRepo.ReadByID(&modelProduct, modelCart.ProductID)
		service.DB.Model(&modelProduct).Update("stock_quantity", modelProduct.StockQuantity-modelCart.Quantity)
		service.DB.Delete(&modelCart)
	}
	return nil
}
