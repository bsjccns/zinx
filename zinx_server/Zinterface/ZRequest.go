package Zinterface

type ZRequest interface {
	GetConn() Iconnection
	GetData() []byte
	GetMsgId() uint32
}
