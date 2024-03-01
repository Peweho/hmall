package test

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestThread(t *testing.T) {
	wg := &sync.WaitGroup{}
	f(wg)
	wg.Wait()
	fmt.Println("主线程结束")
}

func f(wg *sync.WaitGroup) {
	wg.Add(1)
	go func() {
		defer wg.Done()
		time.Sleep(2 * time.Second)
		fmt.Println("睡眠完成")
	}()
	fmt.Println("协程退出")
}
