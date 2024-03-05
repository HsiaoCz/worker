package main

import "fmt"

// simple factory
// 工厂（Factory）角色：简单工厂模式的核心，它负责实现创建所有实例的内部逻辑。
// 工厂类可以被外界直接调用，创建所需的产品对象。
// 抽象产品（AbstractProduct）角色：简单工厂模式所创建的所有对象的父类，它负责描述所有实例所共有的公共接口。
// 具体产品（Concrete Product）角色：简单工厂模式所创建的具体实例对象

// 抽象层
// 水果类 (抽象接口)

type Fruit interface {
	Show() // 接口的方法
}

// 基础模块类
type Apple struct{}

func (a *Apple) Show() {
	fmt.Println("this is apple")
}

type Banana struct{}

func (b *Banana) Show() {
	fmt.Println("this is banana")
}

type Pear struct{}

func (p *Pear) Show() {
	fmt.Println("this is pear")
}

// 工厂模块
type Factory struct{}

func (f *Factory) CreateFruit(kind string) Fruit {
	var fruit Fruit
	if kind == "apple" {
		fruit = new(Apple)
	} else if kind == "banana" {
		fruit = new(Banana)
	} else if kind == "pear" {
		fruit = new(Pear)
	}
	return fruit
}

func main() {
	factory := new(Factory)

	apple := factory.CreateFruit("apple")
	apple.Show()

	banana := factory.CreateFruit("banana")
	banana.Show()

	pear := factory.CreateFruit("pear")
	pear.Show()
}
