package controller

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

//UITesting :: For UI testing
func UITesting(w http.ResponseWriter, r *http.Request) {
	templates := template.Must(template.ParseGlob("templates/*"))
	vars := mux.Vars(r)
	path := vars["path"]

	if err := templates.ExecuteTemplate(w, path, nil); err != nil {
		fmt.Fprint(w, "Error in template")
		log.Println(err)
	}
}
