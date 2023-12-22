package ccache

import (
	"fmt"
	"testing"
)

const (
	B = 1 << (iota * 10)
	KB
)

func TestName(t *testing.T) {
	CacheObj := NewCache()
	CacheObj.Set("key", "value", 30)
	value := CacheObj.Get("key")
	fmt.Println("222", value.String())
	select {}
}
