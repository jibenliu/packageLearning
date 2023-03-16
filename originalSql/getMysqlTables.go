package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	host := "root:root@tcp(localhost:3306)"
	db, err := sql.Open("mysql", host+"/auth_app_4")
	if err != nil {
		panic(err)
	}
	dSql := "show databases"
	tSql := "show tables"
	rows, err := db.Query(dSql)
	if err != nil {
		panic(err)
	}
	for rows.Next() {
		var dbName string
		err = rows.Scan(&dbName)
		dbConn, err := sql.Open("mysql", host+"/"+dbName)
		if err != nil {
			panic(err)
		}
		tRows, err := dbConn.Query(tSql)
		if err != nil {
			panic(err)
		}
		for tRows.Next() {
			var tName string
			err = tRows.Scan(&tName)
			if err != nil {
				panic(err)
			}
			if tName == "metric_one" {
				fmt.Println(dbName)
				return
			}
		}
	}
}
