package impl

import (
	"bytes"
	"encoding/binary"
	"io"
	"log"
	"strconv"
	"zinx2server/interf"
)

type TlvEndecoder struct {
}

// TLV编码 message → []byte
func (endecoder *TlvEndecoder) Encode(message interf.AbstractMessage) ([]byte, error) {
	data := bytes.NewBuffer(make([]byte, 0))
	length := int64(message.GetLength())
	id := int64(message.GetId())
	//使用小端存储，回忆一下小端和大端的区别：小端：11100 大端：00111
	err := binary.Write(data, binary.LittleEndian, &length)
	if err != nil {
		return nil, err
	}

	err = binary.Write(data, binary.LittleEndian, &id)
	if err != nil {
		return nil, err
	}

	err = binary.Write(data, binary.LittleEndian, message.GetData())
	if err != nil {
		return nil, err
	}
	return data.Bytes(), nil
}

func (endecoder *TlvEndecoder) DecodeLength(headData []byte) (interf.AbstractMessage, error) {
	buffer := bytes.NewBuffer(headData)
	message := new(Message)

	/*
			先读长度
			这个方法意思是：从buffer里通过小端读取，读uint64大小的字节内容，然后赋给&message.Length，为什么要传指针？因为方法里面只判断指针，同时还要修改值。可以点进去看看具体实现
		 	&message.length这种写法拿的是length的地址
	*/
	if err := binary.Read(buffer, binary.LittleEndian, &message.Length); err != nil {
		return nil, err
	}

	/*
		再读ID
	*/
	if err := binary.Read(buffer, binary.LittleEndian, &message.Id); err != nil {
		return nil, err
	}
	return message, nil
}

func (endecoder *TlvEndecoder) GetMessage(c interf.AbstractConnection) (interf.AbstractMessage, error) {
	var msg interf.AbstractMessage
	var err error = nil
	// ID和Length占了16个字节（两个uint64）
	headData := make([]byte, 16)
	if _, err = io.ReadFull(c.GetConn(), headData); err != nil {
		log.Printf("id为%s的连接在读消息头时发生异常：%s\n", strconv.Itoa(c.GetId()), err.Error())
	}

	msg, err = endecoder.DecodeLength(headData)

	if err != nil {
		log.Printf("id为%s的连接在消息头转换Message时发生异常：%s\n", strconv.Itoa(c.GetId()), err.Error())
	}
	// 将获取到的msg进行二次读取
	if msg.GetLength() > 0 {
		content := make([]byte, msg.GetLength())
		if _, err = io.ReadFull(c.GetConn(), content); err != nil {
			log.Printf("id为%s的连接在消息头转换Message时发生异常：%s\n", strconv.Itoa(c.GetId()), err.Error())
		}
		msg.SetData(content)
		log.Printf("本次接收的消息ID是%s，消息长度是%s，消息内容是%s", strconv.Itoa(int(msg.GetId())), strconv.Itoa(int(msg.GetLength())),
			string(msg.GetData()))
	}

	return msg, err
}
