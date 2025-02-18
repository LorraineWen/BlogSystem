package dao

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/url"
	"time"
)

var DB *sql.DB

func init() {
	dsn := fmt.Sprintf("root:kylin.2023@tcp(localhost:3306)/goblog?charset=utf8&loc=%s&parseTime=true", url.QueryEscape("Asia/Shanghai"))
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}
	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxLifetime(time.Minute * 5)
	db.SetConnMaxIdleTime(time.Minute * 5)
	err = db.Ping()
	if err != nil {
		db.Close()
		log.Fatal(err)
	}
	DB = db
}
