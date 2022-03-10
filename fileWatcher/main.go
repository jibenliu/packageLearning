package main

import (
	"fmt"
	"syscall"

	"github.com/fsnotify/fsnotify"
	"runtime"
)

var exit chan bool

func main() {
	//1、初始化监控对象watcher
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		fmt.Printf("Fail to create new Watcher[ %s ]\n", err)
	}

	//3、启动监听文件对象事件协程
	go func() {
		fmt.Println("开始监听文件变化")
		for {
			select {
			case e := <-watcher.Events:
				fmt.Println(e.Op)
				fmt.Printf("监听到文件 - %s变化\n", e.Name)
				// 以下行为只能监控linux，如果需要监控windows需要调整信号
				// 可参考github.com/howeyc/fsnotify
				if e.Op&syscall.IN_OPEN == syscall.IN_OPEN {
					fmt.Println("监听到文件打开事件")
				}
				// 这里添加根据文件变化的业务逻辑

				if (e.Op&syscall.IN_CREATE) == syscall.IN_CREATE || (e.Op&syscall.IN_MOVED_TO) == syscall.IN_MOVED_TO {
					fmt.Println("监听到文件创建事件")
				}
				if (e.Op&syscall.IN_DELETE_SELF) == syscall.IN_DELETE_SELF || (e.Op&syscall.IN_DELETE) == syscall.IN_DELETE {
					fmt.Println("监听到文件删除事件")
				}
				if (e.Op&syscall.IN_MODIFY) == syscall.IN_MODIFY || (e.Op&syscall.IN_ATTRIB) == syscall.IN_ATTRIB {
					fmt.Println("监听到文件修改事件")
				}
				if (e.Op&syscall.IN_MOVE_SELF) == syscall.IN_MOVE_SELF || (e.Op&syscall.IN_MOVED_FROM) == syscall.IN_MOVED_FROM {
					fmt.Println("监听到文件重命名事件")
				}
				if (e.Op & syscall.IN_ATTRIB) == syscall.IN_ATTRIB {
					fmt.Println("监听到文件属性修改事件")
				}
				fmt.Println("根据文件变化开始执行业务逻辑")
			case err := <-watcher.Errors:
				fmt.Printf(" %s\n", err.Error())
			}
		}
	}()
	// 2、将需要监听的文件加入到watcher的监听队列中
	paths := []string{"config.yml"}
	for _, path := range paths {
		err = watcher.Add(path) //将文件加入监听
		if err != nil {
			fmt.Sprintf("Fail to watch directory[ %s ]\n", err)
		}
	}

	<-exit
	runtime.Goexit()
}
