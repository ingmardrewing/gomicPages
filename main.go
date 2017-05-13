package main

import (
	"log"
	"net/http"

	restful "github.com/emicklei/go-restful"
	"github.com/ingmardrewing/gomicRest/db"
	"github.com/ingmardrewing/gomicRest/gomicRest"
)

func main() {
	db.Initialize()

	restful.Add(gomicRest.New())

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err.Error())
	}
}
