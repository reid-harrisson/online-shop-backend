package prodvarsvc

import (
	"OnlineStoreBackend/models"
	"OnlineStoreBackend/repositories"
	"OnlineStoreBackend/requests"
)

func (service *Service) Update(attributeID uint64, productID uint64, modelVars *[]models.ProductVariationsWithName, req *requests.RequestProductVariation) {
	filterKeys := make(map[string]int)
	for _, modelVar := range *modelVars {
		filterKeys[modelVar.Variant] = 1
	}
	for _, variant := range req.Variants {
		if filterKeys[variant] == 1 {
			filterKeys[variant] = 3
		} else {
			filterKeys[variant] = 2
		}
	}

	for variant, key := range filterKeys {
		if key == 1 {
			service.Delete(variant, productID)
		} else if key == 2 {
			service.Create(attributeID, variant, productID)
		}
	}

	varRepo := repositories.NewRepositoryProductVariation(service.DB)
	varRepo.ReadByID(modelVars, attributeID)
}
