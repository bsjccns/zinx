package Impl

import (
	"errors"
	"fmt"
	"zinx/zinx_server/Zinterface"
	"zinx/zinx_server/util"
)

type RouterHolder struct {
	routerMap  map[uint32]Zinterface.ZRouter
	WorkerChan []chan Zinterface.ZRequest
}

func NewRouterHolder() *RouterHolder {
	RH := &RouterHolder{
		routerMap: make(map[uint32]Zinterface.ZRouter),
	}
	if util.Config1.TaskQueueSize > 0 {
		RH.WorkerChan = make([]chan Zinterface.ZRequest, util.Config1.WorkerPoolSize)
	}
	return RH
}

func (RH *RouterHolder) DoRouter(r Zinterface.ZRequest) {
	if router, ok := RH.routerMap[r.GetMsgId()]; ok {
		router.PreHandle(r)
		router.Handle(r)
		router.PostHandle(r)
	} else {
		fmt.Println("msgId not exist,无法处理该消息")
	}
}
func (RH *RouterHolder) AddRouter(msgId uint32, R Zinterface.ZRouter) error {
	if _, ok := RH.routerMap[msgId]; ok {
		return errors.New("msgId exist")
	}
	RH.routerMap[msgId] = R
	fmt.Println("AddRouter success, msgId:", msgId)
	return nil
}

func (RH *RouterHolder) StartWorkPool() {
	if util.Config1.WorkerPoolSize <= 0 || util.Config1.TaskQueueSize <= 0 {
		fmt.Println("未开启工作池")
		return
	}
	fmt.Println("开启工作池 workerPoolSize=", util.Config1.WorkerPoolSize, "taskQueueSize=", util.Config1.TaskQueueSize)
	for i := 0; i < int(util.Config1.WorkerPoolSize); i++ {
		RH.WorkerChan[i] = make(chan Zinterface.ZRequest, util.Config1.TaskQueueSize)
		go RH.StartOneWorker(i)
	}
}

func (RH *RouterHolder) StartOneWorker(workerId int) {
	fmt.Println("workerId=", workerId, "开始执行任务")
	for {
		//监听 自己的Channel 获取任务
		select {
		case request := <-RH.WorkerChan[workerId]:
			fmt.Println("workerId=", workerId, "开始执行任务")
			RH.DoRouter(request)
		}
	}
}

func (RH *RouterHolder) DoMsgByWorkerPool(request Zinterface.ZRequest) {
	workerId := request.GetConn().GetConnectId() % util.Config1.WorkerPoolSize
	RH.WorkerChan[workerId] <- request
}
