package Zinterface

type IConnManager interface {
	Add(c Iconnection)
	Remove(c Iconnection)
	Get(connID uint32) (Iconnection, error)
	Len() uint32
	Clear()
}
