package {{.pkg}}

import (
    "autoGenCache/cachex"
	"context"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"time"
)

var _ {{.upperStartCamelObject}}Cache = (*custom{{.upperStartCamelObject}}Cache)(nil)

type (
	// {{.upperStartCamelObject}}Cache is an interface to be customized, add more methods here,
	// and implement the added methods in custom{{.upperStartCamelObject}}Cache.
	{{.upperStartCamelObject}}Cache interface {
		{{.lowerStartCamelObject}}Cache
	}

	custom{{.upperStartCamelObject}}Cache struct {
		*default{{.upperStartCamelObject}}Cache
	}
)

// New{{.upperStartCamelObject}}Cache returns a cache for the cache key.
func New{{.upperStartCamelObject}}Cache(c cache.CacheConf, opts ...cache.Option) {{.upperStartCamelObject}}Cache {
    {{if .expiry}}opts = append(opts, cache.WithExpiry(time.Second*{{.expiry}})){{end}}
    {{if .notFoundExpiry}}opts = append(opts, cache.WithNotFoundExpiry(time.Second*{{.notFoundExpiry}})){{end}}
	return &custom{{.upperStartCamelObject}}Cache{
		default{{.upperStartCamelObject}}Cache: new{{.upperStartCamelObject}}Cache(c, opts...),
	}
}

// queryCtxFn
//
//	@Description: When the cache does not exist, get data from here
func queryCtxFn(ctx context.Context, v any{{if .args}}, {{.args}}{{end}}) error {
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
