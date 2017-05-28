package service

import (
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

	srv.Route(srv.GET("/{page-id}").Filter(basicAuthenticate).To(GetPage))
	srv.Route(srv.GET("/").Filter(basicAuthenticate).To(GetPages))
	srv.Route(srv.PUT("/").Filter(basicAuthenticate).To(PutPage))
	srv.Route(srv.POST("/{page-id}").Filter(basicAuthenticate).To(PostPage))
	srv.Route(srv.DELETE("/{page-id}").Filter(basicAuthenticate).To(DeletePage))
	return srv
}

func basicAuthenticate(request *restful.Request, response *restful.Response, chain *restful.FilterChain) {
	err := authenticate(request)
	log.Println(err)
	if err != nil {
		response.AddHeader("WWW-Authenticate", "Basic realm=Protected Area")
		response.WriteErrorString(401, "401: Not Authorized")
		return
	}

	chain.ProcessFilter(request, response)
}

func authenticate(req *restful.Request) error {
	user, pass, _ := req.Request.BasicAuth()
	given_pass := []byte(pass)
	stored_hash := []byte(config.GetPasswordHashForUser(user))
	//hash, _ := bcrypt.GenerateFromPassword(given_pass, coast)
	return bcrypt.CompareHashAndPassword(stored_hash, given_pass)
}

func PutPage(request *restful.Request, response *restful.Response) {
	p := new(content.Page)
	request.ReadEntity(p)
	nr := db.GetHighestPageNumber()
	p.PageNumber = nr + 1
	db.Insert(p)
	response.WriteEntity(p)
}

func GetPage(request *restful.Request, response *restful.Response) {
	id := request.PathParameter("page-id")
	page := db.GetPage(id)
	response.WriteEntity(page)
}

func GetPages(request *restful.Request, response *restful.Response) {
	pgs := db.GetAllPages()
	p := &content.Pages{pgs}
	response.WriteEntity(p)
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
