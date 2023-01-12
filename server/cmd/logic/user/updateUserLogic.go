package user

import (
	"context"
	"wpm-project/cmd/database/model"
	"wpm-project/cmd/svc"
	"wpm-project/cmd/types"

	"github.com/go-kratos/kratos/v2/log"
)

type UpdateUserLogic struct {
	ctx       context.Context
	svcCtx    *svc.ServiceContext
	logHelper *log.Helper
}

func NewUpdateUserLogic(ctx context.Context, svcCtx *svc.ServiceContext, logHelper *log.Helper) UpdateUserLogic {
	return UpdateUserLogic{
		ctx:       ctx,
		svcCtx:    svcCtx,
		logHelper: logHelper,
	}
}

func (l *UpdateUserLogic) UpdateUser(path *types.UpdateUserRequest, body *types.UpdateUserBody) (*types.UpdateUserResponse, error) {
	l.logHelper.Infof("Start process update user")
	err := l.svcCtx.UserRepo.UpdateUser(l.ctx, &model.User{
		ID:         path.ID,
		Name:       body.Name,
		Password:   body.Password,
		Mail:       body.Mail,
	})
	if err != nil {
		l.logHelper.Errorf("Failed while update user, error: %s", err.Error())
		return nil, err
	}

	return &types.UpdateUserResponse{}, nil
}
