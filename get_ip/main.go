package main

import (
	"fmt"
	"net"
)

func GetLocalIp() (ip string, err error){
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return
	}

	for _, addr := range addrs{
		ipAddr, ok := addr.(*net.IPNet) // 类型断言
		if !ok{
			continue
		}

		if ipAddr.IP.IsLoopback(){
			continue
		}
		if !ipAddr.IP.IsGlobalUnicast(){
			continue
		}
		fmt.Println(ipAddr)
		return ipAddr.IP.String(), nil
	}

	return 
}

func GetOutboundIP() string{
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil{
		fmt.Println(err)
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)
	fmt.Println(localAddr.String())
	return localAddr.IP.String()
}
func main(){
	// GetLocalIp()
	GetOutboundIP()
}