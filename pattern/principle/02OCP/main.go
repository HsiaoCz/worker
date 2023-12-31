package main

import "fmt"

type Driver interface {
	Drive(Car)
}

type Car interface {
	Run()
}

type Zhangsan struct{}

func (z *Zhangsan) Drive(car Car) {
	fmt.Println("zhangsan is Driving")
	car.Run()
}

type Lisi struct{}

func (l *Lisi) Drive(car Car) {
	fmt.Println("lisi is Driving")
	car.Run()
}

type BMW struct{}

func (b *BMW) Run() {
	fmt.Println("BMW is running...")
}

type Benz struct{}

func (b *Benz) Run() {
	fmt.Println("Benz is running...")
}

func main() {

}
