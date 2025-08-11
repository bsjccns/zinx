package Zinterface

type IRouterHolder interface {
	DoRouter(r ZRequest)
	AddRouter(msgId uint32, R ZRouter) error
	StartWorkPool()
	DoMsgByWorkerPool(request ZRequest)
}
