package impl

import "Zinx2Server/interf"

type RequestImpl struct {
	connection *Connection
	data       []byte
}

func (request *RequestImpl) GetConnection() interf.AbstractConnection {
	return request.connection
}

func (request *RequestImpl) GetData() []byte {
	return request.data
}
