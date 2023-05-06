package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	host := "root:root@tcp(localhost:3306)"
	dbName := "ai_chat_3"
	db, err := sql.Open("mysql", host+"/"+dbName)
	if err != nil {
		panic(err)
	}
	dSql := "alter database " + dbName + " charset=utf8mb4 collate=utf8mb4_unicode_520_ci"
	_, err = db.Exec(dSql)
	if err != nil {
		panic(err)
	}
	tSql := "show tables"
	tRows, err := db.Query(tSql)
	if err != nil {
		panic(err)
	}
	for tRows.Next() {
		var tName string
		err = tRows.Scan(&tName)
		if err != nil {
			panic(err)
		}
		cSql := "alter table " + tName + " charset=utf8mb4 collate=utf8mb4_unicode_520_ci"
		db.Exec(cSql)
		sSql := `SELECT
    COLUMN_NAME,
    COLUMN_DEFAULT,
    IS_NULLABLE,
    COLUMN_TYPE,
    COLUMN_COMMENT
FROM
    information_schema.COLUMNS
WHERE
    TABLE_SCHEMA = "` + dbName + `" AND TABLE_NAME = "` + tName + `" AND (DATA_TYPE = "varchar" OR DATA_TYPE = "text")`
		sRows, err := db.Query(sSql)
		if err != nil {
			panic(err)
		}
		for sRows.Next() {
			var cName string
			var cDefault *string
			var isNullAble string
			var cType string
			var cComment string
			err = sRows.Scan(&cName, &cDefault, &isNullAble, &cType, &cComment)
			if err != nil {
				panic(err)
			}
			aSql := "ALTER TABLE `" + tName + "` MODIFY COLUMN `" + cName + "` " + cType + ` CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_520_ci `
			if isNullAble == "NO" {
				if cDefault != nil && *cDefault != "" {
					aSql += " DEFAULT '" + *cDefault + "'"
				}
			} else {
				aSql += " DEFAULT NULL "
			}
			if cComment != "" {
				aSql += " COMMENT '" + cComment + "'"
			}
			_, err = db.Exec(aSql)
			if err != nil {
				fmt.Println(aSql)
				panic(err)
			}
		}
	}
}
