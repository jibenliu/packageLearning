package slog

import (
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
)

type Logger struct {
	console bool
	warn    bool
	info    bool
	tFormat func() string
	file    chan string
}

func NewLog(level string, console bool, File *os.File, buf int) (*Logger, error) {
	log := &Logger{console: console, tFormat: format}
	if File != nil {
		FileInfo, err := File.Stat()
		if err != nil {
			return nil, err
		}
		mode := strings.Split(FileInfo.Mode().String(), "-")
		if strings.Contains(mode[1], "w") {
			strChan := make(chan string, buf)
			log.file = strChan
			go func() {
				for {
					fmt.Fprintln(File, <-strChan)
				}
			}()
			defer func() {
				for len(strChan) > 0 {
					time.Sleep(1e9)
				}
			}()
		} else {
			return nil, errors.New("can't write.")
		}
	}
	switch level {
	case "Warn":
		log.warn = true
		return log, nil
	case "Info":
		log.warn = true
		log.info = true
		return log, nil
	}
	return nil, errors.New("level must be Warn or Info.")
}

func (l *Logger) Error(info interface{}) {
	if l.console {
		fmt.Println("Error", l.tFormat(), info)
	}
	if l.file != nil {
		l.file <- fmt.Sprintf("Error %s", info)

	}
}

func (l *Logger) Warn(info interface{}) {
	if l.warn && l.console {
		fmt.Println("Warn", info)
	}
	if l.file != nil {
		l.file <- fmt.Sprintf("Warn %s", info)
	}
}
func (l *Logger) Info(info interface{}) {
	if l.info && l.console {
		fmt.Println("Info", info)
	}
	if l.file != nil {
		l.file <- fmt.Sprintf("Info %s", info)
	}
}
func (l *Logger) Close() {
	for len(l.file) > 0 {
		time.Sleep(1e8)
	}
}
func format() string {
	return time.Now().Format("2006-01-02 15:04:05")
}
