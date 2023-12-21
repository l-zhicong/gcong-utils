package abs

import "github.com/l-zhicong/gcong-utils/conv"

type Value string

type Data struct {
	Value   *Value
	ExpireD int64 //秒作为计算
}

type CacheInterface interface {
	Get(key any) (value *Value)
	Set(key any, value Value, expireD int64) (ok bool)
	Del(key any)
}

func (v *Value) String() string {
	return conv.String(v)
}

func (v *Value) Int() int {
	return conv.Int(v)
}
