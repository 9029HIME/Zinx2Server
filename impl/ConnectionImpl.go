package impl

import (
	_interface "Zinx2Server/interface"
	"net"
)

type Connection struct {
	Conn *net.TCPConn
	Id int
	IsClosed bool
	HandleAPI _interface.HandleFunc
	//用来告知当前连接已经关闭
	Exit chan bool
}

func GetConnection(conn *net.TCPConn,id int,handleAPI _interface.HandleFunc) *Connection{
	connection := &Connection{
		Conn: conn,
		Id: id,
		IsClosed: false,
		HandleAPI: handleAPI,
		Exit: make(chan bool,1),
	}

	return connection
}
