package Zinterface

type ZRouter interface {
	PreHandle(r ZRequest)
	Handle(r ZRequest)
	PostHandle(r ZRequest)
}
