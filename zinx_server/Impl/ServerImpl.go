package Impl

import (
	"fmt"
	"net"
	"zinx/zinx_server/Zinterface"
	"zinx/zinx_server/util"
)

type Server struct {
	Ip          string
	Port        uint32
	Name        string
	RH          Zinterface.IRouterHolder
	ConnManager Zinterface.IConnManager

	OnConnStart func(conn Zinterface.Iconnection)
	OnConnStop  func(conn Zinterface.Iconnection)
}

func NewServe() (res *Server) {
	util.Config1.AnalysisConfig()
	fmt.Println("NewServe", util.Config1)
	return &Server{
		Ip:          util.Config1.Ip,
		Port:        util.Config1.Port,
		Name:        util.Config1.Name,
		RH:          NewRouterHolder(),
		ConnManager: NewConnManager(),
	}
}

//func callback(conn *net.TCPConn, b []byte, l int) error {
//	_, err := conn.Write(b[:l])
//	if err != nil {
//		fmt.Println("回写错误")
//		return err
//	}
//
//	return nil
//}

func (s *Server) Start() {

	go func() {
		s.RH.StartWorkPool()
		// 解析地址（也能解析域名）
		addr, err := net.ResolveTCPAddr("tcp4", fmt.Sprintf("%s:%d", s.Ip, s.Port))
		if err != nil {
			fmt.Println("服务器启动错误 [解析tcp地址错误]", err)
			return
		}
		//监听地址
		listener, err := net.ListenTCP("tcp4", addr)
		if err != nil {
			fmt.Println("ListenTcp 出错", err)
		}
		fmt.Printf("%s Starting Listen %s:%d\n", s.Name, s.Ip, s.Port)
		var cid uint32 = 1
		// 接受客户端连接，拿到连接对象
		for {
			tcpConnect, err := listener.AcceptTCP()
			if err != nil {
				fmt.Println("Accept Err", err)
				continue
			}
			if s.ConnManager.Len() > util.Config1.MaxConn {
				fmt.Println("-----------------------------当前连接数已达上限,拒绝连接！-------------------------")
				tcpConnect.Close()
				continue
			}
			fmt.Println("客户端连接成功 conId=", cid, " addr=", tcpConnect.RemoteAddr().String())
			conn := NewConnection(s, tcpConnect, cid, s.RH)
			cid++
			go conn.Start()
		}
	}()

}
func (s *Server) Serve() {
	//开始监听
	s.Start()
	// TODO 额外业务

	// 阻塞
	select {}
}
func (s *Server) Stop() {
	fmt.Println("Server Stop, clear all conn")
	s.ConnManager.Clear()
}

func (s *Server) AddRouter(msgId uint32, R Zinterface.ZRouter) {
	err := s.RH.AddRouter(msgId, R)
	if err != nil {
		fmt.Println("AddRouter err", err)
	}
}

func (s *Server) GetConnManager() Zinterface.IConnManager {
	return s.ConnManager
}

func (s *Server) SetOnConnStart(callBack func(conn Zinterface.Iconnection)) {
	s.OnConnStart = callBack

}

func (s *Server) SetOnConnStop(callback func(conn Zinterface.Iconnection)) {
	s.OnConnStop = callback
}

func (s *Server) CallOnConnStart(conn Zinterface.Iconnection) {
	if s.OnConnStart != nil {
		s.OnConnStart(conn)
	}
}

func (s *Server) CallOnConnStop(conn Zinterface.Iconnection) {
	if s.OnConnStop != nil {
		s.OnConnStop(conn)
	}
}
