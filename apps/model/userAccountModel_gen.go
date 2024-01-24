// Code generated by goctl. DO NOT EDIT!

package model

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"time"

	"github.com/Masterminds/squirrel"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
	"zinx-zero/apps/acommon/globalkey"
)

var (
	userAccountFieldNames          = builder.RawFieldNames(&UserAccount{})
	userAccountRows                = strings.Join(userAccountFieldNames, ",")
	userAccountRowsExpectAutoSet   = strings.Join(stringx.Remove(userAccountFieldNames, "`create_time`", "`update_time`"), ",")
	userAccountRowsWithPlaceHolder = strings.Join(stringx.Remove(userAccountFieldNames, "`account_id`", "`create_time`", "`update_time`"), "=?,") + "=?"

	cacheGamexUserAccountAccountIdPrefix = "cache:gamex:userAccount:accountId:"
	cacheGamexUserAccountMobilePrefix    = "cache:gamex:userAccount:mobile:"
)

type (
	userAccountModel interface {
		Insert(ctx context.Context, session sqlx.Session, data *UserAccount) (sql.Result, error)
		FindOne(ctx context.Context, accountId int64) (*UserAccount, error)
		FindOneByMobile(ctx context.Context, mobile string) (*UserAccount, error)
		Update(ctx context.Context, session sqlx.Session, data *UserAccount) (sql.Result, error)
		UpdateWithVersion(ctx context.Context, session sqlx.Session, data *UserAccount) error
		Trans(ctx context.Context, fn func(context context.Context, session sqlx.Session) error) error
		SelectBuilder() squirrel.SelectBuilder
		DeleteSoft(ctx context.Context, session sqlx.Session, data *UserAccount) error
		FindSum(ctx context.Context, sumBuilder squirrel.SelectBuilder, field string) (float64, error)
		FindCount(ctx context.Context, countBuilder squirrel.SelectBuilder, field string) (int64, error)
		FindAll(ctx context.Context, rowBuilder squirrel.SelectBuilder, orderBy string) ([]*UserAccount, error)
		FindPageListByPage(ctx context.Context, rowBuilder squirrel.SelectBuilder, page, pageSize int64, orderBy string) ([]*UserAccount, error)
		FindPageListByPageWithTotal(ctx context.Context, rowBuilder squirrel.SelectBuilder, page, pageSize int64, orderBy string) ([]*UserAccount, int64, error)
		FindPageListByIdDESC(ctx context.Context, rowBuilder squirrel.SelectBuilder, preMinId, pageSize int64) ([]*UserAccount, error)
		FindPageListByIdASC(ctx context.Context, rowBuilder squirrel.SelectBuilder, preMaxId, pageSize int64) ([]*UserAccount, error)
		Delete(ctx context.Context, session sqlx.Session, accountId int64) error
	}

	defaultUserAccountModel struct {
		sqlc.CachedConn
		table string
	}

	UserAccount struct {
		AccountId  int64     `db:"account_id"`
		CreateTime time.Time `db:"create_time"`
		UpdateTime time.Time `db:"update_time"`
		DeleteTime time.Time `db:"delete_time"`
		DelState   int64     `db:"del_state"`
		Version    int64     `db:"version"` // 版本号
		Mobile     string    `db:"mobile"`
		Password   string    `db:"password"`
	}
)

func newUserAccountModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) *defaultUserAccountModel {
	return &defaultUserAccountModel{
		CachedConn: sqlc.NewConn(conn, c, opts...),
		table:      "`user_account`",
	}
}

func (m *defaultUserAccountModel) Insert(ctx context.Context, session sqlx.Session, data *UserAccount) (sql.Result, error) {
	data.DeleteTime = time.Unix(0, 0)
	data.DelState = globalkey.Sql_DelStateNo
	gamexUserAccountAccountIdKey := fmt.Sprintf("%s%v", cacheGamexUserAccountAccountIdPrefix, data.AccountId)
	gamexUserAccountMobileKey := fmt.Sprintf("%s%v", cacheGamexUserAccountMobilePrefix, data.Mobile)
	return m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?)", m.table, userAccountRowsExpectAutoSet)
		if session != nil {
			return session.ExecCtx(ctx, query, data.AccountId, data.DeleteTime, data.DelState, data.Version, data.Mobile, data.Password)
		}
		return conn.ExecCtx(ctx, query, data.AccountId, data.DeleteTime, data.DelState, data.Version, data.Mobile, data.Password)
	}, gamexUserAccountAccountIdKey, gamexUserAccountMobileKey)
}

func (m *defaultUserAccountModel) FindOne(ctx context.Context, accountId int64) (*UserAccount, error) {
	gamexUserAccountAccountIdKey := fmt.Sprintf("%s%v", cacheGamexUserAccountAccountIdPrefix, accountId)
	var resp UserAccount
	err := m.QueryRowCtx(ctx, &resp, gamexUserAccountAccountIdKey, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `account_id` = ? and del_state = ? limit 1", userAccountRows, m.table)
		return conn.QueryRowCtx(ctx, v, query, accountId, globalkey.Sql_DelStateNo)
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

func (m *defaultUserAccountModel) FindOneByMobile(ctx context.Context, mobile string) (*UserAccount, error) {
	gamexUserAccountMobileKey := fmt.Sprintf("%s%v", cacheGamexUserAccountMobilePrefix, mobile)
	var resp UserAccount
	err := m.QueryRowIndexCtx(ctx, &resp, gamexUserAccountMobileKey, m.formatPrimary, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) (i interface{}, e error) {
		query := fmt.Sprintf("select %s from %s where `mobile` = ? and del_state = ? limit 1", userAccountRows, m.table)
		if err := conn.QueryRowCtx(ctx, &resp, query, mobile, globalkey.Sql_DelStateNo); err != nil {
			return nil, err
		}
		return resp.AccountId, nil
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

func (m *defaultUserAccountModel) Update(ctx context.Context, session sqlx.Session, newData *UserAccount) (sql.Result, error) {
	data, err := m.FindOne(ctx, newData.AccountId)
	if err != nil {
		return nil, err
	}
	gamexUserAccountAccountIdKey := fmt.Sprintf("%s%v", cacheGamexUserAccountAccountIdPrefix, data.AccountId)
	gamexUserAccountMobileKey := fmt.Sprintf("%s%v", cacheGamexUserAccountMobilePrefix, data.Mobile)
	return m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `account_id` = ?", m.table, userAccountRowsWithPlaceHolder)
		if session != nil {
			return session.ExecCtx(ctx, query, newData.DeleteTime, newData.DelState, newData.Version, newData.Mobile, newData.Password, newData.AccountId)
		}
		return conn.ExecCtx(ctx, query, newData.DeleteTime, newData.DelState, newData.Version, newData.Mobile, newData.Password, newData.AccountId)
	}, gamexUserAccountAccountIdKey, gamexUserAccountMobileKey)
}

func (m *defaultUserAccountModel) UpdateWithVersion(ctx context.Context, session sqlx.Session, newData *UserAccount) error {

	oldVersion := newData.Version
	newData.Version += 1

	var sqlResult sql.Result
	var err error

	data, err := m.FindOne(ctx, newData.AccountId)
	if err != nil {
		return err
	}
	gamexUserAccountAccountIdKey := fmt.Sprintf("%s%v", cacheGamexUserAccountAccountIdPrefix, data.AccountId)
	gamexUserAccountMobileKey := fmt.Sprintf("%s%v", cacheGamexUserAccountMobilePrefix, data.Mobile)
	sqlResult, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `account_id` = ? and version = ? ", m.table, userAccountRowsWithPlaceHolder)
		if session != nil {
			return session.ExecCtx(ctx, query, newData.DeleteTime, newData.DelState, newData.Version, newData.Mobile, newData.Password, newData.AccountId, oldVersion)
		}
		return conn.ExecCtx(ctx, query, newData.DeleteTime, newData.DelState, newData.Version, newData.Mobile, newData.Password, newData.AccountId, oldVersion)
	}, gamexUserAccountAccountIdKey, gamexUserAccountMobileKey)
	if err != nil {
		return err
	}
	updateCount, err := sqlResult.RowsAffected()
	if err != nil {
		return err
	}
	if updateCount == 0 {
		return ErrNoRowsUpdate
	}

	return nil
}

func (m *defaultUserAccountModel) DeleteSoft(ctx context.Context, session sqlx.Session, data *UserAccount) error {
	data.DelState = globalkey.Sql_DelStateYes
	data.DeleteTime = time.Now()
	if err := m.UpdateWithVersion(ctx, session, data); err != nil {
		return errors.Wrapf(errors.New("delete soft failed "), "UserAccountModel delete err : %+v", err)
	}
	return nil
}

func (m *defaultUserAccountModel) FindSum(ctx context.Context, builder squirrel.SelectBuilder, field string) (float64, error) {

	if len(field) == 0 {
		return 0, errors.Wrapf(errors.New("FindSum Least One Field"), "FindSum Least One Field")
	}

	builder = builder.Columns("IFNULL(SUM(" + field + "),0)")

	query, values, err := builder.Where("del_state = ?", globalkey.Sql_DelStateNo).ToSql()
	if err != nil {
		return 0, err
	}

	var resp float64
	err = m.QueryRowNoCacheCtx(ctx, &resp, query, values...)
	switch err {
	case nil:
		return resp, nil
	default:
		return 0, err
	}
}

func (m *defaultUserAccountModel) FindCount(ctx context.Context, builder squirrel.SelectBuilder, field string) (int64, error) {

	if len(field) == 0 {
		return 0, errors.Wrapf(errors.New("FindCount Least One Field"), "FindCount Least One Field")
	}

	builder = builder.Columns("COUNT(" + field + ")")

	query, values, err := builder.Where("del_state = ?", globalkey.Sql_DelStateNo).ToSql()
	if err != nil {
		return 0, err
	}

	var resp int64
	err = m.QueryRowNoCacheCtx(ctx, &resp, query, values...)
	switch err {
	case nil:
		return resp, nil
	default:
		return 0, err
	}
}

func (m *defaultUserAccountModel) FindAll(ctx context.Context, builder squirrel.SelectBuilder, orderBy string) ([]*UserAccount, error) {

	builder = builder.Columns(userAccountRows)

	if orderBy == "" {
		builder = builder.OrderBy("id DESC")
	} else {
		builder = builder.OrderBy(orderBy)
	}

	query, values, err := builder.Where("del_state = ?", globalkey.Sql_DelStateNo).ToSql()
	if err != nil {
		return nil, err
	}

	var resp []*UserAccount
	err = m.QueryRowsNoCacheCtx(ctx, &resp, query, values...)
	switch err {
	case nil:
		return resp, nil
	default:
		return nil, err
	}
}

func (m *defaultUserAccountModel) FindPageListByPage(ctx context.Context, builder squirrel.SelectBuilder, page, pageSize int64, orderBy string) ([]*UserAccount, error) {

	builder = builder.Columns(userAccountRows)

	if orderBy == "" {
		builder = builder.OrderBy("id DESC")
	} else {
		builder = builder.OrderBy(orderBy)
	}

	if page < 1 {
		page = 1
	}
	offset := (page - 1) * pageSize

	query, values, err := builder.Where("del_state = ?", globalkey.Sql_DelStateNo).Offset(uint64(offset)).Limit(uint64(pageSize)).ToSql()
	if err != nil {
		return nil, err
	}

	var resp []*UserAccount
	err = m.QueryRowsNoCacheCtx(ctx, &resp, query, values...)
	switch err {
	case nil:
		return resp, nil
	default:
		return nil, err
	}
}

func (m *defaultUserAccountModel) FindPageListByPageWithTotal(ctx context.Context, builder squirrel.SelectBuilder, page, pageSize int64, orderBy string) ([]*UserAccount, int64, error) {

	total, err := m.FindCount(ctx, builder, "id")
	if err != nil {
		return nil, 0, err
	}

	builder = builder.Columns(userAccountRows)

	if orderBy == "" {
		builder = builder.OrderBy("id DESC")
	} else {
		builder = builder.OrderBy(orderBy)
	}

	if page < 1 {
		page = 1
	}
	offset := (page - 1) * pageSize

	query, values, err := builder.Where("del_state = ?", globalkey.Sql_DelStateNo).Offset(uint64(offset)).Limit(uint64(pageSize)).ToSql()
	if err != nil {
		return nil, total, err
	}

	var resp []*UserAccount
	err = m.QueryRowsNoCacheCtx(ctx, &resp, query, values...)
	switch err {
	case nil:
		return resp, total, nil
	default:
		return nil, total, err
	}
}

func (m *defaultUserAccountModel) FindPageListByIdDESC(ctx context.Context, builder squirrel.SelectBuilder, preMinId, pageSize int64) ([]*UserAccount, error) {

	builder = builder.Columns(userAccountRows)

	if preMinId > 0 {
		builder = builder.Where(" id < ? ", preMinId)
	}

	query, values, err := builder.Where("del_state = ?", globalkey.Sql_DelStateNo).OrderBy("id DESC").Limit(uint64(pageSize)).ToSql()
	if err != nil {
		return nil, err
	}

	var resp []*UserAccount
	err = m.QueryRowsNoCacheCtx(ctx, &resp, query, values...)
	switch err {
	case nil:
		return resp, nil
	default:
		return nil, err
	}
}

func (m *defaultUserAccountModel) FindPageListByIdASC(ctx context.Context, builder squirrel.SelectBuilder, preMaxId, pageSize int64) ([]*UserAccount, error) {

	builder = builder.Columns(userAccountRows)

	if preMaxId > 0 {
		builder = builder.Where(" id > ? ", preMaxId)
	}

	query, values, err := builder.Where("del_state = ?", globalkey.Sql_DelStateNo).OrderBy("id ASC").Limit(uint64(pageSize)).ToSql()
	if err != nil {
		return nil, err
	}

	var resp []*UserAccount
	err = m.QueryRowsNoCacheCtx(ctx, &resp, query, values...)
	switch err {
	case nil:
		return resp, nil
	default:
		return nil, err
	}
}

func (m *defaultUserAccountModel) Trans(ctx context.Context, fn func(ctx context.Context, session sqlx.Session) error) error {

	return m.TransactCtx(ctx, func(ctx context.Context, session sqlx.Session) error {
		return fn(ctx, session)
	})

}

func (m *defaultUserAccountModel) SelectBuilder() squirrel.SelectBuilder {
	return squirrel.Select().From(m.table)
}
func (m *defaultUserAccountModel) Delete(ctx context.Context, session sqlx.Session, accountId int64) error {
	data, err := m.FindOne(ctx, accountId)
	if err != nil {
		return err
	}

	gamexUserAccountAccountIdKey := fmt.Sprintf("%s%v", cacheGamexUserAccountAccountIdPrefix, accountId)
	gamexUserAccountMobileKey := fmt.Sprintf("%s%v", cacheGamexUserAccountMobilePrefix, data.Mobile)
	_, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `account_id` = ?", m.table)
		if session != nil {
			return session.ExecCtx(ctx, query, accountId)
		}
		return conn.ExecCtx(ctx, query, accountId)
	}, gamexUserAccountAccountIdKey, gamexUserAccountMobileKey)
	return err
}
func (m *defaultUserAccountModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheGamexUserAccountAccountIdPrefix, primary)
}
func (m *defaultUserAccountModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where `account_id` = ? and del_state = ? limit 1", userAccountRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary, globalkey.Sql_DelStateNo)
}

func (m *defaultUserAccountModel) tableName() string {
	return m.table
}