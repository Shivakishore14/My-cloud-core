package controller

import (
	"fmt"
	"net/http"
	"os/exec"

	"github.com/Shivakishore14/My-Cloud/app/console"
	"github.com/Shivakishore14/My-Cloud/app/model"
)

//CreateContainer : to create a container
func CreateContainer(w http.ResponseWriter, r *http.Request) {
	session, err := store.Get(r, "session-name")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	user := session.Values["user"]

	if user == user {
		//	http.Redirect(w, r, "/login", http.StatusUnauthorized)
		//	return
	}

	if r.Method == post {
		/*name := r.FormValue("name")
		f := model.Container{CreatedBy: "sk", Name: name, DisplayName: "TestUbunt1", Status: "Creating"}

		if obj := db.Create(f); obj.Error != nil {
			console.PrintError("Error Writing to DB")
			return
		}*/
		fmt.Fprint(w, "Started")
		go create(r, user.(string))
	}
}

func create(r *http.Request, user string) {

	name := r.FormValue("name")
	cargs := []string{"lxc-create", "--vgname=lxc-volume-group", "-B", "lvm", "--fssize", "2011M", "-t", "ubuntu", "-n", name}
	cargs = append(cargs)
	output, err := exec.Command(cargs[0], cargs[1:]...).CombinedOutput()
	if err != nil {
		console.PrintError("Failed Creation")
		return
	}
	fmt.Print(string(output))
	con := model.Container{}
	db.Where("created_by = ? and name = ?", user, name).First(&con).Update("status", "done")
}
