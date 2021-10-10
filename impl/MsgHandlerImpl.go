package impl

import (
	"log"
	"strconv"
	"zinx2server/interf"
)

type MsgHandler struct {
	maps map[uint64]interf.AbstractRouter
}

func NewMsgHandler() interf.AbstractMsgHandler {
	return &MsgHandler{
		map[uint64]interf.AbstractRouter{},
	}
}

func CreateHandler() interf.AbstractMsgHandler {
	return &MsgHandler{
		maps: make(map[uint64]interf.AbstractRouter),
	}
}

func (msgHandler *MsgHandler) AddRouter(id uint64, router interf.AbstractRouter) interf.AbstractMsgHandler {
	_, ok := msgHandler.maps[id]
	if ok {
		panic("多个路由器不能绑定同一个id")
	}
	msgHandler.maps[id] = router
	return msgHandler
}

func (msgHandler *MsgHandler) Dispatch(request interf.AbstractRequest) {
	msgId := request.GetMsgId()
	router := msgHandler.maps[msgId]
	if router == nil {
		log.Printf("id = %s的消息不合法，此消息丢弃", strconv.Itoa(int(msgId)))
	}

	router.PreHandler(request)
	router.DoHandle(request)
	router.PostHandler(request)
}
