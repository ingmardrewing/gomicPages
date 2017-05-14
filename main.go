package main

import (
	"log"
	"net/http"

	restful "github.com/emicklei/go-restful"
	"github.com/ingmardrewing/gomicRest/db"
	"github.com/ingmardrewing/gomicRest/gomicRest"
)

/**
 * struct for the comic page
 */

type Page struct {
	Id, PageNumber                     int
	Title, Path, ImgUrl, DisqusId, Act string
}

func main() {
	db.Initialize()

	restful.Add(gomicRest.NewPagesService())
	restful.Add(gomicRest.NewSocMedService())

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err.Error())
	}
}
