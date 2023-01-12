package repo

import (
	"context"
	"wpm-project/cmd/database/model"
)

type PointConditions struct {
	ID     int64
	Search string
}

type PointRepo interface {
	GetPoint(context.Context, *PointConditions) ([]*model.Point, error)
	CreatePoint(context.Context, *model.Point) error
}
