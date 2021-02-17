package interf

/*
	编解码器
	首先要明确一点，我一开始并不知道数据包的内容大小，我只知道长度的偏移量是多少，假设数据包如下
	[(固定的偏移量，能够获取长度)(实际内容)]
	我不能直接获取[]的内容，因为我根本不知道[]有多大，直接获取还是会拆包
	但是我知道(固定的偏移量，能够获取长度)有多大，因为(固定的偏移量，能够获取长度)是length(uint64)与id(uint64)转换成字节后的流，大小是固定的
	所以必须分两步读取，第一次获取16个字节，得到(固定的偏移量，能够获取长度)里的内容长度
	第二次获取长度同等大小的字节数，这样就能完整地拿到(实际内容)了
*/
type AbstractEndecoder interface {
	Encode(message AbstractMessage) ([]byte, error)
	DecodeLength(headData []byte) (AbstractMessage, error)
}
