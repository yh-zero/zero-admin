// Code generated by goctl. DO NOT EDIT.

package smsmodel

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
	smsHomeRecommendSubjectFieldNames          = builder.RawFieldNames(&SmsHomeRecommendSubject{})
	smsHomeRecommendSubjectRows                = strings.Join(smsHomeRecommendSubjectFieldNames, ",")
	smsHomeRecommendSubjectRowsExpectAutoSet   = strings.Join(stringx.Remove(smsHomeRecommendSubjectFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	smsHomeRecommendSubjectRowsWithPlaceHolder = strings.Join(stringx.Remove(smsHomeRecommendSubjectFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"
)

type (
	smsHomeRecommendSubjectModel interface {
		Insert(ctx context.Context, data *SmsHomeRecommendSubject) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*SmsHomeRecommendSubject, error)
		Update(ctx context.Context, data *SmsHomeRecommendSubject) error
		Delete(ctx context.Context, id int64) error
	}

	defaultSmsHomeRecommendSubjectModel struct {
		conn  sqlx.SqlConn
		table string
	}

	SmsHomeRecommendSubject struct {
		Id              int64  `db:"id"`
		SubjectId       int64  `db:"subject_id"`       // 专题id
		SubjectName     string `db:"subject_name"`     // 专题名称
		RecommendStatus int64  `db:"recommend_status"` // 推荐状态：0->不推荐;1->推荐
		Sort            int64  `db:"sort"`             // 排序
	}
)

func newSmsHomeRecommendSubjectModel(conn sqlx.SqlConn) *defaultSmsHomeRecommendSubjectModel {
	return &defaultSmsHomeRecommendSubjectModel{
		conn:  conn,
		table: "`sms_home_recommend_subject`",
	}
}

func (m *defaultSmsHomeRecommendSubjectModel) Delete(ctx context.Context, id int64) error {
	query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, id)
	return err
}

func (m *defaultSmsHomeRecommendSubjectModel) FindOne(ctx context.Context, id int64) (*SmsHomeRecommendSubject, error) {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", smsHomeRecommendSubjectRows, m.table)
	var resp SmsHomeRecommendSubject
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

func (m *defaultSmsHomeRecommendSubjectModel) Insert(ctx context.Context, data *SmsHomeRecommendSubject) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?)", m.table, smsHomeRecommendSubjectRowsExpectAutoSet)
	ret, err := m.conn.ExecCtx(ctx, query, data.SubjectId, data.SubjectName, data.RecommendStatus, data.Sort)
	return ret, err
}

func (m *defaultSmsHomeRecommendSubjectModel) Update(ctx context.Context, data *SmsHomeRecommendSubject) error {
	query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, smsHomeRecommendSubjectRowsWithPlaceHolder)
	_, err := m.conn.ExecCtx(ctx, query, data.SubjectId, data.SubjectName, data.RecommendStatus, data.Sort, data.Id)
	return err
}

func (m *defaultSmsHomeRecommendSubjectModel) tableName() string {
	return m.table
}
