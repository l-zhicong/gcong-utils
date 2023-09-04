package clog

import (
	"fmt"
	"os"
	"testing"
)

func TestName(t *testing.T) {
	logs := New(&Config{"logsss", 2, true})
	logs.Info("111111%v", "222")
	for true {

	}
}

func TestFile(t *testing.T) {
	rootPath, _ := getPath()
	logPath := rootPath + "/" + "log/lll"
	//dir := filepath.Dir(logPath)
	err := os.MkdirAll(logPath, os.ModePerm)
	if err != nil {
		panic(err)
	}
	fmt.Println(logPath)
	//return logPath

	//file, err := os.Create("/Users/lzc/wwwgo/myproject/gcong-test/2021/error.log")
	//if err != nil {
	//	fmt.Println("Error creating file:", err)
	//	return
	//}
	//defer file.Close()
	//return
}
