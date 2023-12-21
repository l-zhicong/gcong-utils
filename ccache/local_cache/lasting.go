package cache

import "fmt"

type Lasting struct {
}

func (*Lasting) lasting() {
	fmt.Println("持久化")
}
