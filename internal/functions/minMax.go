package groupieTrecker

import data "groupieTrecker/internal/data"

func MinMax() (int, int) {
	min := data.ArtistsList[0].CreationDate
	max := data.ArtistsList[0].CreationDate
	for _, artist := range data.ArtistsList {
		if artist.CreationDate > max {
			max = artist.CreationDate
		}
		if artist.CreationDate < min {
			min = artist.CreationDate
		}
	}
	return min, max
}
