package zonesvc

import (
	"OnlineStoreBackend/models"
	"OnlineStoreBackend/requests"
	"strconv"
)

func (service *Service) Create(storeID uint64, req *requests.RequestShippingZone, modelZones *[]models.ShippingZones) error {
	places := []string{}
	indices := map[string]int{}
	for index, place := range req.Places {
		places = append(places, strconv.FormatUint(storeID, 10)+":"+place)
		*modelZones = append(*modelZones, models.ShippingZones{
			StoreID: storeID,
			Name:    place,
		})
		indices[place] = index
	}

	modelNewZones := []models.ShippingZones{}
	service.DB.Where("Concat(store_id, ':', name) In (?)", places).Find(modelNewZones)
	for _, modelZone := range modelNewZones {
		index := indices[modelZone.Name]
		(*modelZones)[index].ID = modelZone.ID
	}
	return service.DB.Save(modelZones).Error
}
