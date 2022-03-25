package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/url"

	_ "github.com/denisenkom/go-mssqldb"
)

func main() {
	u := &url.URL{
		Scheme: "sqlserver",
		User:   url.UserPassword("tryout", "tryout"),
		Host:   fmt.Sprintf("%s:%d", "10.77.3.211", 1433),
	}
	log.Println(u.String())
	db, err := sql.Open("sqlserver", u.String())
	if err != nil {
		log.Fatal(err)
	}
	res, err := db.Exec("insert into dbo.Table1 (ID, NAME) values (@p1, @p2)", "1", "NAME1")
	if err != nil {
		log.Fatal(err)
	}
	log.Println(res)
}
