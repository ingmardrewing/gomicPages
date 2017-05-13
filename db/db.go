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
	if nil != err {
		panic(err)
	}
	err = d.Ping()
	if nil != err {
		panic(err)
	}
	return d
}

func Query(query string) *sql.Rows {
	rows, err := db.Query(query)
	if err != nil {
		log.Println("Query error")
		panic(err.Error())
	}
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

func handleErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
