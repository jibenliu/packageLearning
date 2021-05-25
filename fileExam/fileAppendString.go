package main

import "os"

func main() {
	f, err := os.OpenFile(`test.txt`, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0660)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	_, _ = f.WriteString("文本追加\n")

}
