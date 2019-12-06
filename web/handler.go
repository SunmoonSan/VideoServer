package main

import (
	"github.com/julienschmidt/httprouter"
	"html/template"
	"log"
	"net/http"
)

type HomePage struct {
	Name string
}

func homeHandler(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	p := &HomePage{Name: "San"}
	t, e := template.ParseFiles("./template/home.html")
	if e != nil {
		log.Printf("Parsing template homt.html error: %s", e)
		return
	}
	t.Execute(w, p)
	return
}
