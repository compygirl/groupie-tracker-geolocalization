package groupieTrecker

import (
	"log"
	"strings"

	models "groupieTrecker/internal/models"
)

func UnmarshalArtistsAndRelations() ([]models.Artist, error) {
	artistsList, err := GetArtistsList()
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}
	relationlist, err := GetRelations()
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}
	tempMaps := make([]map[string][]string, 0)
	for i := 0; i < len(artistsList); i++ {
		tempMaps = append(tempMaps, relationlist.Index[i].DatesLocation)
	}
	// now we will fill our map of maps so that the key of map of json(state+conutry) will changed and our map will be as below:
	for i := 0; i < len(tempMaps); i++ {
		artistsList[i].DatesLocation = make(map[string]map[string][]string)
		for location, dates := range tempMaps[i] {
			innerMap := make(map[string][]string, 0)
			tempLoc := strings.Split(location, "-") // before: texas-usa, after: [texas,usa]
			if val, ok := artistsList[i].DatesLocation[tempLoc[1]]; ok {
				val[tempLoc[0]] = dates
			} else {
				innerMap[tempLoc[0]] = dates
				artistsList[i].DatesLocation[tempLoc[1]] = innerMap
			}

		}
	}
	return artistsList, nil
}
