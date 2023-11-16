package endpoints

import (
	"errors"
	"fmt"
	"net/http"
	"text/template"
)

func RenderTemplate(w http.ResponseWriter, htmltemplate string, resp interface{}) {
	temp, err := template.ParseFiles(htmltemplate)
	if err != nil {
		ErrorHandler(w, http.StatusInternalServerError, errors.New("problems with parsing"))
		fmt.Println(err)
		return
	}

	err = temp.Execute(w, resp)
	if err != nil {
		ErrorHandler(w, http.StatusInternalServerError, errors.New("problems with executing"))
		fmt.Println(err)
		return
	}
}
