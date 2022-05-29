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

	if err := config.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			fmt.Println("找不到配置文件..")
		} else {
			fmt.Println("配置文件出错..")
		}
	}
	udpport := config.GetInt("udp.port")
	udpip := config.GetString("udp.host")
	mes := config.GetString("udp.mes")
	udpAddr := net.UDPAddr{IP: net.ParseIP(udpip), Port: udpport}
	udpConn, err := net.DialUDP("udp", nil, &udpAddr)
	if err != nil {
		fmt.Printf("连接服务端失败: %s \n", err)
		return
	}
	defer udpConn.Close()
	sendData := []byte(mes)
	fmt.Printf("发送数据完成")
	_, err = udpConn.Write(sendData)
	if err != nil {
		fmt.Printf("发送数据失败: %s \n", err)
		return
	}
	// readData := [1024]byte{}
	// udp, addr, err := udpConn.ReadFromUDP(readData[:])
	// if err != nil {
	// 	fmt.Printf("接收数据失败: %s \n", err)
	// 	return
	// }
	// fmt.Printf("recv: %v addr: %v count: %v \n", string(readData[:udp]), addr, udp)
}
