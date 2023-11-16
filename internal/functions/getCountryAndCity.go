package groupieTrecker

import (
	models "groupieTrecker/internal/models"
)

func GetCountry(artistsList []models.Artist) map[string]bool {
	countryMap := make(map[string]bool, 0)
	for _, artist := range artistsList {
		for country, _ := range artist.DatesLocation {
			countryMap[country] = true
		}
	}

	return countryMap
}

func GetCity(artistsList []models.Artist) map[string]bool {
	cityMap := make(map[string]bool, 0)
	for _, artist := range artistsList {
		for _, mapa := range artist.DatesLocation {
			for city, _ := range mapa {
				cityMap[city] = true
			}
		}
	}
	return cityMap
}
