package clog

import (
	"fmt"
	"os"
	"testing"
)

func TestName(t *testing.T) {
	logs := New()
	logs.Info("%v", "123")
}

func TestFile(t *testing.T) {

	file, err := os.Create("/Users/lzc/wwwgo/myproject/gcong-test/2021/error.log")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()
	return
}
