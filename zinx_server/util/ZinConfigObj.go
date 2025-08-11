package util

import (
	"encoding/json"
	"fmt"
	"os"
)

type ZinConfigObj struct {
	Ip         string `json:"Ip"`
	Port       uint32 `json:"Port"`
	Name       string `json:"Name"`
	MaxConn    uint32 `json:"MaxConn"`
	MaxPkgSize uint32 `json:"MaxPkgSize"`

	WorkerPoolSize uint32 `json:"WorkerPoolSize"`
	TaskQueueSize  uint32 `json:"TaskQueueSize"`
}

var Config1 *ZinConfigObj

func (C *ZinConfigObj) AnalysisConfig() {
	// 尝试从环境变量读取配置文件路径
	configPath := os.Getenv("ZINX_CONFIG_PATH")
	if configPath == "" {
		// 环境变量未设置，打印错误并退出
		fmt.Println("环境变量 ZINX_CONFIG_PATH 未设置，无法找到配置文件。")
		panic("配置文件路径未指定")
	}

	// 读取配置文件
	data, err := os.ReadFile(configPath)
	if err != nil {
		fmt.Println("失败：", err)
		panic(err)
	}

	err = json.Unmarshal(data, C)

	if err != nil {
		panic(err)
	}
}

func init() {
	C := &ZinConfigObj{
		Name:           "zinx",
		Ip:             "0.0.0.0",
		Port:           9999,
		MaxConn:        1000,
		MaxPkgSize:     4096,
		WorkerPoolSize: 0,
		TaskQueueSize:  0,
	}

	//加载自定义参数
	//C.AnalysisConfig()
	Config1 = C
}
