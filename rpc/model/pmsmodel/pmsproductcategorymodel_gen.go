// Code generated by goctl. DO NOT EDIT.

package pmsmodel

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	pmsProductCategoryFieldNames          = builder.RawFieldNames(&PmsProductCategory{})
	pmsProductCategoryRows                = strings.Join(pmsProductCategoryFieldNames, ",")
	pmsProductCategoryRowsExpectAutoSet   = strings.Join(stringx.Remove(pmsProductCategoryFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	pmsProductCategoryRowsWithPlaceHolder = strings.Join(stringx.Remove(pmsProductCategoryFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"
)

type (
	pmsProductCategoryModel interface {
		Insert(ctx context.Context, data *PmsProductCategory) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*PmsProductCategory, error)
		Update(ctx context.Context, data *PmsProductCategory) error
		Delete(ctx context.Context, id int64) error
	}

	defaultPmsProductCategoryModel struct {
		conn  sqlx.SqlConn
		table string
	}

	PmsProductCategory struct {
		Id           int64          `db:"id"`
		ParentId     int64          `db:"parent_id"` // 上机分类的编号：0表示一级分类
		Name         string         `db:"name"`
		Level        int64          `db:"level"` // 分类级别：0->1级；1->2级
		ProductCount int64          `db:"product_count"`
		ProductUnit  string         `db:"product_unit"`
		NavStatus    int64          `db:"nav_status"`  // 是否显示在导航栏：0->不显示；1->显示
		ShowStatus   int64          `db:"show_status"` // 显示状态：0->不显示；1->显示
		Sort         int64          `db:"sort"`
		Icon         string         `db:"icon"` // 图标
		Keywords     string         `db:"keywords"`
		Description  sql.NullString `db:"description"` // 描述
	}
)

func newPmsProductCategoryModel(conn sqlx.SqlConn) *defaultPmsProductCategoryModel {
	return &defaultPmsProductCategoryModel{
		conn:  conn,
		table: "`pms_product_category`",
	}
}

func (m *defaultPmsProductCategoryModel) Delete(ctx context.Context, id int64) error {
	query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, id)
	return err
}

func (m *defaultPmsProductCategoryModel) FindOne(ctx context.Context, id int64) (*PmsProductCategory, error) {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", pmsProductCategoryRows, m.table)
	var resp PmsProductCategory
	err := m.conn.QueryRowCtx(ctx, &resp, query, id)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultPmsProductCategoryModel) Insert(ctx context.Context, data *PmsProductCategory) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", m.table, pmsProductCategoryRowsExpectAutoSet)
	ret, err := m.conn.ExecCtx(ctx, query, data.ParentId, data.Name, data.Level, data.ProductCount, data.ProductUnit, data.NavStatus, data.ShowStatus, data.Sort, data.Icon, data.Keywords, data.Description)
	return ret, err
}

func (m *defaultPmsProductCategoryModel) Update(ctx context.Context, data *PmsProductCategory) error {
	query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, pmsProductCategoryRowsWithPlaceHolder)
	_, err := m.conn.ExecCtx(ctx, query, data.ParentId, data.Name, data.Level, data.ProductCount, data.ProductUnit, data.NavStatus, data.ShowStatus, data.Sort, data.Icon, data.Keywords, data.Description, data.Id)
	return err
}

func (m *defaultPmsProductCategoryModel) tableName() string {
	return m.table
}