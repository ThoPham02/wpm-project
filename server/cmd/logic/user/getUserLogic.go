package user

import (
	"context"
	"database/sql"
	"wpm-project/cmd/svc"
	"wpm-project/cmd/types"

	"github.com/go-kratos/kratos/v2/log"
)

type GetUserLogic struct {
	ctx       context.Context
	svcCtx    *svc.ServiceContext
	logHelper *log.Helper
}

func NewGetUserLogic(ctx context.Context, svcCtx *svc.ServiceContext, logHelper *log.Helper) GetUserLogic {
	return GetUserLogic{
		ctx:       ctx,
		svcCtx:    svcCtx,
		logHelper: logHelper,
	}
}

func (l *GetUserLogic) GetUser() (*types.GetUserResponse, error) {
	l.logHelper.Infof("Start process get user")
	users, err := l.svcCtx.UserRepo.GetUser(l.ctx, nil)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		l.logHelper.Errorf("Failed while get users, error: %s", err.Error())
		return nil, err
	}

	var dataResponse []types.User
	for _, value := range users {
		dataResponse = append(dataResponse, types.User{
			ID: value.ID,
			Name: value.Name,
			Mail: value.Mail,
			CreatedAt: value.Created_at,
		})
	}

	return &types.GetUserResponse{
		Data: dataResponse,
	}, nil
}
