package main

import (
	"log"
	"net/http"

	restful "github.com/emicklei/go-restful"
	"github.com/ingmardrewing/gomicPages/config"
	"github.com/ingmardrewing/gomicPages/db"
	"github.com/ingmardrewing/gomicPages/service"
)

func main() {
	db.Initialize()

	restful.Add(service.NewPagesService())

	crt, key := config.GetTlsPaths()
	log.Println("Reading crt and key data from files:")
	log.Println(crt)
	log.Println(key)
	err := http.ListenAndServeTLS(":8443", crt, key, nil)
	if err != nil {
		log.Fatal(err.Error())
	}
}
