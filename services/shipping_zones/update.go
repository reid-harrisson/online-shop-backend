package zonesvc

import (
	"OnlineStoreBackend/models"
	"OnlineStoreBackend/requests"
	"strconv"
	"strings"
)

func (service *Service) Update(req *requests.RequestShippingZone, modelZone *models.ShippingZonesWithPlace) {
	modelZone.Name = req.Name
	modelZone.Description = req.Description

	service.DB.Save(&modelZone.ShippingZones)

	zoneID := uint64(modelZone.ID)
	places := []string{}
	indices := map[string]int{}
	modelPlaces := []models.ShippingPlaces{}
	for index, place := range req.Places {
		modelPlaces = append(modelPlaces, models.ShippingPlaces{
			ZoneID: zoneID,
			Name:   place,
		})
		places = append(places, strconv.FormatUint(zoneID, 10)+":"+place)
		indices[place] = index
	}

	modelNewPlaces := []models.ShippingPlaces{}
	service.DB.Where("Concat(zone_id, ':', name) In (?)", places).Find(&modelNewPlaces)
	service.DB.Where("Concat(zone_id, ':', name) Not In (?)", places).Delete(&models.ShippingPlaces{})
	for _, modelPlace := range modelNewPlaces {
		index := indices[modelPlace.Name]
		(modelPlaces)[index].ID = modelPlace.ID
	}
	service.DB.Save(modelPlaces)
	placeIDs := []string{}
	for _, modelPlace := range modelPlaces {
		placeIDs = append(placeIDs, strconv.FormatUint(uint64(modelPlace.ID), 10))
	}
	modelZone.PlaceIDs = strings.Join(placeIDs, ",")
	modelZone.PlaceNames = strings.Join(req.Places, ",")
}
