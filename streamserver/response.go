/*
@desc : Created by San on 2019/10/27 10:50
*/
package main

import (
	"io"
	"net/http"
)

func sendErrorResponse(w http.ResponseWriter, sc int, errMsg string) {
	w.WriteHeader(sc)
	io.WriteString(w, errMsg)
}
