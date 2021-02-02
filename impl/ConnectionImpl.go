package impl

import (
	_interface "Zinx2Server/interface"
	"log"
	"net"
	"strconv"
)

type Connection struct {
	Conn      *net.TCPConn
	Id        int
	IsClosed  bool
	HandleAPI _interface.HandleFunc
	//用来告知当前连接已经关闭
	Exit chan bool
}

func GetConnection(conn *net.TCPConn, id int, handleAPI _interface.HandleFunc) *Connection {
	connection := &Connection{
		Conn:      conn,
		Id:        id,
		IsClosed:  false,
		HandleAPI: handleAPI,
		Exit:      make(chan bool, 1),
	}

	return connection
}

//启动
func (c *Connection) Start() {
	log.Printf("ID为%s的连接开启", c.Id)
	// 当一个连接启动时，服务端要为连接开两个协程，一个用来读一个用来写
	go c.ConnRead()
	// TODO 写协程
}

//获取连接id
func (c *Connection) GetId() int {
	return c.Id
}

//获取Conn
func (c *Connection) GetConn() *net.TCPConn {
	return c.Conn
}

//获取ip
func (c *Connection) GetIP() net.Addr {
	return c.Conn.RemoteAddr()
}

//发送数据
func (c *Connection) Write() {

}

//停止
func (c *Connection) Stop() {
	log.Printf("ID为%s的连接即将关闭\n", strconv.Itoa(c.Id))

	if c.IsClosed == true {
		log.Printf("尝试关闭一个已关闭的连接：%s\n", strconv.Itoa(c.Id))
		return
	}

	c.Conn.Close()

	close(c.Exit)

}

func (c *Connection) ConnRead() {
	log.Printf("ID为%s的连接正在读\n", strconv.Itoa(c.Id))
	defer c.Stop()
	endFlag := 0

	for {
		buffer := make([]byte, 512)
		count, err := c.Conn.Read(buffer)
		if err != nil {
			endFlag++
			log.Printf("%s read error:%s", strconv.Itoa(c.Id), err)
			if endFlag == 4 {
				log.Printf("ID为%s的连接即将关闭，异常原因是：%s\n", strconv.Itoa(c.Id), err.Error())
				break
			}
			continue
		}

		if err = c.HandleAPI(c.Conn, buffer, count); err != nil {
			log.Printf("ID为%s的连接即将关闭，异常原因是：%s\n", strconv.Itoa(c.Id), err.Error())
			break
		}
	}
}
