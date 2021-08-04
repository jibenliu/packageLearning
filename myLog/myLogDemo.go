package main

import (
	"io/ioutil"
	"myLog/myLog"
	"os"
)

func main() {
	myLog.InitLogger(
		myLog.WithFileName("test.log"),
		myLog.WithLogDir("./runtime/logs"),
		myLog.WithSize(1*1024*1024),
	)
	fi, err := os.Open("../readFileExam/medium.txt")
	if err != nil {
		panic(err)
	}
	defer fi.Close()
	fd, err := ioutil.ReadAll(fi)
	if err != nil {
		panic(err)
	}
	for i := 0; i < 50; i++ {
		myLog.T("%s", string(fd))
		myLog.W("%s", string(fd))
		myLog.E("%s", string(fd))
	}
}
