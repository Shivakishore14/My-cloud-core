package controller

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/Shivakishore14/My-Cloud/app/console"
	"github.com/Shivakishore14/My-Cloud/app/model"
	//mysql driver
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/sessions"
	"github.com/jinzhu/gorm"
	lxc "github.com/lxc/go-lxc"
)

var database = "cloud"
var user = "test"
var password = "test"

var db *gorm.DB
var store *sessions.CookieStore

var (
	post       = "POST"
	get        = "GET"
	lxcpath    = lxc.DefaultConfigPath()
	ctemplate  = "ubuntu"
	distro     = "ubuntu"
	release    = "xenial"
	arch       = "amd64"
	verbose    = true
	flush      = false
	validation = false
)

func init() {
	var err error

	store = sessions.NewCookieStore([]byte("secret-key-to-be-changed"))
	if db, err = gorm.Open("mysql", user+":"+password+"@/"+database+"?charset=utf8&parseTime=True&loc=Local"); err != nil {
		console.PrintError("Error Connecting to database")
	}

	if db.HasTable(&model.User{}) == false {
		db.CreateTable(&model.User{})
	}
	if db.HasTable(&model.Container{}) == false {
		db.CreateTable(&model.Container{})
	}
	//f := model.Container{CreatedBy: "sk1", Name: "Myubuntu11", DisplayName: "TestUbuntu1"}
	//a := db.Create(&f)
	//fmt.Println(a.Error)
	//fmt.Println(reflect.TypeOf(a))
}

func webresponse(msg string, err error, data interface{}, w http.ResponseWriter) (string, error) {
	obj := model.WebResponse{}
	obj.Message = msg
	obj.Error = err
	obj.Data = data

	var resErr error
	var resTxt string

	if jsonData, e := json.Marshal(obj); e != nil {
		resErr = e
	} else {
		resTxt = string(jsonData)
	}
	fmt.Fprint(w, resTxt)
	return resTxt, resErr
}

func displayTemplate(w http.ResponseWriter, templateName string, data interface{}) {
	var templates = template.Must(template.ParseGlob("templates/*"))
	if err := templates.ExecuteTemplate(w, templateName, data); err != nil {
		fmt.Fprint(w, "Error in template")
		log.Println(err)
	}
}
