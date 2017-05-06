package controller

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func Settings(w http.ResponseWriter, r *http.Request) {
	session, err := store.Get(r, "session-name")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	_, ok := session.Values["user"]
	if !ok {
		http.Redirect(w, r, "/login", 302)
		return
	}

	vars := mux.Vars(r)
	name := vars["name"]
	displayTemplate(w, "settings", nil)
	fmt.Print(name)
}
