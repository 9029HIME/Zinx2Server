package interf

type AbstractRequest interface {
	GetConnection() AbstractConnection
	GetData() []byte
	GetMsgId() uint64
}
