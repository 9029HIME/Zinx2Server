package impl

import (
	"fmt"
	"zinx2server/interf"
)

/**
为什么要弄一个DefaultAbstractRouter？
因为实际使用的Router可能只想实现其中一个方法，直接实现接口的话并不允许，所以直接继承DefaultAbstractRouter，重写想要实现的方法即可。
*/
type DefaultAbstractRouter struct {
}

func (router *DefaultAbstractRouter) PreHandler(request interf.AbstractRequest) {

}

func (router *DefaultAbstractRouter) DoHandle(request interf.AbstractRequest) {

}

func (router *DefaultAbstractRouter) PostHandler(request interf.AbstractRequest) {

}

type PrintCallBackRouter struct {
	DefaultAbstractRouter
}

func (router *PrintCallBackRouter) PreHandler(request interf.AbstractRequest) {
	//fmt.Println("重写了Prehandler")
	//request.GetConnection().GetConn().Write([]byte("服务端这里调用了PreHandler\n"))
}

func (router *PrintCallBackRouter) DoHandle(request interf.AbstractRequest) {
	//fmt.Println(string(request.GetData()))
	request.GetConnection().Write(1, []byte(fmt.Sprintf("已经收到你的消息了:%s\n", string(request.GetData()))))
	//request.GetConnection().GetConn().Write([]byte("服务端这里调用了DoHandle\n"))
}

func (router *PrintCallBackRouter) PostHandler(request interf.AbstractRequest) {
	//fmt.Println("重写了PostHandler")
	//request.GetConnection().GetConn().Write([]byte("服务端这里调用了PostHandler\n"))
}
