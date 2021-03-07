package impl

import (
	"io"
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

	if _, err := c.Conn.Write(binaryContent); err != nil {
		log.Printf("id为%s的连接在写数据时发生异常：%s", strconv.Itoa(c.Id), err.Error())
	}
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
		//buffer := make([]byte, 512)
		//count, err := c.Conn.Read(buffer)
		//if err != nil {
		//	endFlag++
		//	log.Printf("%s read error:%s", strconv.Itoa(c.Id), err)
		//	if endFlag == 4 {
		//		log.Printf("ID为%s的连接即将关闭，异常原因是：%s\n", strconv.Itoa(c.Id), err.Error())
		//		break
		//	}
		//	continue
		//}

		// 现在不是直接读数据到缓冲区，而是用自定义的endecoder来解码字节流，防止拆包粘包 on 2021-02-17
		endecoder := new(TlvEndecoder)
		// ID和Length占了16个字节（两个uint64）
		headData := make([]byte, 16)
		if _, err := io.ReadFull(c.Conn, headData); err != nil {
			log.Printf("id为%s的连接在读消息头时发生异常：%s\n", strconv.Itoa(c.Id), err.Error())
			endFlag++
			if endFlag == 4 {
				log.Printf("ID为%s的连接即将关闭，异常原因是：%s\n", strconv.Itoa(c.Id), err.Error())
				break
			}
			continue
		}

		msg, err := endecoder.DecodeLength(headData)

		if err != nil {
			log.Printf("id为%s的连接在消息头转换Message时发生异常：%s\n", strconv.Itoa(c.Id), err.Error())
		}
		// 将获取到的msg进行二次读取
		if msg.GetLength() > 0 {
			content := make([]byte, msg.GetLength())
			if _, err := io.ReadFull(c.Conn, content); err != nil {
				log.Printf("id为%s的连接在消息头转换Message时发生异常：%s\n", strconv.Itoa(c.Id), err.Error())
				continue
			}
			msg.SetData(content)
			log.Printf("本次接收的消息ID是%s，消息长度是%s，消息内容是%s", strconv.Itoa(int(msg.GetId())), strconv.Itoa(int(msg.GetLength())),
				string(msg.GetData()))
		}

		// 现在不用HandleAPI了，直接将Connection包装成Request，调用router来处理请求
		request := &Request{
			connection: c,
			msg:        msg,
		}
		/**
		TODO 有一个疑问，方法定义参数是接口，实际传参可以是接口实现类的指针
		这又回到一个基础知识点：Go里超集可以转为子集，但子集不能转为超集。参数定义为超集，可以传子集，参数定义为子集，无法传超集
		对于interface来说，如果子集通过指针实现方法，那么只认该子集的指针
		*/
		func(request interf.AbstractRequest) {
			c.router.PreHandler(request)
			c.router.DoHandle(request)
			c.router.PostHandler(request)
		}(request)

	}
}
