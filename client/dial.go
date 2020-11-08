package main

import (
	//"github.com/pkg/errors"
	//kcp "github.com/xtaci/kcp-go/v5"
	//"github.com/xtaci/tcpraw"
	"kcp-go"
	"fmt"
	"net"
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
