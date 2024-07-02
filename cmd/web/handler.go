package main

import "net/http"

// ShowHome Handler
func (app *application) ShowHome(w http.ResponseWriter, r *http.Request) {

	res := "Show home."
	resBytes := []byte(res)

	w.Write(resBytes)
}
