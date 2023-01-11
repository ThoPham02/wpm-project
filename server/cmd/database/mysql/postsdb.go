package mysql

import (
	"context"
	"database/sql"
	"time"
	"wpm-project/cmd/config"
	"wpm-project/cmd/database/model"
	"wpm-project/cmd/database/repo"
	"wpm-project/core/logger"

	sq "github.com/Masterminds/squirrel"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type PostDB struct {
	table string
	connect *sqlx.DB
	ignoreInsertColumns []string
	datatimeColumns []string
}

func NewPostDB(c config.Config) (repo.PostRepo, error){ 
	db, err := sqlx.Open(c.Database.Driver, c.Database.Source)
	if err != nil {
		panic(err)
	}
	return &PostDB{
		table: "posts",
		connect: db,
		ignoreInsertColumns: []string{"id"},
		datatimeColumns: []string{"created_at", "updated_at", "deleted_at"},
	}, nil
}

func (post *PostDB) Close() {
	err := post.connect.Close()
	if err != nil {
		return
	}
}

func (u *PostDB) GetPost(ctx context.Context, condition *repo.PostConditions) ([]*model.Post, error) {
	ctxLogger := logger.NewContextLog(ctx)
	db := sq.Select("*").From(u.table)
	if condition != nil {

	}
	query, arg, err := db.ToSql()
	if err != nil {
		ctxLogger.Errorf("Failed while build query, error: %s", err.Error())
		return nil, err
	}
	var result []*model.Post
	err = u.connect.Select(&result, query, arg...)
	if err != nil {
		ctxLogger.Errorf("Failed while select user, error: %s", err.Error())
		return nil, err
	}
	if result == nil {
		return nil, sql.ErrNoRows
	}
	return result, nil
}

func (u *PostDB) CreatePost(ctx context.Context, post *model.Post) error {
	ctxLogger := logger.NewContextLog(ctx)
	db := sq.Insert(u.table).
	Columns(GetListColumn(post, u.ignoreInsertColumns, u.datatimeColumns)...).
	Values(GetListValues(post, u.ignoreInsertColumns, u.datatimeColumns)...).
	Columns("created_at").Values(time.Now())

	query, arg, err := db.ToSql()
	if err != nil {
		ctxLogger.Errorf("Failed while build query, error: %s", err.Error())
		return err
	}
	_, err = u.connect.Exec(query, arg...)
	if err != nil {
		ctxLogger.Errorf("Failed while create post, error: %s", err.Error())
		return err
	}
	return nil
}

func (u *PostDB) UpdatePost(ctx context.Context,post *model.Post) error {
	return nil
}

func (u *PostDB) DeletePost(ctx context.Context, condition *repo.PostConditions) error {
	return nil
}