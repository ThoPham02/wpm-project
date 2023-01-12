package handler

import (
	"wpm-project/cmd/svc"
	"wpm-project/cmd/handler/user"

	"github.com/gin-gonic/gin"
)

func RegisterHandlers(router *gin.Engine, serviceCtx *svc.ServiceContext) {
	router.GET("api/users", user.GetUserHandler(serviceCtx))
	router.POST("api/users", user.CreateUserHandler(serviceCtx))
	router.PUT("api/users/:id", user.UpdateUserHandler(serviceCtx))
}