package main

import (

	"sync"
)

type RecursiveMutex struct {
	sync.Mutex
	owner int64//当前持有锁的goroutine id
	recursion int64 //重入次数
}

func (m* RecursiveMutex)Lock()  {


}
func main() {

}
