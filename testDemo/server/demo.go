package main

import "zinx2server/impl"

/**
设计得有点像Netty，最近学Netty学魔怔了
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
	impl.Config(
		"").AddRouter(new(impl.PrintCallBackRouter)).Serve()

}
