package main

import (
	"fmt"
	"net"

	"github.com/spf13/viper"
)

//
func main() {
	config := viper.New()
	config.AddConfigPath(".")   // 文件所在目录
	config.SetConfigName("udp") // 文件名
	config.SetConfigType("ini") // 文件类型
	var datacount int
	datacount = 1

	if err := config.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			fmt.Println("找不到配置文件..")
		} else {
			fmt.Println("配置文件出错..")
		}
	}
	udpport := config.GetInt("udp.port")
	// upd地址
	updAddr := net.UDPAddr{
		IP:   net.IPv4(0, 0, 0, 0),
		Port: udpport,
	}
	fmt.Printf("UDP server running...  IP: %v Port: %v  \n", updAddr.IP, updAddr.Port)

	listen, err := net.ListenUDP("udp", &updAddr)
	if err != nil {
		fmt.Println("listen failed err: ", err)
		return
	}
	// 关闭连接
	defer func(listen *net.UDPConn) {
		err := listen.Close()
		if err != nil {
			fmt.Println("listen failed err: ", err)
		}
	}(listen)
	for {
		// 接收数据
		var data [1024]byte

		udp, addr, err := listen.ReadFromUDP(data[:])
		if err != nil {
			fmt.Println("listen failed err: ", err)
			continue
		}

		fmt.Printf("data: %v add: %v count: %v datacount: %d     \n\n ", string(data[:udp]), addr, udp, datacount)
		// 发送数据
		datacount++
		_, err = listen.WriteToUDP(data[:udp], addr)
		if err != nil {
			fmt.Println("listen failed err: ", err)
			continue
		}
	}
}
