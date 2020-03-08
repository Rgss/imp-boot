package webserver

type WebServer interface {
	Run()
	Create()
	RegisterRouters()
}