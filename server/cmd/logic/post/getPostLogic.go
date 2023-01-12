package post

import (
	"context"
	"database/sql"
	"wpm-project/cmd/svc"
	"wpm-project/cmd/types"

	"github.com/go-kratos/kratos/v2/log"
)

type GetPostLogic struct {
	ctx       context.Context
	svcCtx    *svc.ServiceContext
	logHelper *log.Helper
}

func NewGetPostLogic(ctx context.Context, svcCtx *svc.ServiceContext, logHelper *log.Helper) GetPostLogic {
	return GetPostLogic{
		ctx:       ctx,
		svcCtx:    svcCtx,
		logHelper: logHelper,
	}
}

func (l *GetPostLogic) GetPost() (*types.GetPostResponse, error) {
	l.logHelper.Infof("Start process get posts")
	posts, err := l.svcCtx.PostRepo.GetPost(l.ctx, nil)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		l.logHelper.Errorf("Failed while get posts, error: %s", err.Error())
		return nil, err
	}

	var dataResponse []types.Post
	for _, value := range posts {
		dataResponse = append(dataResponse, types.Post{
			ID: value.ID,
			Title: value.Title,
			Descriptions: value.Descriptions,
			Content: value.Content,
			UserID: value.UserID,
			CreatedAt: value.Created_at,
		})
	}

	return &types.GetPostResponse{
		Data: dataResponse,
	}, nil
}
