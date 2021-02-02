package impl

import (
	_interface "Zinx2Server/interface"
	"fmt"
	"log"
	"net"
	"strconv"
)

//AbstractServer的实现类
type Server struct {
	ipVersion  string
	serverName string
	host       string
	port       int
}

func CallBack(conn *net.TCPConn, content []byte, length int) error {
	_, err := conn.Write(content[:length])
	if err != nil {
		return err
	}
	return nil
}

func (s *Server) Start() {
	log.Printf("[Start] Server listener at IP :%s,Port :%s,IPversion :%s\n", s.host, strconv.Itoa(s.port), s.ipVersion)
	//TODO ipversion
	listener, err := net.Listen(s.ipVersion, fmt.Sprintf("%s:%s", s.host, strconv.Itoa(s.port)))
	if err != nil {
		fmt.Println("listen err:", err)
		return
	}
	log.Printf("[Start] Server:%s start\n", s.serverName)
	var id int = 0
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("accept err:", err)
			continue
		}
		tcpConn := conn.(*net.TCPConn)
		// 包装成前面定义的Connection
		connection := GetConnection(tcpConn, id, CallBack)
		connection.Start()
		id++
	}

}
func (s *Server) Serve() {
	go s.Start()
	//空select阻塞着
	select {}
}

func (s *Server) Stop() {

}

func Launch(ipVersion string, serverName string, host string, port int) _interface.AbstractServer {
	server := &Server{
		ipVersion:  ipVersion,
		serverName: serverName,
		host:       host,
		port:       port,
	}
	return server
}
