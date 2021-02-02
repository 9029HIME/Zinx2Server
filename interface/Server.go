package _interface

import "net"

type AbstractServer interface {
	//启动
	Start()

	//运行
	Serve()

	//停止
	Stop()
}

type AbstractConnection interface {
	//启动
	Start()

	//获取连接id
	GetId() int

	//获取Conn
	GetConn() *net.TCPConn

	//获取ip
	GetIP() net.Addr

	//发送数据
	Write()

	//停止
	Stop()
}

// Connection处理业务逻辑的方法，参数2是数据，参数3是数据长度
type HandleFunc func(*net.TCPConn, []byte, int) error
