package main

import (
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize/v2"
	"os"
	"strconv"
	"time"
)

var axisMap = map[string]string{
	"A": "customer_nickname",
	"B": "platform",
	"C": "customer_name",
	"D": "industry",
	"E": "second_industry",
	"F": "account_strategy",
	"G": "make_costs",
	"H": "seller",
	"I": "operate_type",
	"J": "remark_rate",
	"K": "mechanism",
	"L": "upper_name",
	"M": "platform_strategy",
	"N": "medium",
}

var cuKey = []string{
	"customer_name",
	"industry",
	"second_industry",
	"platform",
	"account_strategy",
	"customer_nickname",
	"seller",
	"seller_id",
	"mechanism",
	"make_costs",
	"operate_type",
	"remark_rate",
	"created_at",
	"updated_at",
}
var cuInfoKey = []string{
	"customer_id",
	"customer_nickname",
	"platform",
	"upper_name",
	"platform_strategy",
	"medium_id",
	"medium",
	"seller",
	"seller_id",
	"created_at",
	"updated_at",
}
var sellerId = 51
var mediumId = 29

func main() {
	f, err := excelize.OpenFile("doc.xlsx")
	if err != nil {
		fmt.Println(err)
		return
	}
	rows, _ := f.GetRows("Sheet1")
	columns, _ := f.GetCols("Sheet1")
	okSleet := make([]map[string]string, len(rows)-1)
	for index, _ := range rows {
		if index == 0 {
			continue
		}
		temp := make(map[string]string, len(columns))
		for key, item := range axisMap {
			axis := key + strconv.Itoa(index+1)
			temp[item], _ = f.GetCellValue("Sheet1", axis)
		}
		okSleet[index-1] = temp
	}
	line := len(okSleet)
	sqlFile, err := os.Create("insert.sql")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer sqlFile.Close()
	for l, record := range okSleet {
		if l != line-1 { //最后一行的媒介id
			mediumId = 41
		}
		demoSql := `
DELIMITER $$
DROP PROCEDURE IF EXISTS insertData $$
CREATE PROCEDURE insertData()
BEGIN
    DECLARE customerId INT(10) DEFAULT 0;
	START TRANSACTION;
	%s
	SET customerId = (SELECT LAST_INSERT_ID());
	SET @insertSql = CONCAT(REPLACE(%s,'customerId', customerId));
	PREPARE itemSql FROM @insertSql;
	EXECUTE itemSql;
	COMMIT;
END $$
DELIMITER ;

CALL insertData();
DROP PROCEDURE insertData;
`
		cuInsertSql := makeCuSql(record)
		cuInfoInsertSql := makeCuInfoSql(record)
		fmt.Fprintf(sqlFile, demoSql, cuInsertSql, "\""+cuInfoInsertSql+"\"")
	}
}

func makeCuSql(record map[string]string) string {
	cuInsertSql := "INSERT INTO customers "
	keySql := " ("
	valueSql := " ("
	keyLen := len(cuKey)
	for index, value := range cuKey {
		if index != keyLen-1 {
			keySql += "`" + value + "`, "
		} else {
			keySql += "`" + value + "`)"
		}
		temp := ""
		v, ok := record[value]
		if ok {
			if value == "make_costs" || value == "account_strategy" || value == "remark_rate" {
				temp = v
			} else {
				temp = "'" + v + "'"
			}
		} else if value == "seller_id" {
			temp = strconv.Itoa(sellerId)
		} else if value == "created_at" || value == "updated_at" {
			temp = "'" + time.Now().Format("2006-01-02 15:04:05") + "'"
		}
		if index != keyLen-1 {
			valueSql += temp + ", "
		} else {
			valueSql += temp + ");"
		}
	}
	cuInsertSql += keySql
	cuInsertSql += " VALUES "
	cuInsertSql += valueSql
	return cuInsertSql
}

func makeCuInfoSql(record map[string]string) string {
	cuInfoInsertSql := "INSERT INTO customer_infos "
	keySql := " ("
	valueSql := " ("
	keyLen := len(cuInfoKey)
	for index, value := range cuInfoKey {
		if index != keyLen-1 {
			keySql += "`" + value + "`, "
		} else {
			keySql += "`" + value + "`)"
		}
		temp := ""
		v, ok := record[value]
		if ok {
			if value == "platform_strategy" {
				temp = v
			} else {
				temp = "'" + v + "'"
			}
		} else if value == "medium_id" {
			temp = strconv.Itoa(mediumId)
		} else if value == "seller_id" {
			temp = strconv.Itoa(sellerId)
		} else if value == "created_at" || value == "updated_at" {
			temp = "'" + time.Now().Format("2006-01-02 15:04:05") + "'"
		} else if value == "customer_id" {
			temp = "customerId"
		}
		if index != keyLen-1 {
			valueSql += temp + ", "
		} else {
			valueSql += temp + ")"
		}
	}
	cuInfoInsertSql += keySql
	cuInfoInsertSql += " VALUES "
	cuInfoInsertSql += valueSql
	return cuInfoInsertSql
}
