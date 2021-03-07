package interf

type AbstractMsgHandler interface {

	/**
	为Msg添加特定的Router
	*/
	AddRouter(id uint64, router AbstractRouter)

	/**
	用Msg对应的Router来处理Msg
	*/
	DoHandle(request AbstractRequest)
}
