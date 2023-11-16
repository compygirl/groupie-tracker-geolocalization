package endpoints

import (
	"encoding/json"
	"errors"
	"net/http"

	funcs "groupieTrecker/internal/functions"
)

func SearchHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/search" {
		ErrorHandler(w, http.StatusNotFound, errors.New(" "))
		return
	}
	switch r.Method {
	case "GET":
		searchTerm := r.FormValue("searchTerm")
		chekedMembersNumsStr := r.FormValue("membersNums")
		firstAlbumDate := r.FormValue("FAD")
		selectedCountry := r.FormValue("country")
		selectedCity := r.FormValue("city")
		rangeVal1 := r.FormValue(("CD1"))
		rangeVal2 := r.FormValue(("CD2"))
		suggestions, set, err := funcs.SearchSuggestions(searchTerm, rangeVal1, rangeVal2, chekedMembersNumsStr, firstAlbumDate, selectedCountry, selectedCity)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		responseMap := map[string]interface{}{
			"suggestions": suggestions,
			"set":         set,
		}

		response, err := json.Marshal(responseMap)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(response)
		return
	default:
		ErrorHandler(w, http.StatusMethodNotAllowed, errors.New(" "))
		return
	}
}
