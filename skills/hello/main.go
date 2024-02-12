package main

import "fmt"

func main() {
	fmt.Println("hello")
	fmt.Println(add(3, 4))
}

// 函数
func add(a int, b int) int {
	return a + b
}

// 结构体
type AddStruct struct {
	X int
	Y int
}

func (a *AddStruct) Add() int {
	return a.X + a.Y
}
