package zonesvc

import (
	"OnlineStoreBackend/models"
	"OnlineStoreBackend/requests"
	"strconv"
	"strings"
)

func (service *Service) Create(storeID uint64, req *requests.RequestShippingZone, modelZone *models.ShippingZonesWithPlace) error {
	err := service.DB.Where("name = ?", req.Name).First(&modelZone.ShippingZones).Error
	modelZone.StoreID = storeID
	modelZone.Name = req.Name
	modelZone.Description = req.Description
	if err != nil {
		if err = service.DB.Create(modelZone).Error; err != nil {
			return err
		}
	}
	if err = service.DB.Save(modelZone).Error; err != nil {
		return err
	}

	zoneID := uint64(modelZone.ID)
	places := []string{}
	indices := map[string]int{}
	modelPlaces := []models.ShippingLocations{}
	for index, place := range req.Places {
		if indices[place] == 0 {
			modelPlaces = append(modelPlaces, models.ShippingLocations{
				ZoneID: zoneID,
				Name:   place,
			})
			places = append(places, place)
			indices[place] = index + 1
		}
	}

	modelNewPlaces := []models.ShippingLocations{}
	service.DB.Where("zone_id = ? And name In (?)", zoneID, places).Find(&modelNewPlaces)
	service.DB.Where("zone_id = ? And name Not In (?)", zoneID, places).Delete(&models.ShippingLocations{})
	for _, modelPlace := range modelNewPlaces {
		index := indices[modelPlace.Name] - 1
		(modelPlaces)[index].ID = modelPlace.ID
	}
	service.DB.Save(modelPlaces)
	placeIDs := []string{}
	for _, modelPlace := range modelPlaces {
		placeIDs = append(placeIDs, strconv.FormatUint(uint64(modelPlace.ID), 10))
	}
	modelZone.PlaceIDs = strings.Join(placeIDs, ",")
	modelZone.PlaceNames = strings.Join(req.Places, ",")
	return nil
}
