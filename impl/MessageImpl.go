package impl

type MessageImpl struct {
	Id     uint64
	Length uint64
	Data   []byte
}

func (message *MessageImpl) GetId() uint64 {
	return message.Id
}
func (message *MessageImpl) SetId(id uint64) {
	message.Id = id
}

func (message *MessageImpl) GetLength() uint64 {
	return message.Length
}
func (message *MessageImpl) SetLength(length uint64) {
	message.Length = length
}

func (message *MessageImpl) GetData() []byte {
	return message.Data
}

func (message *MessageImpl) SetData(data []byte) {
	message.Data = data
}
