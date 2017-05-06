package controller

import (
	"fmt"
	"net/http"

	"github.com/Shivakishore14/My-cloud-core/app/console"
	"github.com/Shivakishore14/My-cloud-core/app/model"

	lxc "github.com/lxc/go-lxc"
)

//ListConatiners : for listing containers for user
func ListConatiners(w http.ResponseWriter, r *http.Request) {
	session, err := store.Get(r, "session-name")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	user, ok := session.Values["user"]
	if !ok {
		//http.Redirect(w, r, "/login", http.StatusUnauthorized)
		webresponse("Unauthorised", nil, nil, w)
		return
	}
	list := []model.Container{}
	if gobj := db.Where("created_by=?", user).Find(&list); gobj.Error != nil {
		fmt.Printf("Error %s", gobj.Error.Error())
	} else {
		container := []model.ContainerBasicInfo{}
		for _, i := range list {
			obj := model.ContainerBasicInfo{}
			obj.Name = i.Name
			obj.CreatedBy = user.(string)
			obj.DisplayName = i.DisplayName
			obj.Status = i.Status
			if i.Status == "done" {
				if c, err := lxc.NewContainer(i.Name, lxcpath); err != nil {
					console.PrintError("Error :" + err.Error())
				} else {
					obj.IsRunning = c.Running()
					if obj.IsRunning {
						if memUsed, e := c.MemoryUsage(); e != nil {
							console.PrintError("Error Getting memory details" + e.Error())
						} else {
							obj.MemUsed = memUsed.String()
						}

						if memLimit, e := c.MemoryLimit(); e != nil {
							console.PrintError("Error Getting memory details" + e.Error())
						} else {
							obj.MemLimit = memLimit.String()
						}

						if ip, e := c.IPAddresses(); e != nil {
							console.PrintError("Error geting ip :" + e.Error())
						} else {
							obj.IPAddress = ip
						}

					}
				}
			}
			container = append(container, obj)
		}
		webresponse("success", nil, container, w)
	}

}
