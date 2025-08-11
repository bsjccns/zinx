package Zinterface

type IDataPack interface {
	GetPkgHeadLen() uint32
	Pack(message IMessage) ([]byte, error)
	UnPack([]byte) (IMessage, error)
}
