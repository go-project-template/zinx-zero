package cachex

import (
	"context"
	"database/sql"
	"errors"
	"github.com/zeromicro/go-zero/core/logx"
	"time"

	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/syncx"
)

var (
	// ErrNotFound is an alias of sqlx.ErrNotFound.
	ErrNotFound = sqlx.ErrNotFound
	// ErrNotHandleQuery is not handle query then throw err to handle other things.
    ErrNotHandleQuery = errors.New("cache: not handle query")

	// can't use one SingleFlight per conn, because multiple conns may share the same cache key.
	singleFlights = syncx.NewSingleFlight()
	stats         = cache.NewStat("cachex")
)

type (
	// ExecFn defines the sql exec method.
	ExecFn func() (any, error)
	// ExecCtxFn defines the sql exec method.
	ExecCtxFn func(ctx context.Context) (any, error)
	// PrimaryQueryFn defines the query method that based on primary keys.
	PrimaryQueryFn func(v, primary any) error
	// PrimaryQueryCtxFn defines the query method that based on primary keys.
	PrimaryQueryCtxFn func(ctx context.Context, v, primary any) error
	// QueryFn defines the query method.
	QueryFn func(v any) error
	// QueryCtxFn defines the query method.
	QueryCtxFn func(ctx context.Context, v any) error

	// A CachedConn is a DB connection with cache capability.
	CachedConn struct {
		cache cache.Cache
	}
)

// NewConn returns a CachedConn with a redis cluster cache.
func NewConn(c cache.CacheConf, opts ...cache.Option) CachedConn {
	cc := cache.New(c, singleFlights, stats, sql.ErrNoRows, opts...)
	return NewConnWithCache(cc)
}

// NewConnWithCache returns a CachedConn with a custom cache.
func NewConnWithCache(c cache.Cache) CachedConn {
	return CachedConn{
		cache: c,
	}
}

// NewNodeConn returns a CachedConn with a redis node cache.
func NewNodeConn(rds *redis.Redis, opts ...cache.Option) CachedConn {
	c := cache.NewNode(rds, singleFlights, stats, sql.ErrNoRows, opts...)
	return NewConnWithCache(c)
}

// DelCache deletes cache with keys.
func (cc CachedConn) DelCache(keys ...string) error {
	return cc.DelCacheCtx(context.Background(), keys...)
}

// DelCacheCtx deletes cache with keys.
func (cc CachedConn) DelCacheCtx(ctx context.Context, keys ...string) error {
	return cc.cache.DelCtx(ctx, keys...)
}

// GetCache unmarshals cache with given key into v.
func (cc CachedConn) GetCache(key string, v any) error {
	return cc.GetCacheCtx(context.Background(), key, v)
}

// GetCacheCtx unmarshals cache with given key into v.
func (cc CachedConn) GetCacheCtx(ctx context.Context, key string, v any) error {
	return cc.cache.GetCtx(ctx, key, v)
}

// Exec runs given exec on given keys, and returns execution result.
func (cc CachedConn) Exec(exec ExecFn, keys ...string) (any, error) {
	execCtx := func(_ context.Context) (any, error) {
		return exec()
	}
	return cc.ExecCtx(context.Background(), execCtx, keys...)
}

// ExecCtx runs given exec on given keys, and returns execution result.
// If DB operation succeeds, it will delete cache with given keys,
// if DB operation fails, it will return nil result and non-nil error,
// if DB operation succeeds but cache deletion fails, it will return result and non-nil error.
func (cc CachedConn) ExecCtx(ctx context.Context, exec ExecCtxFn, keys ...string) (
	any, error) {
	res, err := exec(ctx)
	if err != nil {
		return nil, err
	}

	return res, cc.DelCacheCtx(ctx, keys...)
}

// QueryRow unmarshals into v with given key and query func.
func (cc CachedConn) QueryRow(v any, key string, query QueryFn) error {
	queryCtx := func(_ context.Context, v any) error {
		return query(v)
	}
	return cc.QueryRowCtx(context.Background(), v, key, queryCtx)
}

// QueryRowCtx unmarshals into v with given key and query func.
func (cc CachedConn) QueryRowCtx(ctx context.Context, v any, key string, query QueryCtxFn) error {
	return cc.cache.TakeCtx(ctx, v, key, func(v any) error {
		return query(ctx, v)
	})
}

// SetCache sets v into cache with given key.
func (cc CachedConn) SetCache(key string, val any) error {
	return cc.SetCacheCtx(context.Background(), key, val)
}

// SetCacheCtx sets v into cache with given key.
func (cc CachedConn) SetCacheCtx(ctx context.Context, key string, val any) error {
	return cc.cache.SetCtx(ctx, key, val)
}

// SetCacheWithExpire sets v into cache with given key with given expire.
func (cc CachedConn) SetCacheWithExpire(key string, val any, expire time.Duration) error {
	return cc.SetCacheWithExpireCtx(context.Background(), key, val, expire)
}

// SetCacheWithExpireCtx sets v into cache with given key with given expire.
func (cc CachedConn) SetCacheWithExpireCtx(ctx context.Context, key string, val any,
	expire time.Duration) error {
	return cc.cache.SetWithExpireCtx(ctx, key, val, expire)
}

// SetCacheCtxFromFn sets v (from getData func) into cache with given key.
func (cc CachedConn) SetCacheCtxFromFn(ctx context.Context, key string, getData func() (any, error)) (any, error) {
	logger := logx.WithContext(ctx)
	data, err := singleFlights.Do(key, func() (any, error) {
		data, err := getData()
		if err != nil {
			return nil, err
		}
		if err = cc.cache.SetCtx(ctx, key, data); err != nil {
			logger.Error("set cache", logx.Field("err", err))
		}
		return data, nil
	})
	return data, err
}