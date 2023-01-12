package handler

import (
	"wpm-project/cmd/handler/point"
	"wpm-project/cmd/handler/user"
	"wpm-project/cmd/handler/post"
	"wpm-project/cmd/svc"

	"github.com/gin-gonic/gin"
)

func RegisterHandlers(router *gin.Engine, serviceCtx *svc.ServiceContext) {
	router.GET("api/users", user.GetUserHandler(serviceCtx))
	router.POST("api/users", user.CreateUserHandler(serviceCtx))
	router.PUT("api/users/:id", user.UpdateUserHandler(serviceCtx))
	router.GET("api/points", point.GetPointHandler(serviceCtx))
	router.POST("api/points", point.CreatePointHandler(serviceCtx))
	router.GET("api/posts", post.GetPostHandler(serviceCtx))
	router.POST("api/posts", post.CreatePostHandler(serviceCtx))
	router.PUT("api/posts/:id", post.UpdatePostHandler(serviceCtx))
	router.DELETE("api/posts/:id", post.DeletePostHandler(serviceCtx))
}