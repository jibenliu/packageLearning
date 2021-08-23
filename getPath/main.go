package main

import (
	"flag"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

// GetAppPath
// when the project is running on build way to get path
func GetAppPath() string {
	file, _ := exec.LookPath(os.Args[0])
	path, _ := filepath.Abs(file)
	index := strings.LastIndex(path, string(os.PathSeparator))

	return path[:index]
}

// GetRunPath
// demo: go run main.go --app-path "Your project address"
func GetRunPath() string {
	var appPath string
	flag.StringVar(&appPath, "app-path", "app-path", "runtime way get path")
	flag.Parse()
	return appPath
}
