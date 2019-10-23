package main

import (
	"io"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func CreateUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	io.WriteString(w, "Create User Handler...")
}

func Login(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	username := p.ByName("username")
	io.WriteString(w, "Login..."+username)
}
