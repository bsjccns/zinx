package Zinterface

type Iserver interface {
	Start()

	Serve()

	Stop()

	AddRouter(msgId uint32, R ZRouter)

	GetConnManager() IConnManager

	SetOnConnStart(func(conn Iconnection))

	SetOnConnStop(func(conn Iconnection))

	CallOnConnStart(conn Iconnection)

	CallOnConnStop(conn Iconnection)
}
