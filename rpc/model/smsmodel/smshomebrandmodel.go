package smsmodel

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"strings"
	"zero-admin/rpc/sms/sms"
)

var _ SmsHomeBrandModel = (*customSmsHomeBrandModel)(nil)

type (
	// SmsHomeBrandModel is an interface to be customized, add more methods here,
	// and implement the added methods in customSmsHomeBrandModel.
	SmsHomeBrandModel interface {
		smsHomeBrandModel
		Count(ctx context.Context, in *sms.HomeBrandListReq) (int64, error)
		FindAll(ctx context.Context, in *sms.HomeBrandListReq) (*[]SmsHomeBrand, error)
		FindOneByBrandId(ctx context.Context, brandId int64) (*SmsHomeBrand, error)
		DeleteByIds(ctx context.Context, ids []int64) error
	}

	customSmsHomeBrandModel struct {
		*defaultSmsHomeBrandModel
	}
)

// NewSmsHomeBrandModel returns a model for the database table.
func NewSmsHomeBrandModel(conn sqlx.SqlConn) SmsHomeBrandModel {
	return &customSmsHomeBrandModel{
		defaultSmsHomeBrandModel: newSmsHomeBrandModel(conn),
	}
}

func (m *customSmsHomeBrandModel) FindAll(ctx context.Context, in *sms.HomeBrandListReq) (*[]SmsHomeBrand, error) {

	where := "1=1"
	if len(in.BrandName) > 0 {
		where = where + fmt.Sprintf(" AND brand_name like '%%%s%%'", in.BrandName)
	}
	if in.RecommendStatus != 2 {
		where = where + fmt.Sprintf(" AND recommend_status = %d", in.RecommendStatus)
	}
	query := fmt.Sprintf("select %s from %s where %s limit ?,?", smsHomeBrandRows, m.table, where)
	var resp []SmsHomeBrand
	err := m.conn.QueryRows(&resp, query, (in.Current-1)*in.PageSize, in.PageSize)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *customSmsHomeBrandModel) FindOneByBrandId(ctx context.Context, brandId int64) (*SmsHomeBrand, error) {

	where := fmt.Sprintf("brand_id = %d", brandId)
	query := fmt.Sprintf("select %s from %s where %s ", smsHomeBrandRows, m.table, where)
	var resp SmsHomeBrand
	err := m.conn.QueryRow(&resp, query)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *customSmsHomeBrandModel) Count(ctx context.Context, in *sms.HomeBrandListReq) (int64, error) {
	where := "1=1"
	if len(in.BrandName) > 0 {
		where = where + fmt.Sprintf(" AND brand_name like '%%%s%%'", in.BrandName)
	}
	if in.RecommendStatus != 2 {
		where = where + fmt.Sprintf(" AND recommend_status = %d", in.RecommendStatus)
	}
	query := fmt.Sprintf("select count(*) as count from %s where %s", m.table, where)

	var count int64
	err := m.conn.QueryRow(&count, query)

	switch err {
	case nil:
		return count, nil
	case sqlc.ErrNotFound:
		return 0, ErrNotFound
	default:
		return 0, err
	}
}

func (m *customSmsHomeBrandModel) DeleteByIds(ctx context.Context, ids []int64) error {
	query := fmt.Sprintf("delete from %s where `id` in (?)", m.table)
	_, err := m.conn.ExecCtx(ctx, query, strings.Replace(strings.Trim(fmt.Sprint(ids), "[]"), " ", ",", -1))
	return err
}
