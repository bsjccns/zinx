package Impl

import (
	"bytes"
	"encoding/binary"
	"errors"
	"zinx/zinx_server/Zinterface"
	"zinx/zinx_server/util"
)

type DataPack struct {
}

func (Dp *DataPack) GetPkgHeadLen() uint32 {
	return 8 // 8字节的头信息
}
func (Dp *DataPack) Pack(message Zinterface.IMessage) ([]byte, error) {
	buffer := bytes.NewBuffer([]byte{})
	// 写入消息长度
	err := binary.Write(buffer, binary.LittleEndian, message.GetMsgLen())
	if err != nil {
		return nil, err
	}
	// 写入消息id
	err = binary.Write(buffer, binary.LittleEndian, message.GetMsgId())
	if err != nil {
		return nil, err
	}
	//写入消息数据
	err = binary.Write(buffer, binary.LittleEndian, message.GetData())
	if err != nil {
		return nil, err
	}

	return buffer.Bytes(), nil

}
func (Dp *DataPack) UnPack(bytesData []byte) (Zinterface.IMessage, error) {
	// io reader
	reader := bytes.NewReader(bytesData)
	//创建msg
	msg := &Message{}
	// 读入消息长度  需要传指针&msg.DataLen 因为要读取到msg的DataLen字段中去
	err := binary.Read(reader, binary.LittleEndian, &msg.DataLen)
	if err != nil {
		return nil, err
	}
	err = binary.Read(reader, binary.LittleEndian, &msg.Id)
	if err != nil {
		return nil, err
	}

	if util.Config1.MaxPkgSize > 0 && msg.DataLen > util.Config1.MaxPkgSize {
		return nil, errors.New("too large msg")
	}

	return msg, nil
}
