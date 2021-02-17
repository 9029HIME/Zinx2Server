package interf

/**
Type-Length-Value
*/
type AbstractMessage interface {
	GetId() uint64
	SetId(id uint64)
	GetLength() uint64
	SetLength(length uint64)
	GetData() []byte
	SetData(data []byte)
}
