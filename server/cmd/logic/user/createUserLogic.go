package user

import (
	"context"
	"wpm-project/cmd/database/model"
	"wpm-project/cmd/svc"
	"wpm-project/cmd/types"

	"github.com/go-kratos/kratos/v2/log"
)

type CreateUserLogic struct {
	ctx       context.Context
	svcCtx    *svc.ServiceContext
	logHelper *log.Helper
}

func NewCreateUserLogic(ctx context.Context, svcCtx *svc.ServiceContext, logHelper *log.Helper) CreateUserLogic {
	return CreateUserLogic{
		ctx:       ctx,
		svcCtx:    svcCtx,
		logHelper: logHelper,
	}
}

func (l *GetUserLogic) CreateUser(user *types.CreateUserRequest) (*types.CreateUserResponse, error) {
	l.logHelper.Infof("Start process create user")
	err := l.svcCtx.UserRepo.CreateUser(l.ctx, &model.User{
		Name:     user.Name,
		Password: user.Password,
		Mail:     user.Mail,
	})
	if err != nil {
		l.logHelper.Errorf("Failed while get users, error: %s", err.Error())
		return nil, err
	}

	return &types.CreateUserResponse{}, nil
}
