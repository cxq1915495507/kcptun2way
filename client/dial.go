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
	"io"
	"net"
	"os"
	"strconv"
	"strings"
)

func dial(config *Config) (*net.UDPConn, uint32,error) {

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



	//2.提取文件的绝对路径  exp ：/home/sxf/hh
	var filePath string
	filePath = `/mnt/d/tl_98dx_torisetu_300.pdf`
	//3.提取不包含路径的文件名 exp:hh
	fileInfo ,err := os.Stat(`/mnt/d/tl_98dx_torisetu_300.pdf`)//fileInfo中包含文件名和文件大小
	if err != nil{
		fmt.Println("os.Stat err",err)
		return
	}
	fileName :=fileInfo.Name() //文件名hh
	fileSize :=fileInfo.Size() //文件大小
	fmt.Println("文件名：",fileName)
	fmt.Println("文件大小",fileSize)

	//4.主动发起连接请求
	conn ,err := net.Dial("tcp","127.0.0.1:8088")
	if err != nil{
		fmt.Println("net.Dial err",err)
		return
	}
	defer conn.Close()

	//5.发送文件名给接收端
	conn.Write([]byte(fileName))

	//6.读取接收端回发的响应数据（ok）
	buf := make([]byte,4096)
	n, err := conn.Read(buf)
	if err != nil{
		fmt.Println("conn.Read err",err)
		return
	}

	//7.判断这个数据是否是ok
	if "ok" == string(buf[:n]){
		//8.是ok，写文件内容给接收端--借助conn
		sendcontent(conn,filePath)
	}


}

//发送文件内容给接受端
func sendcontent(conn net.Conn,filePath string) {

	//8.1只读打开文件
	f, err := os.Open(filePath)
	if err != nil {
		fmt.Println("os.Open err", err)
		return
	}
	defer f.Close()

	//8.2从本地文件读数据，写给接收端
	buf := make([]byte, 4096)
	for {
		n, err := f.Read(buf) //读取本地文件放入缓冲区buf中
		if err != nil {
			if err == io.EOF {
				fmt.Println("读取文件完成")
			} else {
				fmt.Println("f.Read err", err)
			}
			return
		}
		_, err = conn.Write(buf[:n]) //将缓冲区buf中数据写到网络socket中
		if err != nil {
			fmt.Println("conn.Write err", err)
			return
		}
	}

}