package ccache

import (
	"github.com/l-zhicong/gcong-utils/ccache/abs"
	cache "github.com/l-zhicong/gcong-utils/ccache/local_cache"
)

func NewCache() abs.CacheInterface {
	return cache.NewLocalCache()
}
