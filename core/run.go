package core

import (
	"github.com/gin-gonic/gin"
	"github.com/jxcia/go-garden/core/log"
)

// Run service start
func (g *Garden) Run(route func(r *gin.Engine), rpc interface{}, auth func() gin.HandlerFunc) {
	go g.runHttpServer(route, auth)
	go g.runRpcServer(rpc)
	forever := make(chan int, 0)
	<-forever
}

func (g *Garden) runHttpServer(route func(r *gin.Engine), auth func() gin.HandlerFunc) {
	address := g.GetServiceIp()
	if g.cfg.Service.HttpOut {
		address = "0.0.0.0"
	}
	listenAddress := address + ":" + g.cfg.Service.HttpPort
	if err := g.ginListen(listenAddress, route, auth); err != nil {
		log.Fatal("gin", err)
	}
}

func (g *Garden) runRpcServer(rpc interface{}) {
	address := g.GetServiceIp()
	if g.cfg.Service.RpcOut {
		address = "0.0.0.0"
	}
	rpcAddress := address + ":" + g.cfg.Service.RpcPort
	if err := g.rpcListen(g.cfg.Service.ServiceName, "tcp", rpcAddress, rpc, ""); err != nil {
		log.Fatal("rpcRun", err)
	}
}
