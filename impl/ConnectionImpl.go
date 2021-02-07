package impl

import (
	"log"
	"net"
	"strconv"
	"zinx2server/interf"
)

type Connection struct {
	Conn     *net.TCPConn
	Id       int
	IsClosed bool
	router   interf.AbstractRouter
	//用来告知当前连接已经关闭
	Exit chan bool
}

func GetConnection(conn *net.TCPConn, id int, router interf.AbstractRouter) *Connection {
	connection := &Connection{
		Conn:     conn,
		Id:       id,
		IsClosed: false,
		router:   router,
		Exit:     make(chan bool, 1),
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
		// 现在不用HandleAPI了，直接将Connection包装成Request，调用router来处理请求
		request := &RequestImpl{
			connection: c,
			data:       buffer[:count],
		}
		// TODO 有一个疑问，方法定义参数是接口，实际传参可以是接口实现类的指针
		func(request interf.AbstractRequest) {
			c.router.PreHandler(request)
			c.router.DoHandle(request)
			c.router.PostHandler(request)
		}(request)

	}
}
