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

type UserDB struct {
	table string
	connect *sqlx.DB
	ignoreInsertColumns []string
	datatimeColumns []string
}

func NewUserDB(c config.Config) (repo.UserRepo, error){ 
	db, err := sqlx.Open(c.Database.Driver, c.Database.Source)
	if err != nil {
		panic(err)
	}
	return &UserDB{
		table: "users",
		connect: db,
		ignoreInsertColumns: []string{"id"},
		datatimeColumns: []string{"created_at", "updated_at", "deleted_at"},
	}, nil
}

func (user *UserDB) Close() {
	err := user.connect.Close()
	if err != nil {
		return
	}
}

func (u *UserDB) GetUser(ctx context.Context, condition *repo.UserConditions) ([]*model.User, error) {
	ctxLogger := logger.NewContextLog(ctx)
	db := sq.Select("*").From(u.table)
	if condition != nil {

	}
	query, arg, err := db.ToSql()
	if err != nil {
		ctxLogger.Errorf("Failed while build query, error: %s", err.Error())
		return nil, err
	}
	var result []*model.User
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

func (u *UserDB) CreateUser(ctx context.Context, user *model.User) error {
	ctxLogger := logger.NewContextLog(ctx)
	db := sq.Insert(u.table).
	Columns(GetListColumn(user, u.ignoreInsertColumns, u.datatimeColumns)...).
	Values(GetListValues(user, u.ignoreInsertColumns, u.datatimeColumns)...).
	Columns("created_at").Values(time.Now())

	query, arg, err := db.ToSql()
	if err != nil {
		ctxLogger.Errorf("Failed while build query, error: %s", err.Error())
		return err
	}
	_, err = u.connect.Exec(query, arg...)
	if err != nil {
		ctxLogger.Errorf("Failed while create user, error: %s", err.Error())
		return err
	}
	return nil
}

func (u *UserDB) UpdateUser(ctx context.Context,user *model.User) error {
	return nil
}

func (u *UserDB) DeleteUser(ctx context.Context, condition *repo.UserConditions) error {
	return nil
}