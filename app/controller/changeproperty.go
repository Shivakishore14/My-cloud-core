package controller

import (
	"net/http"
	"time"

	"github.com/Shivakishore14/My-Cloud/app/console"
	"github.com/Shivakishore14/My-Cloud/app/model"

	lxc "github.com/lxc/go-lxc"
)

func changeProp(w http.ResponseWriter, r *http.Request) {
	prop := r.FormValue("property")

	displayName := r.FormValue("displayName")
	container := model.Container{}
	db.Where("display_name=?", displayName).Find(&container)

	if prop == "DisplayName" {
		newDisplayName := r.FormValue("newDisplayName")
		container.DisplayName = newDisplayName
		db.Save(container)
		webresponse("success", nil, nil, w)
		return
	}

	//properties below use container object
	c, err := lxc.NewContainer(container.Name, lxcpath)
	if err != nil {
		console.PrintError("ERROR: " + err.Error())
		webresponse("Failed", err, nil, w)
		return
	}

	if prop == "destroy" {

		console.PrintSuccess("Destroying container... : " + container.Name)
		if err := c.Destroy(); err != nil {
			console.PrintError("ERROR: " + err.Error())
			webresponse("Failed", err, nil, w)
			return
		}
		webresponse("Success", nil, nil, w)
		return
	}

	if prop == "start" {
		console.PrintSuccess("starting container... : " + container.Name)

		if err := c.Start(); err != nil {
			console.PrintError("ERROR: " + err.Error())
			webresponse("Failed", err, nil, w)
			return
		}
		console.PrintSuccess("Wating for container to start networking.. : " + container.Name)
		if _, err := c.WaitIPAddresses(5 * time.Second); err != nil {
			console.PrintError("ERROR: " + err.Error())
			webresponse("Failed Networking Error", err, nil, w)
			return
		}

		webresponse("Success", nil, nil, w)
		return
	}

	if prop == "stop" {
		c.SetLogFile("/tmp/" + container.Name + ".log")
		c.SetLogLevel(lxc.TRACE)
		console.PrintSuccess("Stopping container.. : " + container.Name)

		if err := c.Stop(); err != nil {
			console.PrintError("ERROR: " + err.Error())
			webresponse("Failed", err, nil, w)
			return
		}
		webresponse("Success", nil, nil, w)
		return
	}

	if prop == "shutdown" {
		console.PrintSuccess("Shutting down container.. : " + container.Name)
		if err := c.Shutdown(30 * time.Second); err != nil {
			console.PrintError("ERROR: " + err.Error())
			webresponse("Failed", err, nil, w)
			return
		}
		webresponse("Success", nil, nil, w)
		return
	}

}
