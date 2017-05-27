package main

import (
	"log"
	"net/http"

	restful "github.com/emicklei/go-restful"
	"github.com/ingmardrewing/gomicPages/db"
	"github.com/ingmardrewing/gomicPages/service"
)

func main() {
	db.Initialize()

	restful.Add(service.NewPagesService())

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err.Error())
	}
}
