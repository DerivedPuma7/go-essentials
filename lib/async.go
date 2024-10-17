package lib

import (
	"sync"
	"time"
)

func Async(wg *sync.WaitGroup, channel chan *string, param string, sleep int) {
	defer wg.Done()
	time.Sleep(time.Second * time.Duration(sleep))

	msg := ("executing " + param)
	channel <- &msg
}