package impl

import "zinx2server/interf"

type MsgHandler struct {
	maps map[uint64]interf.AbstractRouter
}

func CreateHandler() interf.AbstractMsgHandler {
	return &MsgHandler{
		maps: make(map[uint64]interf.AbstractRouter),
	}
}

func (msgHandler *MsgHandler) AddRouter(id uint64, router interf.AbstractRouter) {
	_, ok := msgHandler.maps[id]
	if ok {
		panic("多个路由器不能绑定同一个id")
	}
	msgHandler.maps[id] = router
}

func (msgHandler *MsgHandler) DoHandle(request interf.AbstractRequest) {
	msgId := request.GetMsgId()
	handler := msgHandler.maps[msgId]

	handler.PreHandler(request)
	handler.DoHandle(request)
	handler.PostHandler(request)
}
