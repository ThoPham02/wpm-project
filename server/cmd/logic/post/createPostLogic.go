package post

import (
	"context"
	"wpm-project/cmd/database/model"
	"wpm-project/cmd/svc"
	"wpm-project/cmd/types"

	"github.com/go-kratos/kratos/v2/log"
)

type CreatePostLogic struct {
	ctx       context.Context
	svcCtx    *svc.ServiceContext
	logHelper *log.Helper
}

func NewCreatePostLogic(ctx context.Context, svcCtx *svc.ServiceContext, logHelper *log.Helper) CreatePostLogic {
	return CreatePostLogic{
		ctx:       ctx,
		svcCtx:    svcCtx,
		logHelper: logHelper,
	}
}

func (l *CreatePostLogic) CreatePost(post *types.CreatePostRequest) (*types.CreatePostResponse, error) {
	l.logHelper.Infof("Start process create post")
	err := l.svcCtx.PostRepo.CreatePost(l.ctx, &model.Post{
		Title: post.Title,
		Descriptions: post.Descriptions,
		Content: post.Content,
		UserID: post.UserID,
	})
	if err != nil {
		l.logHelper.Errorf("Failed while get users, error: %s", err.Error())
		return nil, err
	}

	return &types.CreatePostResponse{}, nil
}
