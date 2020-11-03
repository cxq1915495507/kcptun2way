package main

import (
	//"crypto/rand"
	//	"encoding/binary"
	//"crypto/rand"
	//"encoding/binary"
	//"crypto/rand"
	//"encoding/binary"
	//"github.com/pkg/errors"
	//kcp "github.com/xtaci/kcp-go/v5"
	//"github.com/xtaci/tcpraw"
	//"kcp-go"
	"fmt"
	//"io"
	"net"
	//"os"
	"strconv"
	"strings"
)

func dial(config *Config) (*net.UDPConn, uint32,error) {
        fmt.Printf("client dial")
	udpaddr, err := net.ResolveUDPAddr("udp", config.RemoteAddr)
	if err != nil {
		return nil,0, err
	}
	network := "udp4"
	if udpaddr.IP.To4() == nil {
		network = "udp"
	}

	conn, err := net.ListenUDP(network, nil)
	if err != nil {
		return nil, 0,err
	}
	fmt.Printf("DialWithOptions")




	_, err = conn.WriteToUDP([]byte("hello"),udpaddr)
	fmt.Printf("dial:conn.Write")

	message := make([]byte, 20)
	rlen, remote, err := conn.ReadFromUDP(message[:])
	if err != nil {
		panic(err)
	}

	data := strings.TrimSpace(string(message[:rlen]))
	fmt.Printf("received: %s from %s\n", data, remote)
	idd, _ :=strconv.ParseUint(data[:5], 2, 32)
	id:=uint32(idd)

	return conn,id, nil

}



func dial2()  {


	//4.主动发起连接请求
	conn ,err := net.Dial("tcp","localhost:12900")
	if err != nil{
		fmt.Println("net.Dial err",err)
		return
	}
	defer conn.Close()

	//5.发送文件名给接收端
	conn.Write([]byte("i love you"))
	fmt.Println("client send: i love you")

	//6.读取接收端回发的响应数据（ok）
	buf := make([]byte,4096)
	n, err := conn.Read(buf)
	if err != nil{
		fmt.Println("conn.Read err",err)
		return
	}
	fmt.Println(string(buf[:n]))

	//7.判断这个数据是否是ok
	if "ok" == string(buf[:n]){
		//8.是ok，写文件内容给接收端--借助conn
		conn.Write([]byte("really"))
		fmt.Println("client send: really")

	}


}

