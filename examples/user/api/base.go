package api

import (
	"github.com/gin-gonic/gin"
	"github.com/panco95/go-garden/examples/user/global"
)

func Routes(r *gin.Engine) {
	r.Use(global.Service.CheckCallSafeMiddleware()) // 调用接口安全验证
	r.POST("login", Login)
}
