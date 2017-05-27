package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"

	"github.com/ingmardrewing/gomicPages/config"
	"github.com/ingmardrewing/gomicPages/content"
)

var db *sql.DB

func Initialize() {
	dsn := config.GetDsn()
	db = getDb(dsn)
}

func getDb(dsn string) *sql.DB {
	d, err := sql.Open("mysql", dsn)
	handleErr(err)
	err = d.Ping()
	handleErr(err)
	return d
}

func Query(query string) *sql.Rows {
	rows, err := db.Query(query)
	handleErr(err)
	return rows
}

func Insert(p *content.Page) {
	stmt, err := db.Prepare("INSERT INTO pages(title, path, imgUrl, disqusId, act, pageNumber) VALUES(?, ?, ?, ?, ?, ?)")
	handleErr(err)
	res, err := stmt.Exec(p.Title, p.Path, p.ImgUrl, p.DisqusId, p.Act, p.PageNumber)
	handleErr(err)
	lastId, err := res.LastInsertId()
	handleErr(err)
	rowCnt, err := res.RowsAffected()
	handleErr(err)
	log.Printf("ID = %d, affected = %d\n", lastId, rowCnt)
}

func Update(p *content.Page, id string) {
	stmt, err := db.Prepare("UPDATE pages SET title=?, path=?, imgUrl=?, disqusId=?, act=?, pageNumber=? WHERE id=?")
	handleErr(err)
	res, err := stmt.Exec(p.Title, p.Path, p.ImgUrl, p.DisqusId, p.Act, p.PageNumber, id)
	handleErr(err)
	lastId, err := res.LastInsertId()
	handleErr(err)
	rowCnt, err := res.RowsAffected()
	handleErr(err)
	log.Printf("ID = %d, affected = %d\n", lastId, rowCnt)
}

func Delete(id string) {
	stmt, err := db.Prepare("DELETE FROM pages WHERE id=?")
	handleErr(err)
	res, err := stmt.Exec(id)
	handleErr(err)
	lastId, err := res.LastInsertId()
	handleErr(err)
	rowCnt, err := res.RowsAffected()
	handleErr(err)
	log.Printf("ID = %d, affected = %d\n", lastId, rowCnt)
}

func handleErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func GetPage(id string) content.Page {
	rows := Query(fmt.Sprintf("SELECT * FROM gomic.pages where id = %s", id))
	if rows != nil {
		pages := getDbData(rows)
		if len(pages) > 0 {
			return pages[0]
		}
	}
	return content.EmptyPage()
}

func GetAllPages() []content.Page {
	rows := Query("SELECT * FROM gomic.pages")
	return getDbData(rows)
}

func GetLatestPage() content.Page {
	pages := GetAllPages()
	return pages[len(pages)-1]
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
