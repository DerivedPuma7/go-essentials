package main

import (
	// lib "testego/lib"
	// _ "testego/int"
	// _ "testego/string"
	// _ "testego/slice"
	_ "testego/function"
	// "fmt"
	// "sync"
	// "time"
)

func main() {

}

// func main() {
// 	start := time.Now()
	
// 	// lib.Sync("3", 3)
// 	// lib.Sync("2", 2)
// 	// lib.Sync("1", 1)

// 	var wg sync.WaitGroup
// 	wg.Add(3)
// 	channel := make(chan *string)

// 	go lib.Async(&wg, channel, "3", 3)
// 	go lib.Async(&wg, channel, "2", 2)
// 	go lib.Async(&wg, channel, "1", 1)

// 	go func() {
// 		wg.Wait()
// 		close(channel)
// 	}()

// 	for msg := range channel {
// 		fmt.Println(*msg)
// 	}

// 	diff := time.Since(start)
// 	fmt.Println("total", diff)
// }

