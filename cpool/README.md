# gcong-utils
GMP 模型
G goroutine 执行的任务
M thread 线程
P processor 处理器 

调度器：全局队列G -> N个P的本地队列 ->  数组P -> 线程M

当然 在程序处理量较大时 我们可以设置GOMAXPROCS 限制cpu核心数，
0为自由分配
```
runtime.GOMAXPROCS(2)
```


在协程池中 创建了一个全局队列 放置任务的，用原子性限制最大协程数
```
func main(){
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
```