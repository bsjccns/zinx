package Zinterface

import "net"

type Iconnection interface {
	Start()
	Stop()
	SendMsg(msgId uint32, data []byte) error
	GetConnection() *net.TCPConn
	GetConnectId() uint32
	GetAddr() net.Addr
	GetClosedTag() bool

	SetProperty(key string, value interface{})
	GetProperty(key string) (interface{}, bool)
	RemoveProperty(key string)
}

//type HandleFunc func(conn *net.TCPConn, b []byte, l int) error
