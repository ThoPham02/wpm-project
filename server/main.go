package main

import (
	"fmt"
	"wpm-project/cmd/config"
	"wpm-project/cmd/handler"
	"wpm-project/cmd/svc"

	"github.com/gin-gonic/gin"
)


func main() {
	c := &config.Config{
		Http: config.HttpServer{
			Path: "localhost",
			Port: "8000",
		},
		Database: config.Database{
			Driver: "mysql",
			Source: "admin:secret@tcp(localhost:3306)/wpm",
		},
	}

	serverCtx := svc.NewServiceContext(*c)
	router := gin.Default()
	handler.RegisterHandlers(router, serverCtx)
	fmt.Printf("Starting server at %s:%s...\n", c.Http.Path, c.Http.Port)
	err := router.Run(fmt.Sprintf("%s:%s", c.Http.Path, c.Http.Port))
	if err != nil {
		panic(err)
	}
}