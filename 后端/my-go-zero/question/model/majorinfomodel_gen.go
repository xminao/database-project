// Code generated by goctl. DO NOT EDIT!

package model

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	majorInfoFieldNames          = builder.RawFieldNames(&MajorInfo{}, true)
	majorInfoRows                = strings.Join(majorInfoFieldNames, ",")
	majorInfoRowsExpectAutoSet   = strings.Join(stringx.Remove(majorInfoFieldNames, "create_time", "update_time"), ",")
	majorInfoRowsWithPlaceHolder = builder.PostgreSqlJoin(stringx.Remove(majorInfoFieldNames, "major_id", "create_time", "update_time"))

	cachePublicMajorInfoMajorIdPrefix = "cache:public:majorInfo:majorId:"
)

type (
	majorInfoModel interface {
		Insert(ctx context.Context, data *MajorInfo) (sql.Result, error)
		FindOne(ctx context.Context, majorId int64) (*MajorInfo, error)
		Update(ctx context.Context, data *MajorInfo) error
		Delete(ctx context.Context, majorId int64) error
	}

	defaultMajorInfoModel struct {
		sqlc.CachedConn
		table string
	}

	MajorInfo struct {
		MajorId   int64  `db:"major_id"`   // 专业编号
		MajorName string `db:"major_name"` // 专业名称
		CollegeId int64  `db:"college_id"` // 所属专业
	}
)

func newMajorInfoModel(conn sqlx.SqlConn, c cache.CacheConf) *defaultMajorInfoModel {
	return &defaultMajorInfoModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      `"public"."major_info"`,
	}
}

func (m *defaultMajorInfoModel) Insert(ctx context.Context, data *MajorInfo) (sql.Result, error) {
	publicMajorInfoMajorIdKey := fmt.Sprintf("%s%v", cachePublicMajorInfoMajorIdPrefix, data.MajorId)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values ($1, $2, $3)", m.table, majorInfoRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.MajorId, data.MajorName, data.CollegeId)
	}, publicMajorInfoMajorIdKey)
	return ret, err
}

func (m *defaultMajorInfoModel) FindOne(ctx context.Context, majorId int64) (*MajorInfo, error) {
	publicMajorInfoMajorIdKey := fmt.Sprintf("%s%v", cachePublicMajorInfoMajorIdPrefix, majorId)
	var resp MajorInfo
	err := m.QueryRowCtx(ctx, &resp, publicMajorInfoMajorIdKey, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where major_id = $1 limit 1", majorInfoRows, m.table)
		return conn.QueryRowCtx(ctx, v, query, majorId)
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

func (m *defaultMajorInfoModel) Update(ctx context.Context, data *MajorInfo) error {
	publicMajorInfoMajorIdKey := fmt.Sprintf("%s%v", cachePublicMajorInfoMajorIdPrefix, data.MajorId)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where major_id = $1", m.table, majorInfoRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, data.MajorId, data.MajorName, data.CollegeId)
	}, publicMajorInfoMajorIdKey)
	return err
}

func (m *defaultMajorInfoModel) Delete(ctx context.Context, majorId int64) error {
	publicMajorInfoMajorIdKey := fmt.Sprintf("%s%v", cachePublicMajorInfoMajorIdPrefix, majorId)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where major_id = $1", m.table)
		return conn.ExecCtx(ctx, query, majorId)
	}, publicMajorInfoMajorIdKey)
	return err
}

func (m *defaultMajorInfoModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cachePublicMajorInfoMajorIdPrefix, primary)
}

func (m *defaultMajorInfoModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where major_id = $1 limit 1", majorInfoRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultMajorInfoModel) tableName() string {
	return m.table
}
