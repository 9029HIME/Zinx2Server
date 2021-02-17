package impl

import (
	"fmt"
	"log"
	"net"
	"zinx2server/config"
	_interface "zinx2server/interf"
)

//AbstractServer的实现类
type Server struct {
	ipVersion  string
	serverName string
	host       string
	port       string
	//TODO 目前还是一个服务器一个router，后期会改成一个服务器多个router
	router    _interface.AbstractRouter
	endecoder _interface.AbstractEndecoder
}

func CallBack(conn *net.TCPConn, content []byte, length int) error {
	_, err := conn.Write(content[:length])
	if err != nil {
		return err
	}
	return nil
}

func (s *Server) Start() {

	log.Printf("[Start] Server listener at IP :%s,Port :%s,IPversion :%s\n", s.host, s.port, s.ipVersion)
	listener, err := net.Listen(s.ipVersion, fmt.Sprintf("%s:%s", s.host, s.port))
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
		connection := GetConnection(tcpConn, id, s.router)
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

func (s *Server) AddRouter(router _interface.AbstractRouter) _interface.AbstractServer {
	s.router = router
	return s
}

func (s *Server) AddEndecoder(endecoder _interface.AbstractEndecoder) _interface.AbstractServer {
	s.endecoder = endecoder
	return s
}

func Launch(ipVersion string, serverName string, host string, port string) _interface.AbstractServer {
	server := &Server{
		ipVersion:  ipVersion,
		serverName: serverName,
		host:       host,
		port:       port,
	}
	return server
}

func Config(configPath string) _interface.AbstractServer {
	config := config.Init(configPath)
	server := &Server{
		ipVersion:  config.IPVersion,
		serverName: config.ServerName,
		host:       config.Host,
		port:       config.Port,
	}
	return server
}
