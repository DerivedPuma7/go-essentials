package goroutine

import (
	"fmt"
	"time"
)

// func init() {
// 	go load()
// 	countUntil(5)
// }

func load() {
	for {
		fmt.Println("...")
		time.Sleep(time.Millisecond * 600)
	}
}

func countUntil(limit int) {
	for i := 0; i < limit; i++ {
		time.Sleep(time.Second * 1)
	}
	fmt.Println(limit)
}
