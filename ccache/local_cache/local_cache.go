package cache

import (
	"context"
	"github.com/l-zhicong/gcong-utils/ccache/abs"
	"sync"
	"time"
)

const (
	DefaultExpiration  = 0               //默认过期时间 0为无限
	ClearExpirInterval = time.Minute * 1 // 默认1分钟清理一次过期
)

type localCache struct {
	//sync.RWMutex 缓存读比较多
	context.Context
	sync.Map
	interval time.Duration //定时清理的时间
	//Lasting
}

func NewLocalCache() abs.CacheInterface {
	cacheObj := localCache{}
	cacheObj.interval = ClearExpirInterval
	go func() {
		cacheObj.clearExpire()
	}()
	return &cacheObj
}

func (lc *localCache) SetInterval(duration time.Duration) {
	lc.interval = duration
}

// clearExpire 定时清理过期的
func (lc *localCache) clearExpire() {
	timeTicker := time.NewTicker(lc.interval)
	for {
		select {
		case <-timeTicker.C:
			lc.Range(func(key, value interface{}) bool {
				data := value.(abs.Data)
				if time.Now().UnixMilli() > data.ExpireD {
					//fmt.Println("过期key", key, data)
					lc.Delete(key)
				}
				//fmt.Println("Key:", key, "Value:", time.Now().UnixMilli(), data)
				return true
			})
		}
	}
}
