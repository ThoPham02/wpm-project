package repo

import (
	"context"
	"wpm-project/cmd/database/model"
)

type PostConditions struct {

}

type PostRepo interface {
	GetPost(context.Context, *PostConditions) ([]*model.Post, error)
	CreatePost(context.Context, *model.Post) error
	UpdatePost(context.Context, *model.Post) error
	DeletePost(context.Context, *PostConditions) error
}