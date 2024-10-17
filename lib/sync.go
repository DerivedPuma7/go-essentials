package lib

import (
	"fmt"
	"time"
)

func Sync(param string, sleep int) {
	time.Sleep(time.Second * time.Duration(sleep))
	fmt.Println("executing:", param)
}