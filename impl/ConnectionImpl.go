package impl

import (
	"log"
	"net"
	"strconv"
	"zinx2server/interf"
)

type Connection struct {
	Conn      *net.TCPConn
	Id        int
	IsClosed  bool
	handler   interf.AbstractMsgHandler
	endecoder interf.AbstractEndecoder
	//用来告知当前连接已经关闭
	Exit chan bool
	// “读协程”将写信息写到这个channel里，等待“写协程”读取并写给客户端
	MsgChan chan []byte
}

func GetConnection(conn *net.TCPConn, id int, handler interf.AbstractMsgHandler, endecoder interf.AbstractEndecoder) *Connection {
	connection := &Connection{
		Conn:     conn,
		Id:       id,
		IsClosed: false,
		// 统一连接可能会有多种消息，每种消息有多种路由，所以也用msgHandler来管理
		handler:   handler,
		endecoder: endecoder,
		Exit:      make(chan bool),
		MsgChan:   make(chan []byte),
	}
	// 默认使用tlv编解码器
	if connection.endecoder == nil {
		connection.endecoder = new(TlvEndecoder)
	}
	return connection
}

//启动
func (c *Connection) Start() {
	log.Printf("ID为%s的连接开启", c.Id)
	// 当一个连接启动时，服务端要为连接开两个协程，一个用来读一个用来写
	go c.ConnRead()
	// 写协程
	go c.ConnWrite()
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
func (c *Connection) Write(id uint64, data []byte) {
	// 根据编解码器对信息编码，防止拆包粘包 on 2021-02-17
	if c.IsClosed {
		log.Printf("对一个已关闭通道:%s进行写操作", strconv.Itoa(c.Id))
	}

	endecoder := new(TlvEndecoder)

	binaryContent, err := endecoder.Encode(&Message{
		Id:     id,
		Data:   data,
		Length: uint64(len(data)),
	})

	if err != nil {
		log.Printf("id为%s的连接在编码数据时发生异常：%s", strconv.Itoa(c.Id), err.Error())
	}

	// 往消息通道里传输数据，等待“写协程”读取
	c.MsgChan <- binaryContent
}

//停止
func (c *Connection) Stop() {

	if c.IsClosed == true {
		log.Printf("尝试关闭一个已关闭的连接：%s\n", strconv.Itoa(c.Id))
		return
	}

	c.Conn.Close()

	// 关掉后，读协程就会读到内容，并关闭读协程
	close(c.Exit)

}

func (c *Connection) ConnRead() {
	log.Printf("ID为%s的连接正在读\n", strconv.Itoa(c.Id))
	defer c.Stop()
	for {
		// 这里调用服务器的endecoder来解析信息，此时会读取网络缓冲区的字节数据，即可能会阻塞
		msg, err := c.endecoder.GetMessage(c)
		if err != nil || msg == nil {
			break
		}

		// 现在不用HandleAPI了，直接将Connection包装成Request，调用router来处理请求
		request := &Request{
			connection: c,
			msg:        msg,
		}
		/**
		这又回到一个基础知识点：Go里超集可以转为子集，但子集不能转为超集。参数定义为超集，可以传子集，参数定义为子集，无法传超集
		对于interface来说，如果子集通过指针实现方法，那么只认该子集的指针
		*/
		func(request interf.AbstractRequest) {
			c.handler.Dispatch(request)
		}(request)
	}
}

func (c *Connection) ConnWrite() {
	log.Printf("ID为%s的连接正在写\n", strconv.Itoa(c.Id))
	for {
		select {
		case pendingWrite := <-c.MsgChan:
			if _, err := c.GetConn().Write(pendingWrite); err != nil {
				return
			}
		case <-c.Exit:
			// 这里的return指结束ConnWrite()方法，继续上层返回
			return
		}
	}
}
