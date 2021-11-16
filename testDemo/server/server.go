package main

import "zinx2server/impl"

/**
设计得有点像Netty，最近学Netty学魔怔了
总的来说：
不同的端对端只有一个连接（Connection）
一个Connection会在建立之初开辟两个协程：读协程、写协程，两者底层是通过tcp.conn操作。当调用Connection的写方法时，实际上是将消息写入“消息队列”内，写协程会不断循环select消息队列和ExitChan，直到有消息队列有消息了才通过底层的tcp.conn发送出去。
当Connection发生异常时，会给予engFlag次重试机会，当机会用尽后，会关闭底层的tcp.conn与ExitChan，写协程select ExitChan成功，中止写操作，并关闭相关资源。
一个Connection有多个请求（Request），一个请求有多个信息（Message），一个请求对应多个路由（Router）。
服务端会通过Handler根据一个Message的ID调用对应的Router处理，前提是服务端需要注册Handler
一个Message的编码与解码通过一个编解码器（Endecoder），编解码器最终会落实到Connection的读写
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
