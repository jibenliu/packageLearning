package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"time"
)

var engine *xorm.Engine

type LocalTime time.Time

const localDateTimeFormat string = "2006-01-02 15:04:05"

func (l LocalTime) MarshalJSON() ([]byte, error) {
	b := make([]byte, 0, len(localDateTimeFormat)+2)
	b = append(b, '"')
	b = time.Time(l).AppendFormat(b, localDateTimeFormat)
	b = append(b, '"')
	return b, nil
}

func (l *LocalTime) UnmarshalJSON(b []byte) error {
	now, err := time.ParseInLocation(`"`+localDateTimeFormat+`"`, string(b), time.Local)
	*l = LocalTime(now)
	return err
}

type Admin struct {
	Id      int64
	Name    string
	Salt    string
	Age     int
	Passwd  string    `xorm:"varchar(200)"`
	Created time.Time `xorm:"created"`
	Updated int `xorm:"updated"`
}

func main() {
	var err error
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s", "root", "123456", "127.0.0.1", 3306, "test", "utf8")
	engine, err = xorm.NewEngine("mysql", dsn)
	if err != nil {
		fmt.Print(err)
		return
	}
	err = engine.Sync2(new(Admin))
	if err != nil {
		fmt.Print(err)
		return
	}

}
