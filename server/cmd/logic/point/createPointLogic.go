package point

import (
	"context"
	"wpm-project/cmd/database/model"
	"wpm-project/cmd/svc"
	"wpm-project/cmd/types"

	"github.com/go-kratos/kratos/v2/log"
)

type CreatePointLogic struct {
	ctx       context.Context
	svcCtx    *svc.ServiceContext
	logHelper *log.Helper
}

func NewCreatePointLogic(ctx context.Context, svcCtx *svc.ServiceContext, logHelper *log.Helper) CreatePointLogic {
	return CreatePointLogic{
		ctx:       ctx,
		svcCtx:    svcCtx,
		logHelper: logHelper,
	}
}

func (l *CreatePointLogic) CreatePoint(point *types.CreatePointRequest) (*types.CreatePointResponse, error) {
	l.logHelper.Infof("Start process create point")
	err := l.svcCtx.PointRepo.CreatePoint(l.ctx, &model.Point{
		UserID: point.UserID,
		PostID: point.PostID,
		Wpm: point.Wpm,
	})
	if err != nil {
		l.logHelper.Errorf("Failed while create points, error: %s", err.Error())
		return nil, err
	}

	return &types.CreatePointResponse{}, nil
}