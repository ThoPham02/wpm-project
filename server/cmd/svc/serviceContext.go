package svc

import (
	"wpm-project/cmd/config"
	"wpm-project/cmd/database/mysql"
	"wpm-project/cmd/database/repo"
)

type ServiceContext struct {
	Config    config.Config
	UserRepo      repo.UserRepo
}

func NewServiceContext(c config.Config) *ServiceContext {
	userDB, err := mysql.NewUserDB(c)
	if err != nil {
		panic(err)
	}
	
	return &ServiceContext{
		Config:    c,
		UserRepo: userDB,
	}
}
