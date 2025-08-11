package main

import (
	"fmt"
	"zinx/zinx_server/Impl"
	"zinx/zinx_server/Zinterface"
)

type PingRouter struct {
	Impl.BaseRouter
}

func (BR *PingRouter) PreHandle(r Zinterface.ZRequest) {
	//fmt.Println("pre")
}
func (BR *PingRouter) Handle(r Zinterface.ZRequest) {
	if r.GetConn().GetClosedTag() {
		return
	}
	// fmt.Println("handle:接收到客户端消息:", r.GetMsgId(), string(r.GetData()))
	//回显ping
	err := r.GetConn().SendMsg(1, []byte("ping [回显消息]"))
	if err != nil {
		fmt.Println(" send error:", err)
	}
}
func (BR *PingRouter) PostHandle(r Zinterface.ZRequest) {
	//fmt.Println("post")
}

func DoOnConnStart(conn Zinterface.Iconnection) {
	fmt.Println("连接建立")
}
func DoOnConnStop(conn Zinterface.Iconnection) {
	fmt.Println("连接断开")
}

func main() {
	//addr, err := net.ResolveTCPAddr("tcp4", "www.baidu.com:80")
	//if err != nil {
	//	fmt.Println("错误")
	//}
	//fmt.Println(addr)

	serve := Impl.NewServe()
	serve.SetOnConnStart(DoOnConnStart)
	serve.SetOnConnStop(DoOnConnStop)
	serve.AddRouter(1, &PingRouter{})
	serve.AddRouter(2, &PingRouter{})
	serve.AddRouter(3, &PingRouter{})
	serve.Serve()
	//fmt.Println("hello")
	//estDataPack()
}

//func estDataPack() {
//	fmt.Println("拆包工具测试")
//	listen, err := net.Listen("tcp", "127.0.0.1:8899")
//	if err != nil {
//		fmt.Printf("listen error:%v\n", err)
//	}
//
//	// 服务端
//	go func() {
//		conn, err := listen.Accept()
//		if err != nil {
//			fmt.Printf("accept error:%v", err)
//		}
//		go func(conn net.Conn) {
//			for {
//
//				// 拆包工具
//				dataPack := &Impl.DataPack{}
//				bytes := make([]byte, dataPack.GetPkgHeadLen())
//				_, err = io.ReadFull(conn, bytes)
//				if err != nil {
//					fmt.Printf("read head error:%v", err)
//				}
//				// 解出 消息长度 和消息id 存入message对象
//				message, err := dataPack.UnPack(bytes)
//				if err != nil {
//					fmt.Printf("unpack error:%v", err)
//				}
//				//msg := message.(*Message)
//				//再次根据消息长度读取数据
//				if message.GetMsgLen() > 0 {
//					data := message.GetData()
//					data = make([]byte, message.GetMsgLen())
//					_, err = io.ReadFull(conn, data)
//					if err != nil {
//						fmt.Printf("read data error:%v", err)
//					}
//				}
//
//				fmt.Printf("读取到消息：%v\n", message)
//
//			}
//		}(conn)
//	}()
//	// 模拟客户端
//
//	conn, err := net.Dial("tcp", "127.0.0.1:8899")
//	if err != nil {
//		fmt.Printf("dial error:%v", err)
//	}
//
//	go func(conn net.Conn) {
//
//		defer conn.Close()
//
//		pack := Impl.DataPack{}
//		msg1 := &Impl.Message{
//			Id:      1,
//			DataLen: 5,
//			Data:    []byte("hello"),
//		}
//		bytes1, _ := pack.Pack(msg1)
//
//		msg2 := &Impl.Message{
//			Id:      2,
//			DataLen: 10,
//			Data:    []byte("jackismyna"),
//		}
//		bytes2, _ := pack.Pack(msg2)
//		msg3 := &Impl.Message{
//			Id:      3,
//			DataLen: 5,
//			Data:    []byte("hello"),
//		}
//		bytes3, _ := pack.Pack(msg3)
//		bytes := append(bytes1, bytes2...)
//		bytess := append(bytes, bytes3...)
//		_, err2 := conn.Write(bytess)
//		if err2 != nil {
//			return
//		}
//	}(conn)
//
//	select {}
//
//}
