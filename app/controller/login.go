package controller

import (
	"fmt"
	"net/http"

	"github.com/Shivakishore14/My-cloud-core/app/model"
)

//Login for user
func Login(w http.ResponseWriter, r *http.Request) {

	username := r.FormValue("username")
	password := r.FormValue("password")
	session, err := store.Get(r, "session-name")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Print(session.Values["user"])
	users := model.User{}
	db.Where("user_name=?", username).First(&users)

	if users.Password == password {
		session.Values["user"] = username
		session.Save(r, w)
		//webresponse("success", nil, nil, w)
		http.Redirect(w, r, "dashboard", 302)
	} else {
		webresponse("failed", nil, nil, w)
	}
}
