package tool

import (
	"encoding/json"
	"os"
)

var QSlice []map[string]interface{}

func ReadJsonFile() {
	file := "zhihu.json"
	content, err := os.ReadFile(file)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(content, &QSlice)
	if err != nil {
		panic(err)
	}
}
