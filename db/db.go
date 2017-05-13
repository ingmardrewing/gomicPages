package db

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"

	"github.com/ingmardrewing/gomicRest/config"
	"github.com/ingmardrewing/gomicRest/content"
)

var db *sql.DB

func Initialize() {
	dsn := config.GetDsn()
	log.Println(dsn)
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

func handleErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
