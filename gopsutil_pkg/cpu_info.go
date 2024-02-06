package main

import (
	"fmt"
	"time"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/load"
	"github.com/shirou/gopsutil/mem"
	"github.com/shirou/gopsutil/net"
)
func getCpuInfo() {
	cpuInfos, err := cpu.Info()
	if err != nil {
		fmt.Printf("get cpu info failed, err: %v \n", err)
	}
	for _, ci := range cpuInfos{
		fmt.Println(ci)
	}
	// CPU使用率
	for{
		percent, _ := cpu.Percent(time.Second, false)
		fmt.Printf("cpu percent: %v \n", percent)
	}
}

// cpu负载
func getLoad(){
	info, err := load.Avg()
	if err != nil{
		fmt.Printf("load.Avg() failed, err: %v", err)
		return
	}
	fmt.Println(info)
}
// 内存信息
func getMemInfo(){
	info, err := mem.VirtualMemory()
	if err != nil{
		fmt.Printf("mem.VirtualMemory() failed, err: %v", err)
		return
	}
	fmt.Println(info)
}
// 主机信息
func getHostInfo(){
	info, err := host.Info()
	if err != nil{
		fmt.Printf("host.Info() failed, err: %v", err)
		return
	}
	fmt.Println(info)
}

// 磁盘信息
func getDiskInfo(){
	parts, err := disk.Partitions(true)
	if err != nil{
		fmt.Printf("get disk partitions failed, err: %v", err)
		return
	}
	for _, part := range parts{
		partInfo, err := disk.Usage(part.Mountpoint)
		if err != nil{
			fmt.Printf("get part Info failed, err: %v", err)
			return
		}
		fmt.Println(partInfo)
	}
	// IO
	ioStar, _ := disk.IOCounters()
	for k, v := range ioStar{
		fmt.Printf("%v : %v \n", k, v)
	}
	// fmt.Println(parts)
}
// net io
func getNetInfo(){
	netIOs, err := net.IOCounters(true)
	if err != nil{
		fmt.Printf("get netIO Info failed, err: %v", err)
		return
	}
	for _, netIO := range netIOs{
		fmt.Println(netIO)
	}
}
func main() {
	// getCpuInfo()
	getLoad()
	getMemInfo()
	getHostInfo()
	getDiskInfo()
	getNetInfo()
}