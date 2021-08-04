package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func init() {
	RunExe("tmain")
}

func main() {

}

// RunExe 启动exe
func RunExe(exeName string) {
	fmt.Println("启动 exe：", exeName)
	args := []string{"golang", "语言社区"}
	_ = args
	strPath := getCurrentPath()
	fmt.Println("当前路径：", strPath)
	strPath = strPath + exeName
	//exec.Command(strPath, args)
}

var PathData string = ""

func getCurrentPath() string {
	if len(PathData) == 0 {
		s, _ := exec.LookPath(os.Args[0])
		i := strings.LastIndex(s, "//")
		path := s[0 : i+1]
		PathData = path
		return path
	}
	return PathData
}
