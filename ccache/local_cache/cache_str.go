package cache

import (
	"github.com/l-zhicong/gcong-utils/ccache/abs"
	"time"
)

// Get 没有key时 返回nil
func (lc *localCache) Get(key any) (value *abs.Value) {
	v, ok := lc.Load(key)
	if !ok {
		return nil
	}
	value = v.(abs.Data).Value
	return
}

// Set
// expireD 秒来计算
// TODO expireD 0持久的直接存文件 ok返回值没有
func (lc *localCache) Set(key any, value abs.Value, expireD int64) (ok bool) {
	lc.Store(key, abs.Data{
		Value: &value, ExpireD: time.Now().UnixMilli() + (expireD * 1000),
	})
	//if expireD == 0 {
	//	lc.lasting()
	//}
	return true
}

func (lc *localCache) Del(key any) {
	lc.Delete(key)
}
