package impl

import (
	"bytes"
	"encoding/binary"
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
