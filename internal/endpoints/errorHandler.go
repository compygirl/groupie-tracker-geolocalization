package endpoints

import (
	"net/http"
	"text/template"

	models "groupieTrecker/internal/models"
)

func ErrorHandler(w http.ResponseWriter, errorNum int, errDetails error) {
	var resp models.ErrorResp
	resp.ErrorNum = errorNum
	resp.ErrorMessage = http.StatusText(errorNum) + "\n" + errDetails.Error()
	w.WriteHeader(errorNum)
	temp, err := template.ParseFiles("internal/templates/html_templates/errors.html")
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	err = temp.Execute(w, resp)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
}
