package Impl

import (
	"fmt"
	"io"
	"net"
	"testing"
)

func TestDataPack_Pack(t *testing.T) {
	listen, err := net.Listen("tcp", "127.0.0.1:8899")
	if err != nil {
		t.Fatalf("listen error:%v", err)
	}

	// 服务端
	go func() {
		for { // 这里的for
			conn, err := listen.Accept()
			if err != nil {
				t.Fatalf("accept error:%v", err)
			}
			go func(conn net.Conn) {
				for {

					// 拆包工具
					dataPack := &DataPack{}
					bytes := make([]byte, dataPack.GetPkgHeadLen())
					l, err := io.ReadFull(conn, bytes)
					if err != nil {
						t.Fatalf("read head error:%v", err)
					}
					if l == 0 {
						return
					}
					// 解出 消息长度 和消息id 存入message对象
					message, err := dataPack.UnPack(bytes)
					if err != nil {
						t.Fatalf("unpack error:%v", err)
					}
					//msg := message.(*Message)
					//再次根据消息长度读取数据
					if message.GetMsgLen() > 0 {

						message.SetData(make([]byte, message.GetMsgLen()))
						_, err = io.ReadFull(conn, message.GetData())
						if err != nil {
							t.Fatalf("read data error:%v", err)
						}
					}

					fmt.Printf("读取到消息：%v\n", string(message.GetData()))

				}
			}(conn)
		}
	}()
	// 模拟客户端

	conn, err := net.Dial("tcp", "127.0.0.1:8899")
	if err != nil {
		t.Fatalf("dial error:%v", err)
	}

	pack := DataPack{}
	msg1 := &Message{
		Id:      1,
		DataLen: 5,
		Data:    []byte("hello"),
	}
	bytes1, _ := pack.Pack(msg1)

	msg2 := &Message{
		Id:      2,
		DataLen: 10,
		Data:    []byte("jackismyna"),
	}
	bytes2, _ := pack.Pack(msg2)
	msg3 := &Message{
		Id:      3,
		DataLen: 5,
		Data:    []byte("hello"),
	}
	bytes3, _ := pack.Pack(msg3)
	bytes := append(bytes1, bytes2...)
	bytess := append(bytes, bytes3...)
	_, err2 := conn.Write(bytess)
	if err2 != nil {
		return
	}

	select {}
}
