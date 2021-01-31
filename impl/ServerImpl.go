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

func (s *Server) Start() {
	log.Printf("[Start] Server listener at IP :%s,Port :%s,IPversion :%s\n", s.host, strconv.Itoa(s.port), s.ipVersion)
	//TODO ipversion
	listener, err := net.Listen(s.ipVersion, fmt.Sprintf("%s:%s", s.host, strconv.Itoa(s.port)))
	if err != nil {
		fmt.Println("listen err:", err)
		return
	}
	log.Printf("[Start] Server:%s start\n", s.serverName)
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("accept err:", err)
			continue
		}
		go func() {
			for {
				buffer := make([]byte, 512)
				count, err := conn.Read(buffer)
				if err != nil {
					fmt.Println("read err:", err)
					return
				}
				fmt.Println("from client: ", string(buffer[:count]))
				if _, err := conn.Write(buffer[:count]); err != nil {
					fmt.Println("write err:", err)
				}
			}
		}()
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
