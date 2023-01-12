package post

import (
	"context"
	"wpm-project/cmd/database/model"
	"wpm-project/cmd/svc"
	"wpm-project/cmd/types"

	"github.com/go-kratos/kratos/v2/log"
)

type UpdatePostLogic struct {
	ctx       context.Context
	svcCtx    *svc.ServiceContext
	logHelper *log.Helper
}

func NewUpdatePostLogic(ctx context.Context, svcCtx *svc.ServiceContext, logHelper *log.Helper) UpdatePostLogic {
	return UpdatePostLogic{
		ctx:       ctx,
		svcCtx:    svcCtx,
		logHelper: logHelper,
	}
}

func (l *UpdatePostLogic) UpdatePost(path *types.UpdatePostRequest, body *types.UpdatePostBody) (*types.UpdatePostResponse, error) {
	l.logHelper.Infof("Start process update Post")
	err := l.svcCtx.PostRepo.UpdatePost(l.ctx, &model.Post{
		ID:           path.ID,
		Title:        body.Title,
		Descriptions: body.Descriptions,
		Content:      body.Content,
	})
	if err != nil {
		l.logHelper.Errorf("Failed while update Post, error: %s", err.Error())
		return nil, err
	}

	return &types.UpdatePostResponse{}, nil
}
