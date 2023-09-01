日志

```
路径 默认为根目录log
异步时 请保证主线程运行
Config说明：
    LogPath //路径
    PoolNum //协程最大数
    IsAsync //是否异步
    IsPrint //是否打印到终端
```
实例代码
```
func main(){
    logs := New(&Config{"logsss", 2, false, true})
    logs.Info("111111%v", "222")
}
```

