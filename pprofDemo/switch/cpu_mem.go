package main

import (
	"fmt"
	"github.com/ricochet2200/go-disk-usage/du"
	"github.com/shirou/gopsutil/mem"
	"github.com/shirou/gopsutil/v3/cpu"
	"log"
	"time"
)

var KB = uint64(1024)

type ComputerMonitor struct {
	CPU      float64 `json:"cpu"`
	Mem      float64 `json:"mem"`
	DiskFree float32 `json:"diskFree"`
}

// GetCPUPercent 获取CPU使用率
func GetCPUPercent() float64 {
	percent, err := cpu.Percent(time.Second, false)
	if err != nil {
		log.Fatalln(err.Error())
		return -1
	}
	return percent[0]
}

// GetMemPercent 获取内存使用率
func GetMemPercent() float64 {
	memInfo, err := mem.VirtualMemory()
	if err != nil {
		log.Fatalln(err.Error())
		return -1
	}
	return memInfo.UsedPercent
}

func GetCpuMem() ComputerMonitor {
	var res ComputerMonitor
	res.CPU = GetCPUPercent()
	res.Mem = GetMemPercent()

	usage := du.NewDiskUsage("/")
	fmt.Println("Free:", usage.Free()/(KB*KB))
	fmt.Println("Available:", usage.Available()/(KB*KB))
	fmt.Println("Size:", usage.Size()/(KB*KB))
	fmt.Println("Used:", usage.Used()/(KB*KB))
	fmt.Println("Usage:", usage.Usage()*100, "%")
	res.DiskFree = usage.Usage() * 100
	fmt.Printf("cpu使用率：%.2f%%\n", res.CPU)
	fmt.Printf("内存使用率：%.2f%%\n", res.Mem)
	fmt.Printf("磁盘空闲率：%.2f%%\n", res.DiskFree)
	return res
}
