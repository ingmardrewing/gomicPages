package service

import (
	"fmt"
	"log"

	"golang.org/x/crypto/bcrypt"

	restful "github.com/emicklei/go-restful"
	"github.com/ingmardrewing/gomicPages/config"
	"github.com/ingmardrewing/gomicPages/content"
	"github.com/ingmardrewing/gomicPages/db"
)

func NewPagesService() *restful.WebService {
	path := "/0.1/gomic/page"
	srv := new(restful.WebService)
	srv.
		Path(path).
		Consumes(restful.MIME_JSON).
		Produces(restful.MIME_JSON)

	log.Printf("Creating pages servicel at localhost:8080 -- access with http://localhost:8080%s\n", path)
	srv.Route(srv.GET("/{page-id}").Filter(basicAuthenticate).To(GetPage))
	srv.Route(srv.GET("/").Filter(basicAuthenticate).To(GetPages))
	srv.Route(srv.PUT("/").Filter(basicAuthenticate).To(PutPage))
	srv.Route(srv.POST("/{page-id}").Filter(basicAuthenticate).To(PostPage))
	srv.Route(srv.DELETE("/{page-id}").Filter(basicAuthenticate).To(DeletePage))
	return srv
}

func basicAuthenticate(req *restful.Request, resp *restful.Response, chain *restful.FilterChain) {
	fmt.Println("authenticating ...")
	user, pass, _ := req.Request.BasicAuth()
	password := []byte(pass)
	stored_hash := []byte(config.GetPasswordHashForUser(user))

	err := bcrypt.CompareHashAndPassword(stored_hash, password)
	if err != nil {
		resp.AddHeader("WWW-Authenticate", "Basic realm=Protected Area")
		resp.WriteErrorString(401, "401: Not Authorized")
		return
	}

	chain.ProcessFilter(req, resp)
}

func PutPage(request *restful.Request, response *restful.Response) {
	p := new(content.Page)
	request.ReadEntity(p)
	db.Insert(p)
	response.WriteEntity(p)
}

func GetPage(request *restful.Request, response *restful.Response) {
	id := request.PathParameter("page-id")
	page := db.GetPage(id)
	response.WriteEntity(page)
}

func GetPages(request *restful.Request, response *restful.Response) {
	pages := db.GetAllPages()
	response.WriteEntity(pages)
}

func PostPage(request *restful.Request, response *restful.Response) {
	p := new(content.Page)
	request.ReadEntity(p)
	id := request.PathParameter("page-id")
	db.Update(p, id)
	response.WriteEntity(p)
}

func DeletePage(request *restful.Request, response *restful.Response) {
	id := request.PathParameter("page-id")
	db.Delete(id)
}
