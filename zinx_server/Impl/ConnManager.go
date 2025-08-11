package Impl

import (
	"errors"
	"sync"
	"zinx/zinx_server/Zinterface"
)

type ConnManager struct {
	connections map[uint32]Zinterface.Iconnection
	connLock    sync.RWMutex
}

func NewConnManager() *ConnManager {
	return &ConnManager{
		connections: make(map[uint32]Zinterface.Iconnection),
	}
}

func (cm *ConnManager) Add(c Zinterface.Iconnection) {
	cm.connLock.Lock()
	defer cm.connLock.Unlock()
	cm.connections[c.GetConnectId()] = c
}
func (cm *ConnManager) Remove(c Zinterface.Iconnection) {
	cm.connLock.Lock()
	defer cm.connLock.Unlock()
	delete(cm.connections, c.GetConnectId())
}
func (cm *ConnManager) Get(connID uint32) (Zinterface.Iconnection, error) {
	cm.connLock.RLock()
	defer cm.connLock.RUnlock()
	if conn, ok := cm.connections[connID]; ok {
		return conn, nil
	} else {
		return nil, errors.New("连接不存在")
	}
}
func (cm *ConnManager) Len() uint32 {
	return uint32(len(cm.connections))
}
func (cm *ConnManager) Clear() {
	cm.connLock.Lock()
	defer cm.connLock.Unlock()
	for k, conn := range cm.connections {
		delete(cm.connections, k)
		conn.Stop()
	}
}
