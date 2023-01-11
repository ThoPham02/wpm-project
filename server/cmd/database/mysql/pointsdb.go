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

type PointDB struct {
	table string
	connect *sqlx.DB
	ignoreInsertColumns []string
	datatimeColumns []string
}

func NewPointDB(c config.Config) (repo.PointRepo, error){ 
	db, err := sqlx.Open(c.Database.Driver, c.Database.Source)
	if err != nil {
		panic(err)
	}
	return &PointDB{
		table: "points",
		connect: db,
		ignoreInsertColumns: []string{"id"},
		datatimeColumns: []string{"created_at", "updated_at", "deleted_at"},
	}, nil
}

func (point *PointDB) Close() {
	err := point.connect.Close()
	if err != nil {
		return
	}
}

func (u *PointDB) GetPoint(ctx context.Context, condition *repo.PointConditions) ([]*model.Point, error) {
	ctxLogger := logger.NewContextLog(ctx)
	db := sq.Select("*").From(u.table)
	if condition != nil {

	}
	query, arg, err := db.ToSql()
	if err != nil {
		ctxLogger.Errorf("Failed while build query, error: %s", err.Error())
		return nil, err
	}
	var result []*model.Point
	err = u.connect.Select(&result, query, arg...)
	if err != nil {
		ctxLogger.Errorf("Failed while select point, error: %s", err.Error())
		return nil, err
	}
	if result == nil {
		return nil, sql.ErrNoRows
	}
	return result, nil
}

func (u *PointDB) CreatePoint(ctx context.Context, point *model.Point) error {
	ctxLogger := logger.NewContextLog(ctx)
	db := sq.Insert(u.table).
	Columns(GetListColumn(point, u.ignoreInsertColumns, u.datatimeColumns)...).
	Values(GetListValues(point, u.ignoreInsertColumns, u.datatimeColumns)...).
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