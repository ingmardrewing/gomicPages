package gomicRest

import (
	"database/sql"
	"fmt"
	"log"

	restful "github.com/emicklei/go-restful"
	"github.com/ingmardrewing/gomicRest/content"
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
	pages := getDbData(rows)
	return pages[0]
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
	log.Printf("Posting page")
	p := new(content.Page)
	request.ReadEntity(p)
	response.WriteEntity(p)
}

func DeletePage(request *restful.Request, response *restful.Response) {
	log.Printf("Delete page")
	msg := Msg{"page deleted"}
	response.WriteEntity(msg)
}

/**
 * struct for a no-read request
 */

type Msg struct {
	msg string
}
