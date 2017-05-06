package routes

import (
	"net/http"

	"github.com/Shivakishore14/My-Cloud/app/console"
	"github.com/Shivakishore14/My-Cloud/app/controller"

	"github.com/gorilla/mux"
)

//LoadRoutes :for loading routing
func LoadRoutes() {
	r := mux.NewRouter()
	r.HandleFunc("/login", controller.Login)
	r.HandleFunc("/signup", controller.SignUp)
	r.HandleFunc("/dashboard", controller.Dashboard)
	r.HandleFunc("/settings/{name}", controller.Settings)
	r.HandleFunc("/createcontainer", controller.CreateContainer)
	r.HandleFunc("/listcontainers", controller.ListConatiners)
	r.HandleFunc("/getstats", controller.GetStatsContainer)
	r.HandleFunc("/test", controller.Templatetest)
	r.HandleFunc("/UI/{path}", controller.UITesting)
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./static/")))
	console.PrintSuccess("Listening on 9090")
	console.PrintError(http.ListenAndServe(":9090", r))
}
