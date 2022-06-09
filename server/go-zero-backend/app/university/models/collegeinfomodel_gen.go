// Code generated by goctl. DO NOT EDIT!

package models

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	collegeInfoFieldNames          = builder.RawFieldNames(&CollegeInfo{}, true)
	collegeInfoRows                = strings.Join(collegeInfoFieldNames, ",")
	collegeInfoRowsExpectAutoSet   = strings.Join(stringx.Remove(collegeInfoFieldNames, "create_time", "update_time"), ",")
	collegeInfoRowsWithPlaceHolder = builder.PostgreSqlJoin(stringx.Remove(collegeInfoFieldNames, "college_id", "create_time", "update_time"))

	cachePublicCollegeInfoCollegeIdPrefix   = "cache:public:collegeInfo:collegeId:"
	cachePublicCollegeInfoCollegeNamePrefix = "cache:public:collegeInfo:collegeName:"
)

type (
	collegeInfoModel interface {
		Insert(ctx context.Context, data *CollegeInfo) (sql.Result, error)
		FindOne(ctx context.Context, collegeId int64) (*CollegeInfo, error)
		FindOneByCollegeName(ctx context.Context, collegeName string) (*CollegeInfo, error)
		Update(ctx context.Context, data *CollegeInfo) error
		Delete(ctx context.Context, collegeId int64) error
	}

	defaultCollegeInfoModel struct {
		sqlc.CachedConn
		table string
	}

	CollegeInfo struct {
		CollegeId   int64     `db:"college_id"`   // 学院编号
		CollegeName string    `db:"college_name"` // 学院名称
		CreatedAt   time.Time `db:"created_at"`   // 创建时间
	}
)

func newCollegeInfoModel(conn sqlx.SqlConn, c cache.CacheConf) *defaultCollegeInfoModel {
	return &defaultCollegeInfoModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      `"public"."college_info"`,
	}
}

func (m *defaultCollegeInfoModel) Insert(ctx context.Context, data *CollegeInfo) (sql.Result, error) {
	publicCollegeInfoCollegeIdKey := fmt.Sprintf("%s%v", cachePublicCollegeInfoCollegeIdPrefix, data.CollegeId)
	publicCollegeInfoCollegeNameKey := fmt.Sprintf("%s%v", cachePublicCollegeInfoCollegeNamePrefix, data.CollegeName)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values ($1, $2, $3)", m.table, collegeInfoRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.CollegeId, data.CollegeName, data.CreatedAt)
	}, publicCollegeInfoCollegeIdKey, publicCollegeInfoCollegeNameKey)
	return ret, err
}

func (m *defaultCollegeInfoModel) FindOne(ctx context.Context, collegeId int64) (*CollegeInfo, error) {
	publicCollegeInfoCollegeIdKey := fmt.Sprintf("%s%v", cachePublicCollegeInfoCollegeIdPrefix, collegeId)
	var resp CollegeInfo
	err := m.QueryRowCtx(ctx, &resp, publicCollegeInfoCollegeIdKey, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where college_id = $1 limit 1", collegeInfoRows, m.table)
		return conn.QueryRowCtx(ctx, v, query, collegeId)
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

func (m *defaultCollegeInfoModel) FindOneByCollegeName(ctx context.Context, collegeName string) (*CollegeInfo, error) {
	publicCollegeInfoCollegeNameKey := fmt.Sprintf("%s%v", cachePublicCollegeInfoCollegeNamePrefix, collegeName)
	var resp CollegeInfo
	err := m.QueryRowIndexCtx(ctx, &resp, publicCollegeInfoCollegeNameKey, m.formatPrimary, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) (i interface{}, e error) {
		query := fmt.Sprintf("select %s from %s where college_name = $1 limit 1", collegeInfoRows, m.table)
		if err := conn.QueryRowCtx(ctx, &resp, query, collegeName); err != nil {
			return nil, err
		}
		return resp.CollegeId, nil
	}, m.queryPrimary)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultCollegeInfoModel) Update(ctx context.Context, data *CollegeInfo) error {
	publicCollegeInfoCollegeIdKey := fmt.Sprintf("%s%v", cachePublicCollegeInfoCollegeIdPrefix, data.CollegeId)
	publicCollegeInfoCollegeNameKey := fmt.Sprintf("%s%v", cachePublicCollegeInfoCollegeNamePrefix, data.CollegeName)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where college_id = $1", m.table, collegeInfoRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, data.CollegeId, data.CollegeName, data.CreatedAt)
	}, publicCollegeInfoCollegeIdKey, publicCollegeInfoCollegeNameKey)
	return err
}

func (m *defaultCollegeInfoModel) Delete(ctx context.Context, collegeId int64) error {
	data, err := m.FindOne(ctx, collegeId)
	if err != nil {
		return err
	}

	publicCollegeInfoCollegeIdKey := fmt.Sprintf("%s%v", cachePublicCollegeInfoCollegeIdPrefix, collegeId)
	publicCollegeInfoCollegeNameKey := fmt.Sprintf("%s%v", cachePublicCollegeInfoCollegeNamePrefix, data.CollegeName)
	_, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where college_id = $1", m.table)
		return conn.ExecCtx(ctx, query, collegeId)
	}, publicCollegeInfoCollegeIdKey, publicCollegeInfoCollegeNameKey)
	return err
}

func (m *defaultCollegeInfoModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cachePublicCollegeInfoCollegeIdPrefix, primary)
}

func (m *defaultCollegeInfoModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where college_id = $1 limit 1", collegeInfoRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultCollegeInfoModel) tableName() string {
	return m.table
}
