package main

import "fmt"

// 单例模式
type signel struct{}

var instance *signel = new(signel)

func NewInstance() *signel {
	return instance
}

func (s *signel) DoSomeThing() { fmt.Println("单例的某种方法") }

func main() {

}
