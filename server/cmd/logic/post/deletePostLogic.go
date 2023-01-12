package post

import (
	"context"
	"wpm-project/cmd/database/repo"
	"wpm-project/cmd/svc"
	"wpm-project/cmd/types"

	"github.com/go-kratos/kratos/v2/log"
)

type DeletePostLogic struct {
	ctx       context.Context
	svcCtx    *svc.ServiceContext
	logHelper *log.Helper
}

func NewDeletePostLogic(ctx context.Context, svcCtx *svc.ServiceContext, logHelper *log.Helper) DeletePostLogic {
	return DeletePostLogic{
		ctx:       ctx,
		svcCtx:    svcCtx,
		logHelper: logHelper,
	}
}

func (l *DeletePostLogic) DeletePost(request *types.DeletePostRequest) (*types.DeletePostResponse, error) {
	l.logHelper.Infof("Start process Delete Post")
	err := l.svcCtx.PostRepo.DeletePost(l.ctx, &repo.PostConditions{
		ID: request.ID,
	})
	if err != nil {
		l.logHelper.Errorf("Failed while Delete Post, error: %s", err.Error())
		return nil, err
	}

	return &types.DeletePostResponse{}, nil
}