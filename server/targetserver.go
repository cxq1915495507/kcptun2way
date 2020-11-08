package main

import (
	"fmt"
	"net"
)


func targetserver()  {
	//1.指定服务器通信协议、IP地址、Port端口，创建一个用于监听的socket---listener
	listener ,err := net.Listen("tcp","127.0.0.1:8089")
	if err != nil{
		fmt.Println("net.Listener err:",err)
		return
	}
	defer listener.Close()//关闭socket


	//2.阻塞监听客户端连接请求,成功建立连接，返回用于通信的socket---conn
	conn ,err := listener.Accept()
	if err != nil{
		fmt.Println("listener.Accept err:",err)
		return
	}
	defer conn.Close()//关闭socket

	//3.从conn套接字中获取文件名，写入缓存buf中
	buf := make([]byte,4096)
	n ,err := conn.Read(buf)
	if err != nil{
		fmt.Println("conn.Read err:",err)
		return
	}

	//4.从buf中提取文件名

	message := string(buf[:n])

	if n > 0 {
		fmt.Println(message)
	}	//5.回写给发送端ok
	conn.Write([]byte("ok"))

	//6.获取文件内容
	recivefile(conn)
}

func recivefile(conn net.Conn)  {

	//6.2从网络socket中读数据，写入本地文件中
	buf := make([]byte,4096)

	for  {

		n,_ := conn.Read(buf) //从conn中读数据到buf中
		if n == 0{  //判断是否读取数据完毕
			fmt.Println("接收文件完毕")
			return
		}

		//将buf中的数据写入到本地文件
		conn.Write(buf[:n])
	}

}
