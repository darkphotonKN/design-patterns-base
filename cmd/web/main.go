package main

import (
	"log"
	"net/http"
)

const port string = ":4000"

type application struct{}

func main() {

	// app := application{}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		res := "testing root reply!"
		byteRes := []byte(res)

		w.Write(byteRes)
	})

	log.Println("Server starting and listening on ", port)

	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Panic("Error when attempting to start server:", err)
	}

}
