package controller

import (
	"net/http"

	"github.com/Shivakishore14/My-cloud-core/app/model"
)

//SignUp for user
func SignUp(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	password := r.FormValue("password")
	name := r.FormValue("name")
	email := r.FormValue("email")
	phone := r.FormValue("phone")

	users := model.User{UserName: username, Password: password, Name: name, Email: email, Phone: phone}
	if gobj := db.Create(&users); gobj.Error != nil {
		webresponse("failed", gobj.Error, nil, w)
	} else {
		webresponse("success", nil, nil, w)
	}
}
