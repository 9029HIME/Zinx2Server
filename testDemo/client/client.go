package main

import (
	"fmt"
	"net"
	"strconv"
	"time"
	"zinx2server/impl"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:6789")

	if err != nil {
		fmt.Println("dial err: ", err)
	}

	var flag uint64 = 1

	endecoder := new(impl.TlvEndecoder)
	contentTemplate := "helloworldhelloworldhelloworldhelloworldhelloworldhelloworldhelloworldhelloworldhelloworldhelloworldhelloworldhelloworldhelloworldhelloworldhelloworldhelloworldhelloworldhelloworldhelloworldhelloworld"

	for {
		contentTemplate = fmt.Sprint(contentTemplate, "add")
		data := []byte(fmt.Sprintf("%s:%s", contentTemplate, strconv.Itoa(int(flag))))
		fmt.Println("准备发送的数据长度：", len(data))

		binaryContent, err := endecoder.Encode(&impl.MessageImpl{
			Id:     flag,
			Data:   data,
			Length: uint64(len(data)),
		})

		_, err = conn.Write(binaryContent)
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
