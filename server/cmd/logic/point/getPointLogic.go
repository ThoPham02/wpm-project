package point

import (
	"context"
	"database/sql"
	"wpm-project/cmd/database/repo"
	"wpm-project/cmd/svc"
	"wpm-project/cmd/types"

	"github.com/go-kratos/kratos/v2/log"
)

type GetPointLogic struct {
	ctx       context.Context
	svcCtx    *svc.ServiceContext
	logHelper *log.Helper
}

func NewGetPointLogic(ctx context.Context, svcCtx *svc.ServiceContext, logHelper *log.Helper) GetPointLogic {
	return GetPointLogic{
		ctx:       ctx,
		svcCtx:    svcCtx,
		logHelper: logHelper,
	}
}

func (l *GetPointLogic) GetPoint() (*types.GetPointResponse, error) {
	l.logHelper.Infof("Start process get point")
	points, err := l.svcCtx.PointRepo.GetPoint(l.ctx, nil)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		l.logHelper.Errorf("Failed while get point, error: %s", err.Error())
		return nil, err
	}

	var dataResponse []types.Point
	for _, value := range points {
		users, err := l.svcCtx.UserRepo.GetUser(l.ctx, &repo.UserConditions{
			
		})
		if err != nil {
			l.logHelper.Errorf("Failed while get user name, error: %s", err.Error())
			return nil, err
		}
		user := users[0]
		dataResponse = append(dataResponse, types.Point{
			ID:         value.ID,
			Name:       user.Name,
			PostID:     value.PostID,
			Wpm:        value.Wpm,
			Created_at: value.Created_at,
		})
	}

	return &types.GetPointResponse{
		Data: dataResponse,
	}, nil
}
