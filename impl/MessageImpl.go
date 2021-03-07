package impl

type Message struct {
	Id     uint64
	Length uint64
	Data   []byte
}

func (message *Message) GetId() uint64 {
	return message.Id
}
func (message *Message) SetId(id uint64) {
	message.Id = id
}

func (message *Message) GetLength() uint64 {
	return message.Length
}
func (message *Message) SetLength(length uint64) {
	message.Length = length
}

func (message *Message) GetData() []byte {
	return message.Data
}

func (message *Message) SetData(data []byte) {
	message.Data = data
}
