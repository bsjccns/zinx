package Impl

type Message struct {
	Id      uint32
	DataLen uint32
	Data    []byte
}

func (M *Message) GetData() []byte {
	return M.Data
}
func (M *Message) SetData(d []byte) {
	M.Data = d
}
func (M *Message) GetMsgId() uint32 {
	return M.Id
}
func (M *Message) SetMsgId(id uint32) {
	M.Id = id
}
func (M *Message) GetMsgLen() uint32 {
	return M.DataLen
}
func (M *Message) SetMsgLen(L uint32) {
	M.DataLen = L
}
