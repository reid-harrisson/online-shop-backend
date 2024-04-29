package zonesvc

import (
	"OnlineStoreBackend/models"
	"OnlineStoreBackend/requests"
	"strconv"
	"strings"
)

func (service *Service) Update(req *requests.RequestShippingZone, modelZone *models.ShippingZonesWithPlace) error {
	modelZone.Name = req.Name
	modelZone.Description = req.Description

	if err := service.DB.Save(&modelZone.ShippingZones).Error; err != nil {
		return err
	}

	zoneID := uint64(modelZone.ID)
	places := []string{}
	indices := map[string]int{}
	modelPlaces := []models.ShippingLocations{}

	for index, place := range req.Places {
		modelPlaces = append(modelPlaces, models.ShippingLocations{
			ZoneID: zoneID,
			Name:   place,
		})
		places = append(places, strconv.FormatUint(zoneID, 10)+":"+place)
		indices[place] = index
	}

	modelNewPlaces := []models.ShippingLocations{}

	if err := service.DB.Where("Concat(zone_id, ':', name) In (?)", places).Find(&modelNewPlaces).Error; err != nil {
		return err
	}

	if err := service.DB.Where("Concat(zone_id, ':', name) Not In (?)", places).Delete(&models.ShippingLocations{}).Error; err != nil {
		return err
	}

	for _, modelPlace := range modelNewPlaces {
		index := indices[modelPlace.Name]
		(modelPlaces)[index].ID = modelPlace.ID
	}

	if err := service.DB.Save(modelPlaces).Error; err != nil {
		return err
	}

	placeIDs := []string{}
	for _, modelPlace := range modelPlaces {
		placeIDs = append(placeIDs, strconv.FormatUint(uint64(modelPlace.ID), 10))
	}

	modelZone.PlaceIDs = strings.Join(placeIDs, ",")
	modelZone.PlaceNames = strings.Join(req.Places, ",")

	return nil
}
