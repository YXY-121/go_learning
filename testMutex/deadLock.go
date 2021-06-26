package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var aIdentify sync.Mutex
	var bIdentify sync.Mutex
	var wg sync.WaitGroup
	wg.Add(2)

	//先处理a认证
	go func() {
		defer wg.Done()
		aIdentify.Lock()
		defer aIdentify.Unlock()

		//
		time.Sleep(5*time.Second)
		bIdentify.Unlock()


	}()
	go func() {
		defer wg.Done()
		bIdentify.Lock()
		defer bIdentify.Unlock()

		time.Sleep(5*time.Second)

		aIdentify.Lock()
		aIdentify.Unlock()

	}()
	wg.Wait()
	fmt.Println("成功完成")


}
