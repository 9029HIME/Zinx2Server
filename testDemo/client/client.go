package main

import (
	"fmt"
	"net"
	"strconv"
	"time"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:7001")

	if err != nil {
		fmt.Println("dial err: ", err)
	}

	var flag int = 1

	for {
		_, err := conn.Write([]byte(fmt.Sprintf("helloworld:%s", strconv.Itoa(flag))))
		if err != nil {
			fmt.Println("write err: ", err)
			return
		}
		buffer := make([]byte, 512)
		count, err := conn.Read(buffer)
		if err != nil {
			fmt.Println("read err", err)
			return
		}
		fmt.Println("from server: ", string(buffer[:count]))
		flag++
		time.Sleep(time.Second * 1)
	}
}
