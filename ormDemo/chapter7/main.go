package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"newTest/ormDemo/chapter7/geeorm"
)

func main() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s", "root", "123456", "127.0.0.1", 3306, "test", "utf8")
	engine, _ := geeorm.NewEngine("mysql", dsn)
	defer engine.Close()
	s := engine.NewSession()
	_, _ = s.Raw("DROP TABLE IF EXISTS user;").Exec()
	sql := `
CREATE TABLE user (
    uid INT(10) NOT NULL AUTO_INCREMENT,
    name VARCHAR(64)  DEFAULT NULL,
    PRIMARY KEY (uid)
)ENGINE=InnoDB DEFAULT CHARSET=utf8;
`
	_, _ = s.Raw(sql).Exec()
	_, _ = s.Raw(sql).Exec()
	result, _ := s.Raw("INSERT INTO user(`name`) values (?), (?)", "Tom", "Sam").Exec()
	count, _ := result.RowsAffected()
	fmt.Printf("Exec success, %d affected\n", count)
}
