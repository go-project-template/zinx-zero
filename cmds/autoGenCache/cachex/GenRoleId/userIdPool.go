package GenRoleId

import (
	"autoGenCache/cachex"
	"context"
	"time"

	"github.com/zeromicro/go-zero/core/stores/cache"
)

var _ UserIdPoolCache = (*customUserIdPoolCache)(nil)

type (
	// UserIdPoolCache is an interface to be customized, add more methods here,
	// and implement the added methods in customUserIdPoolCache.
	UserIdPoolCache interface {
		userIdPoolCache
	}

	customUserIdPoolCache struct {
		*defaultUserIdPoolCache
	}
)

// NewUserIdPoolCache returns a cache for the cache key.
func NewUserIdPoolCache(c cache.CacheConf, opts ...cache.Option) UserIdPoolCache {
	opts = append(opts, cache.WithExpiry(time.Second*604800))
	opts = append(opts, cache.WithNotFoundExpiry(time.Second*60))
	return &customUserIdPoolCache{
		defaultUserIdPoolCache: newUserIdPoolCache(c, opts...),
	}
}

// queryCtxFn
//
//	@Description: When the cache does not exist, get data from here
func queryCtxFn(ctx context.Context, v any) error {
	// todo: add your logic here and delete this line
	return cachex.ErrNotHandleQuery
}

// execDelCtxFn
//
//	@Description: Before delete cache to do something, default to doing nothing.
//	@param ctx
//	@return any
//	@return error err is nil will delete cache.
func execDelCtxFn(ctx context.Context) (any, error) {
	// todo: add your logic here and delete this line
	return nil, nil
}
