package db

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"

	"github.com/ingmardrewing/gomicRest/config"
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
		panic(err.Error())
	}
	return rows
}

func InsertPage() {
	/*
		ins := fmt.Sprintf("INSERT INTO pages (title, path, imgUrl, disqusId, act) VALUES('%s', '%s', '%s', '%s', '%s');\n", p.Title(), p.FSPath(), p.ImgUrl(), p.DisqusId(), "Act III")
		_, err := db.Exec(ins)
		if err != nil {
			panic(err.Error())
		}
	*/
}
