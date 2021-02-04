package main

import "Zinx2Server/impl"

func main() {
	server := impl.Launch("tcp4", "myServer", "localhost", 7001)
	router := new(impl.PrintCallBackRouter)
	server.AddRouter(router)
	server.Serve()
}
