package clog

import (
	"fmt"
	"github.com/l-zhicong/gcong-utils/cpool"
	"log"
	"os"
	"os/exec"
	"path/filepath"
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
	isAsync bool
	log     log.Logger
	grPool  *cpool.Pool
}

type Config struct {
	LogPath string
	PoolNum int
}

func New(config ...*Config) (logs clog) {
	logs.grPool = cpool.New().SetConfig(poolNum)
	logs.SetPath("log/" + time.Now().Format("2006-01")) //默认路径
	if len(config) > 1 {
		con := config[0]
		logs.SetPath(con.LogPath)
		logs.SetPool(con.PoolNum)
	}
	return
}

func (l *clog) SetPath(path ...string) *clog {
	if len(path) > 0 && len(path[0]) > 0 {
		rootPath, _ := getPath()
		l.logPath = rootPath + "/" + path[0] + "/" + fileName + ".log"
	}
	l.log.SetOutput(cFile(l.logPath))
	return l
}

func (l *clog) SetPool(poolNum int) clog {
	l.grPool.SetConfig(poolNum)
	return *l
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

//TODO 未处理异常
func cFile(logPath string) *os.File {
	dir := filepath.Dir(logPath)
	//检测目录，创建
	logFile, err := os.OpenFile(logPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		err = os.MkdirAll(dir, os.ModePerm)
		if err != nil {
			panic(err)
		}
		logFile, err = os.Create(logPath)
		if err != nil {
			panic(err)
		}
		return logFile
	} else {
		return logFile
	}
}

func (l *clog) Info(format string, param ...any) {
	l.log.Println(INFO + ":" + fmt.Sprintf(format, param...) + "\t" + time.Now().Format("2006-01-02 15:04:05"))
}

func (l *clog) Error(format string, param ...any) {
	l.log.Println(Err + ":" + fmt.Sprintf(format, param...) + "\t" + time.Now().Format("2006-01-02 15:04:05"))
}
