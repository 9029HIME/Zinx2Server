package main

import "zinx2server/impl"

/**
设计得有点像Netty，最近学Netty学魔怔了
总的来说：
不同的端对端只有一个连接（Connection）
一个连接有多个请求（Request），多个路由（Router），通过连接进行读与写操作
一个请求包含一个信息（Message）
一个信息的编码与解码通过一个编解码器（Endecoder）
服务端会根据一个信息的ID调用对应的路由（Router）处理
*/
func main() {
	/*
		基于手动输入信息打开
	*/

	//impl.Launch(
	//	"tcp4",
	//	"myServer",
	//	"localhost",
	//	"7001").AddRouter(new(impl.PrintCallBackRouter)).Serve()

	/*
		基于配置文件打开，传""则用默认值
	*/
	impl.Config("").
		AddMsgHandler(
			impl.NewMsgHandler().
				AddRouter(uint64(1), new(impl.PrintCallBackRouter))).
		Serve()

}
