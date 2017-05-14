package gomicRest

import (
	"database/sql"
	"fmt"
	"log"

	"golang.org/x/crypto/bcrypt"

	restful "github.com/emicklei/go-restful"
	"github.com/ingmardrewing/gomicRest/config"
	"github.com/ingmardrewing/gomicRest/content"
	"github.com/ingmardrewing/gomicRest/db"
)

func NewPagesService() *restful.WebService {
	path := "/0.1/gomic/page"
	service := new(restful.WebService)
	service.
		Path(path).
		Consumes(restful.MIME_JSON).
		Produces(restful.MIME_JSON)

	log.Printf("Creating pages servicel at localhost:8080 -- access with http://localhost:8080%s\n", path)
	service.Route(service.GET("/{page-id}").Filter(basicAuthenticate).To(GetPage))
	service.Route(service.GET("/").Filter(basicAuthenticate).To(GetPages))
	service.Route(service.PUT("/").Filter(basicAuthenticate).To(PutPage))
	service.Route(service.POST("/{page-id}").Filter(basicAuthenticate).To(PostPage))
	service.Route(service.DELETE("/{page-id}").Filter(basicAuthenticate).To(DeletePage))
	return service
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
	page := getPage(id)
	response.WriteEntity(page)
}

func getPage(id string) content.Page {
	rows := db.Query(fmt.Sprintf("SELECT * FROM gomic.pages where id = %s", id))
	if rows != nil {
		pages := getDbData(rows)
		if len(pages) > 0 {
			return pages[0]
		}
	}
	return content.EmptyPage()
}

func GetPages(request *restful.Request, response *restful.Response) {
	pages := getAllPages()
	response.WriteEntity(pages)
}

func getAllPages() []content.Page {
	rows := db.Query("SELECT * FROM gomic.pages")
	return getDbData(rows)
}

func getDbData(rows *sql.Rows) []content.Page {
	pages := []content.Page{}
	if rows != nil {
		defer rows.Close()
		for rows.Next() {
			var (
				id         int
				title      sql.NullString
				path       sql.NullString
				imgUrl     sql.NullString
				disqusId   sql.NullString
				act        sql.NullString
				pageNumber int
			)

			rows.Scan(
				&id,
				&title,
				&path,
				&imgUrl,
				&disqusId,
				&act,
				&pageNumber)

			pages = append(pages, content.Page{
				Id:         id,
				Title:      title.String,
				Path:       path.String,
				ImgUrl:     imgUrl.String,
				DisqusId:   disqusId.String,
				Act:        act.String,
				PageNumber: pageNumber})
		}
	}
	return pages
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
