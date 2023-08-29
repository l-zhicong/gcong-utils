package clog

import (
	"fmt"
	"github.com/l-zhicong/gcong-utils/cpool"
	"log"
	"os"
	"os/exec"
	"strings"
	"time"
)

var (
	fileName = time.Now().Format("2006-01-02")
	poolNum  = 1 //默认池数
)

const (
	INFO = "info"
	Err  = "error"
)

type clog struct {
	logPath string
	isAsync bool //开异步必须保证主线程在运行
	isPrint bool
	log     log.Logger
	grPool  *cpool.Pool
}

type Config struct {
	LogPath string
	PoolNum int
	isAsync bool
}

func New(config ...*Config) (logs clog) {
	if len(config) > 0 {
		con := config[0]
		logs.SetPath(con.LogPath)
		logs.grPool = cpool.New().SetConfig(con.PoolNum)
		logs.isAsync = con.isAsync
		return
	}
	logs.grPool = cpool.New().SetConfig(poolNum)
	logs.SetPath("log/" + time.Now().Format("2006-01")) //默认路径
	return
}

func (l *clog) SetPath(path ...string) *clog {
	if len(path) > 0 && len(path[0]) > 0 {
		rootPath, _ := getPath()
		l.logPath = rootPath + "/" + path[0]
	}
	checkPath(l.logPath)
	return l
}

func (l *clog) SetPool(poolNum int) clog {
	l.grPool.SetConfig(poolNum)
	return *l
}

func (l *clog) cFile(fileName string) *os.File {
	path := l.logPath + "/" + fileName + ".log"
	//检测目录，创建
	logFile, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		logFile, err = os.Create(path)
		if err != nil {
			panic(err)
		}
		return logFile
	} else {
		return logFile
	}
}

func (l *clog) Info(format string, param ...any) {
	Job := func() {
		l.log.SetOutput(l.cFile(fileName))
		l.log.Println(INFO + ":" + fmt.Sprintf(format, param...) + "\t" + time.Now().Format("2006-01-02 15:04:05"))
	}
	l.asyncLog(Job)
}

func (l *clog) Error(format string, param ...any) {
	Job := func() {
		l.log.SetOutput(l.cFile(fileName))
		l.log.Println(Err + ":" + fmt.Sprintf(format, param...) + "\t" + time.Now().Format("2006-01-02 15:04:05"))
	}
	l.asyncLog(Job)
}

func (l *clog) asyncLog(Job func()) {
	if l.isAsync {
		l.grPool.AddJob(Job)
	} else {
		Job()
	}
}

func (l *clog) print() {
	if l.isPrint {

	}
}

func getPath() (projectRoot string, err error) {
	cmd := exec.Command("go", "list", "-f", "{{.Dir}}", ".")
	output, err := cmd.Output()
	if err != nil {
		return
	}
	projectRoot = strings.TrimSpace(string(output)) //去掉空格
	return
}

func checkPath(logPath string) string {
	err := os.MkdirAll(logPath, os.ModePerm)
	if err != nil {
		panic(err)
	}
	return logPath
}
