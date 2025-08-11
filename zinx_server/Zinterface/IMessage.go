package Zinterface

type IMessage interface {
	GetData() []byte
	SetData([]byte)
	GetMsgId() uint32
	SetMsgId(i uint32)
	GetMsgLen() uint32
	SetMsgLen(uint32)
}
