package main

import (
	"fmt"
	"sync"
)

func main() {
	//foo1()


	//var c Counter
	//c.Lock()
	//defer c.Unlock()
	//c.count++
	////call of foo2 copies lock value:
	//foo2(c)


	l:=&sync.Mutex{}
	foo3(l)
}
//释放没有加锁的锁
func foo1(){
	var mu sync.Mutex
	defer mu.Unlock()
	fmt.Println("hello world!")
}

type Counter struct {
	sync.Mutex
	count int
}
//mutex被复用的情况call of foo2 copies lock value:
func foo2(c Counter)  {
	c.Lock()
	defer c.Unlock()
	fmt.Println("in foo2")
}

//mutex不可重入：因为mutex不像java的reentrantLock 可以记录哪个线程拥有了当前这把锁
//发生死锁的现象
//foo3锁住了，bar无法获取锁，所以就没办法从bar中退出，因此foo3的锁也没办法解锁+》死锁
func foo3(l sync.Locker ){
	fmt.Println("in foo")
	l.Lock()
	bar(l)
	l.Unlock()
}
func bar(l sync.Locker){
	fmt.Println("in bar")
	l.Lock()

	l.Unlock()
}


//解决可重入
