package Impl

import (
	"errors"
	"fmt"
	"io"
	"net"
	"sync"
	"zinx/zinx_server/Zinterface"
	"zinx/zinx_server/util"
)

type Connection struct {
	Conn     *net.TCPConn
	ConnId   uint32
	IsClosed bool
	//业务函数
	RH Zinterface.IRouterHolder
	//用于通知writer连接关闭的管道
	ExitChan chan bool
	// 读协程和写协程通过管道通信
	msgChan chan []byte
	//当前连接所属服务器
	server   Zinterface.Iserver
	property map[string]interface{}

	lock sync.RWMutex
}

func (c *Connection) GetClosedTag() bool {
	return c.IsClosed
}

func NewConnection(s Zinterface.Iserver, c *net.TCPConn, cId uint32, callBackFunc Zinterface.IRouterHolder) *Connection {
	connection := &Connection{
		Conn:     c,
		ConnId:   cId,
		ExitChan: make(chan bool, 1),
		RH:       callBackFunc,
		IsClosed: false,
		msgChan:  make(chan []byte),
		server:   s,
		property: make(map[string]interface{}),
	}
	s.GetConnManager().Add(connection)
	return connection
}

// Start 启动连接 开始业务操作
func (c *Connection) Start() {

	// 读协程
	go c.StartReader()
	// 写协程
	go c.StartWriter()
	c.server.CallOnConnStart(c)
}
func (c *Connection) Stop() {
	if c.IsClosed {
		return
	}
	c.IsClosed = true
	c.ExitChan <- true
	c.server.CallOnConnStop(c)
	err := c.Conn.Close()
	if err != nil {
		return
	}
	c.server.GetConnManager().Remove(c)
	close(c.ExitChan)
	close(c.msgChan)
	fmt.Println("reader goroutine exit: connId=", c.ConnId)
}
func (c *Connection) SendMsg(msgId uint32, b []byte) error {
	if c.IsClosed {
		return errors.New("连接已关闭")
	}
	msg := &Message{
		Id:      msgId,
		Data:    b,
		DataLen: uint32(len(b)),
	}
	dp := &DataPack{}
	bytes, err := dp.Pack(msg)
	if err != nil {
		fmt.Println("pack error!")
		return err
	}
	// 发送数据给writer
	c.msgChan <- bytes
	return nil

}
func (c *Connection) GetConnection() *net.TCPConn {
	return c.Conn
}
func (c *Connection) GetConnectId() uint32 {
	return c.ConnId
}
func (c *Connection) GetAddr() net.Addr {
	return c.Conn.RemoteAddr()
}

func (c *Connection) StartReader() {
	fmt.Println("reader goroutine is running")

	defer c.Stop()
	for !c.IsClosed {
		// 读取id和数据长度
		dataPack := &DataPack{}
		b := make([]byte, dataPack.GetPkgHeadLen())
		_, err := io.ReadFull(c.Conn, b)
		if err != nil {
			fmt.Println("读取出错!")
			c.ExitChan <- true
			break
		}
		msg, err := dataPack.UnPack(b)
		if err != nil {
			fmt.Println("拆包出错!")
			continue
		}
		if msg.GetMsgLen() > 0 {
			msg.SetData(make([]byte, msg.GetMsgLen()))
			_, err := io.ReadFull(c.Conn, msg.GetData())
			if err != nil {
				fmt.Println("读取数据出错!")
				continue
			}
		}
		r := Request{
			c,
			msg,
		}
		if util.Config1.WorkerPoolSize > 0 {
			c.RH.DoMsgByWorkerPool(&r)
		} else {
			//执行业务
			go func() {
				c.RH.DoRouter(&r)
			}()
		}
	}

}

func (c *Connection) StartWriter() {
	fmt.Println("writer goroutine is running")
	defer fmt.Println("writer exit: connId=", c.ConnId)
	for {
		select {
		case msg := <-c.msgChan:
			if _, err := c.Conn.Write(msg); err != nil {
				fmt.Println("writer send msg error")
				continue
			}
		case <-c.ExitChan:
			return
		}
	}

}

func (c *Connection) SetProperty(key string, value interface{}) {
	c.lock.Lock()
	defer c.lock.Unlock()
	c.property[key] = value
}
func (c *Connection) GetProperty(key string) (interface{}, bool) {

	c.lock.RLock()
	defer c.lock.RUnlock()
	if value, ok := c.property[key]; ok {
		return value, true
	} else {
		return nil, false
	}
}
func (c *Connection) RemoveProperty(key string) {
	c.lock.Lock()
	defer c.lock.Unlock()
	delete(c.property, key)
}
