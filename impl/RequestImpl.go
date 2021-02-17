package impl

import "zinx2server/interf"

type Request struct {
	connection interf.AbstractConnection
	msg        interf.AbstractMessage
}

func (request *Request) GetConnection() interf.AbstractConnection {
	return request.connection
}

func (request *Request) GetData() []byte {
	return request.msg.GetData()
}

func (request *Request) GetMsgId() uint64 {
	return request.msg.GetId()
}
