package interf

import "net"

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
	Write(id uint64, data []byte)

	//停止
	Stop()
}
