package controller

import "net/http"

//Dashboard for user
func Dashboard(w http.ResponseWriter, r *http.Request) {
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

	//displayTemplate(w, "dashboard", nil)
}
