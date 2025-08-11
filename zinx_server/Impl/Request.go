package Impl

import "zinx/zinx_server/Zinterface"

type Request struct {
	conn Zinterface.Iconnection
	msg  Zinterface.IMessage
}

func (R *Request) GetConn() Zinterface.Iconnection {
	return R.conn
}
func (R *Request) GetData() []byte {

	return R.msg.GetData()
}
func (R *Request) GetMsgId() uint32 {
	return R.msg.GetMsgId()
}
