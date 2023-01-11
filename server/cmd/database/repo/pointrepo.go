package repo

import (
	"context"
	"wpm-project/cmd/database/model"
)

type PointConditions struct {
}

type PointRepo interface {
	GetPoint(context.Context, *PointConditions) ([]*model.Point, error)
	CreatePoint(context.Context, *model.Point) error
}