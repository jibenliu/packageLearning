package main

import (
	"bufio"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

const suffixName = ".go"

func main() {
	fmt.Println(pathWalk("/media/jibenliu/HDD/goProjects/packageLearning/connectPool"))
}

func pathWalk(dir string) int{
	line := 0
	filepath.Walk(dir, func(path string, info fs.FileInfo, err error) error {
		if !info.IsDir() {
			if strings.HasSuffix(path, suffixName){
				 file, err := os.Open(path)
				 if err != nil {
				 	return err
				 }
				defer file.Close()

				reader := bufio.NewReader(file)
				for {
					_, isPrefix, err := reader.ReadLine()
					if err != nil {
						break
					}
					if !isPrefix {
						line++
					}
				}
			}
			return nil
		}
		return nil
	})
	return line
}
