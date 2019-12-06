package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func RegisterHandler() *httprouter.Router {
	router := httprouter.New()
	router.GET("/", homeHandler)
	router.POST("/", homeHandler)

	router.ServeFiles("/static/*filepath", http.Dir("./templates"))
	return router
}

func main() {
	r := RegisterHandler()
	http.ListenAndServe(":8080", r)
}
