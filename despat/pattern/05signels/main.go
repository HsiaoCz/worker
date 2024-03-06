package main

import (
	"fmt"
	"sync"
)

// 之前的单例属于懒汉式单例
// 不用考虑并发安全问题
// 缺点是 无论使用不使用 都会创建内存

// 饿汉式单例 只有在使用的时候才会创建内存

var lock sync.Mutex

type singelton struct{}

var instance *singelton

func GetInstance() *singelton {
	lock.Lock()
	defer lock.Unlock()

	if instance == nil {
		return new(singelton)
	} else {
		return instance
	}
}

func (s *singelton) SomeThing() {
	fmt.Println("单例的某个方法")
}

func main() {
	s := GetInstance()
	s.SomeThing()
}
