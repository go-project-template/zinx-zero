package templ

import (
	_ "embed"
)

//go:embed tpl/cache.tpl
var Cache string

//go:embed tpl/cache_gen.tpl
var CacheGen string

//go:embed tpl/trace.tpl
var Trace string

//go:embed tpl/cacheddata.tpl
var CachedData string
