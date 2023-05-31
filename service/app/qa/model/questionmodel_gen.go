// Code generated by goctl. DO NOT EDIT.

package model

import (
	"context"
	"database/sql"
	"fmt"
	"gorm.io/gorm"
	"strings"
	"time"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	questionFieldNames          = builder.RawFieldNames(&Question{})
	questionRows                = strings.Join(questionFieldNames, ",")
	questionRowsExpectAutoSet   = strings.Join(stringx.Remove(questionFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	questionRowsWithPlaceHolder = strings.Join(stringx.Remove(questionFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"

	cacheQuestionIdPrefix = "cache:question:id:"
)

type (
	questionModel interface {
		Insert(ctx context.Context, data *Question) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*Question, error)
		Update(ctx context.Context, data *Question) error
		Delete(ctx context.Context, id int64) error
	}

	defaultQuestionModel struct {
		sqlc.CachedConn
		table string
	}

	Question struct {
		Id         int64          `db:"id"`
		UserId     int64          `db:"user_id"`
		Title      string         `db:"title"`
		Content    string         `db:"content"`
		CreateTime time.Time      `db:"create_time"`              // 创建时间
		UpdateTime time.Time      `db:"update_time"`              // 更新时间
		DeleteTime gorm.DeletedAt `gorm:"index" db:"delete_time"` // 删除时间
	}
)

func newQuestionModel(conn sqlx.SqlConn, c cache.CacheConf) *defaultQuestionModel {
	return &defaultQuestionModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`question`",
	}
}

func (m *defaultQuestionModel) Delete(ctx context.Context, id int64) error {
	questionIdKey := fmt.Sprintf("%s%v", cacheQuestionIdPrefix, id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.ExecCtx(ctx, query, id)
	}, questionIdKey)
	return err
}

func (m *defaultQuestionModel) FindOne(ctx context.Context, id int64) (*Question, error) {
	questionIdKey := fmt.Sprintf("%s%v", cacheQuestionIdPrefix, id)
	var resp Question
	err := m.QueryRowCtx(ctx, &resp, questionIdKey, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", questionRows, m.table)
		return conn.QueryRowCtx(ctx, v, query, id)
	})
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultQuestionModel) Insert(ctx context.Context, data *Question) (sql.Result, error) {
	questionIdKey := fmt.Sprintf("%s%v", cacheQuestionIdPrefix, data.Id)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?)", m.table, questionRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.UserId, data.Title, data.Content, data.DeleteTime)
	}, questionIdKey)
	return ret, err
}

func (m *defaultQuestionModel) Update(ctx context.Context, data *Question) error {
	questionIdKey := fmt.Sprintf("%s%v", cacheQuestionIdPrefix, data.Id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, questionRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, data.UserId, data.Title, data.Content, data.DeleteTime, data.Id)
	}, questionIdKey)
	return err
}

func (m *defaultQuestionModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheQuestionIdPrefix, primary)
}

func (m *defaultQuestionModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", questionRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultQuestionModel) tableName() string {
	return m.table
}
