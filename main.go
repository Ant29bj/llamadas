package main

import (
	"fmt"
	"llamadas/controller"
	"llamadas/setup"
	"log"
	"net/http"
)

func init() {
	setup.SetDatabase()
}

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "hello World")
	})

	http.HandleFunc("/llamadas", controller.GetAllLlamadas)

	log.Print("Starting server")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
