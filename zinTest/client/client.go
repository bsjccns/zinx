package main

import (
	"fmt"
	"io"
	"net"
	"time"
	"zinx/zinx_server/Impl"
)

func main() {

	conn, err := net.Dial("tcp4", "0.0.0.0:10086")
	if err != nil {
		fmt.Println("连接失败")
		return
	}
	for {
		dp := &Impl.DataPack{}
		// 发送数据
		msg := &Impl.Message{
			Id:      1,
			DataLen: 10,
			Data:    []byte("hello zinx"),
		}
		bytes, err := dp.Pack(msg)
		if err != nil {
			fmt.Println("pack err[客户端]:", err)
		}
		_, err = conn.Write(bytes)
		if err != nil {
			fmt.Println("write err[客户端]:", err)
		}

		// 读取数据
		buffer := make([]byte, dp.GetPkgHeadLen())
		_, err = io.ReadFull(conn, buffer)
		if err != nil {
			fmt.Println("read full err[客户端]:", err)
			return
		}

		message, err := dp.UnPack(buffer)
		if err != nil {
			fmt.Println("unpack err[客户端]:", err)
		}

		message.SetData(make([]byte, message.GetMsgLen()))

		_, err = io.ReadFull(conn, message.GetData())

		if err != nil {
			fmt.Println("read full err[客户端]:", err)
		}

		fmt.Println("客户端收到消息:", string(message.GetData()), "消息id:", message.GetMsgId(), "消息长度:", message.GetMsgLen())

		time.Sleep(1 * time.Second)
	}
}
