package svc

import (
	"wpm-project/cmd/config"
	"wpm-project/cmd/database/mysql"
	"wpm-project/cmd/database/repo"
)

type ServiceContext struct {
	Config    config.Config
	UserRepo      repo.UserRepo
	PointRepo repo.PointRepo
	PostRepo repo.PostRepo
}

func NewServiceContext(c config.Config) *ServiceContext {
	userDB, err := mysql.NewUserDB(c)
	if err != nil {
		panic(err)
	}
	postDB, err := mysql.NewPostDB(c)
	if err != nil {
		panic(err)
	}
	pointDB, err := mysql.NewPointDB(c)
	if err != nil {
		panic(err)
	}
	
	return &ServiceContext{
		Config:    c,
		UserRepo: userDB,
		PostRepo: postDB,
		PointRepo: pointDB,
	}
}
