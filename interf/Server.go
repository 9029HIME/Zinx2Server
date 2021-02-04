package interf

type AbstractServer interface {
	//启动
	Start()

	//运行
	Serve()

	//停止
	Stop()

	//添加路由器
	AddRouter(router AbstractRouter)
}
