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
	table               string
	connect             *sqlx.DB
	ignoreInsertColumns []string
	datatimeColumns     []string
}

func NewUserDB(c config.Config) (repo.UserRepo, error) {
	db, err := sqlx.Open(c.Database.Driver, c.Database.Source)
	if err != nil {
		panic(err)
	}
	return &UserDB{
		table:               "users",
		connect:             db,
		ignoreInsertColumns: []string{"id"},
		datatimeColumns:     []string{},
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
		if condition.ID != 0 {
			db.Where(sq.Eq{"id": condition.ID})
		}
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
	curTime := time.Now().Format(time.RFC3339)
	user.Created_at = &curTime
	db := sq.Insert(u.table).
		Columns(GetListColumn(user, u.ignoreInsertColumns, u.datatimeColumns)...).
		Values(GetListValues(user, u.ignoreInsertColumns, u.datatimeColumns)...)
	query, arg, err := db.ToSql()
	if err != nil {
		ctxLogger.Errorf("Failed while build query, error: %s", err.Error())
		return err
	}
	ctxLogger.Debug(query)
	_, err = u.connect.Exec(query, arg...)
	if err != nil {
		ctxLogger.Errorf("Failed while create user, error: %s", err.Error())
		return err
	}
	return nil
}

func (u *UserDB) UpdateUser(ctx context.Context, user *model.User) error {
	ctxLogger := logger.NewContextLog(ctx)

	curTime := time.Now().Format(time.RFC3339)

	db := sq.Update(u.table).
		Set("name", user.Name).
		Set("password", user.Password).
		Set("mail", user.Mail).
		Set("updated_at", &curTime).
		Where(sq.Eq{"id": user.ID})

	query, arg, err := db.ToSql()
	if err != nil {
		ctxLogger.Errorf("Failed while build query, error: %s", err.Error())
		return err
	}
	ctxLogger.Debug(query)
	_, err = u.connect.Exec(query, arg...)
	if err != nil {
		ctxLogger.Errorf("Failed while update user, error: %s", err.Error())
		return err
	}
	return nil
}
