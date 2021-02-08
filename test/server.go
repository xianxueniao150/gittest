package main

import (
	"fmt"
	"../znet"
	"../ziface"
)

//ping test 自定义路由
type PingRouter struct {
	znet.BaseRouter
}

//Ping Handle
func (this *PingRouter) PreHandle(request ziface.IRequest) {
	//_, _ = request.GetConnection().GetTCPConnection().Write([]byte("before"))
}

//Ping Handle
func (this *PingRouter) Handle(request ziface.IRequest) {
	fmt.Println("recv from client : msgId=", request.GetMsgID(), ", data=", string(request.GetData()))
	_ = request.GetConnection().SendMsg(0,[]byte("ping ping "))
}

type HelloRouter struct {
	znet.BaseRouter
}

//Hello Handle
func (this *HelloRouter) PostHandle(request ziface.IRequest) {
	fmt.Println("recv from client : msgId=", request.GetMsgID(), ", data=", string(request.GetData()))
	_ = request.GetConnection().SendMsg(1,[]byte("hello hello "))
}


func main() {
	s := znet.NewServer("v0.1")
	s.AddRouter(0,&PingRouter{})
	s.AddRouter(1,&HelloRouter{})
	s.Serve()
}
