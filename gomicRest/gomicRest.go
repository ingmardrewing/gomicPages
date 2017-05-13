package gomicRest

import (
	"database/sql"
	"log"

	restful "github.com/emicklei/go-restful"
	"github.com/ingmardrewing/gomicRest/db"
)

func New() *restful.WebService {
	pagePath := "/gomic/page"
	service := new(restful.WebService)
	service.
		Path(pagePath).
		Consumes(restful.MIME_JSON).
		Produces(restful.MIME_JSON)

	log.Printf("Starting server at localhost:8080 -- access with http://localhost:8080%s\n", pagePath)
	service.Route(service.GET("/{page-id}").To(GetPage))
	service.Route(service.GET("/").To(GetPages))
	service.Route(service.PUT("/").To(PutPage))
	service.Route(service.POST("/").To(PostPage))
	service.Route(service.DELETE("/").To(DeletePage))
	return service
}

func PutPage(request *restful.Request, response *restful.Response) {
	log.Print("Put page")
	msg := Msg{"page put"}
	response.WriteEntity(msg)
}

func GetPage(request *restful.Request, response *restful.Response) {
	pageId := request.PathParameter("page-id")
	log.Printf("Get page: %s", pageId)

	page := Page{0, 0, "test title", "/test/path", "test imgurl", "testdisqusid", "testact"}
	response.WriteEntity(page)
}

func GetPages(request *restful.Request, response *restful.Response) {
	log.Println("Get pages")

	pages := getAllPages()
	response.WriteEntity(pages)
}

func getAllPages() []Page {
	pages := []Page{}
	log.Println("about to query")
	rows := db.Query("SELECT * FROM gomic.pages")
	log.Println("actually queried ..")
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

			pages = append(pages, Page{
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
	log.Printf("Posting page")
	p := new(Page)
	request.ReadEntity(p)
	response.WriteEntity(p)
}

func DeletePage(request *restful.Request, response *restful.Response) {
	log.Printf("Delete page")
	msg := Msg{"page deleted"}
	response.WriteEntity(msg)
}

/**
 * struct for the comic page
 */

type Page struct {
	Id, PageNumber                     int
	Title, Path, ImgUrl, DisqusId, Act string
}

/**
 * struct for a no-read request
 */

type Msg struct {
	msg string
}
