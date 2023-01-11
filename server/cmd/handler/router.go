package handler

import (
	"wpm-project/cmd/svc"

	"github.com/gin-gonic/gin"
)

func RegisterHandlers(router *gin.Engine, serviceCtx *svc.ServiceContext) {
	router.GET("/users", GetUserHandler(serviceCtx))
}
