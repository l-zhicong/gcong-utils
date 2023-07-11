package conv

import (
	"fmt"
	"testing"
)

func TestToString(t *testing.T) {
	Any := 1
	value := String(Any)
	fmt.Println(value)
}

func TestToInt(t *testing.T) {
	Any := "1"
	value := Int(Any)
	fmt.Println(value)
}
