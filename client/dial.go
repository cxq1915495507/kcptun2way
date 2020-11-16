package main

import (
	"io"
	//"github.com/pkg/errors"
	//kcp "github.com/xtaci/kcp-go/v5"
	//"github.com/xtaci/tcpraw"
	"kcp-go"
	"fmt"
	"net"
	"os"

	//"strings"

	//"os"
	//"io"


)

func dial(config *Config, block kcp.BlockCrypt) (*kcp.UDPSession, error) {
	con,err:= kcp.DialWithOptions(config.RemoteAddr, block, config.DataShard, config.ParityShard)
	if err != nil {
		panic(err)
	}
	return con,nil
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

	buf := make([]byte, 4096)
	n, err := conn.Read(buf)
	if err != nil {
		fmt.Printf("conn.Read() err:%v\n", err)
		return
	}
	fileName := string(buf[:n])

	//回写ok给发送端
	_, _ = conn.Write([]byte("ok"))


	recivefil(conn,fileName)

}
func recivefil(conn net.Conn,fileName string) {

	file, err := os.Create(fileName)
	if err != nil {
		fmt.Printf("os.Create() err:%v\n", err)
		return
	}
	defer file.Close()

	//从网络中读数据，写入本地文件
	for {
		buf := make([]byte, 4096)
		n, err := conn.Read(buf)

		//写入本地文件，读多少，写多少
		file.Write(buf[:n])
		if err != nil {
			if err == io.EOF {
				fmt.Printf("recieving finish\n")
				fi,err:=os.Stat(fileName)
				if err ==nil {
					fmt.Println("file size is ",fi.Size(),"Bytes")
				}
			} else {
				fmt.Printf("conn.Read() err:%v\n", err)
			}
			return
		}

	}

}
