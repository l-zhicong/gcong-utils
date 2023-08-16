package cpool

import (
	"fmt"
	"sync/atomic"
	"testing"
	"time"
)

func TestName(t *testing.T) {
	f := func(num int) func() {
		j := num
		return func() {
			time.Sleep(10 * time.Millisecond)
			fmt.Println("执行任务", j)
		}
	}
	poolObj := New().SetConfig(10)
	for i := 0; i <= 10000; i++ {
		poolObj.AddJob(f(i))
	}
	for true {
		time.Sleep(time.Second * 1)
		fmt.Println(int(poolObj.GetCount()))
	}
}

func TestName1(t *testing.T) {
	var value int64 = 0
	fmt.Println("value", value)
	if atomic.CompareAndSwapInt64(&value, 0, 1) {
		fmt.Println(value)
	} else {
		fmt.Println("1")
	}
}
