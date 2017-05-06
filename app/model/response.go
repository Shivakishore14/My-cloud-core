package model

//WebResponse is the structure used for web api response
type WebResponse struct {
	Message string
	Error   error
	Data    interface{}
}
