package repo

import (
	"context"
	"wpm-project/cmd/database/model"
)

type UserConditions struct {
	ID int64
}

type UserRepo interface {
	GetUser(context.Context, *UserConditions) ([]*model.User, error)
	CreateUser(context.Context, *model.User) error
	UpdateUser(context.Context, *model.User) error
}