package interf

type AbstractServer interface {
	//启动
	Start()

	//运行
	Serve()

	//停止
	Stop()

	//消息分发器
	AddMsgHandler(router AbstractMsgHandler) AbstractServer

	//
	AddEndecoder(endecoder AbstractEndecoder) AbstractServer
}
