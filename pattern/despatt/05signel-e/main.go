package main

import (
	"sync"
)

// 单例的饿汉式

type signels struct{}

var instance *signels

var lock sync.Mutex

func GetInstance() *signels {
	lock.Lock()
	defer lock.Unlock()
	if instance == nil {
		instance = new(signels)
		return instance
	} else {
		return instance
	}
}

func main() {

}
