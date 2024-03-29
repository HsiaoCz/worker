package main

import "fmt"

// 抽象工厂（Abstract Factory）角色：工厂方法模式的核心，任何工厂类都必须实现这个接口。
// 工厂（Concrete Factory）角色：具体工厂类是抽象工厂的一个实现，负责实例化产品对象。
// 抽象产品（Abstract Product）角色：工厂方法模式所创建的所有对象的父类，它负责描述所有实例所共有的公共接口。
// 具体产品（Concrete Product）角色：工厂方法模式所创建的具体实例对象。

// ======= 抽象层 =========

// 水果类(抽象接口)
type Fruit interface {
	Show() //接口的某方法
}

// 工厂类(抽象接口)
type AbstractFactory interface {
	CreateFruit() Fruit //生产水果类(抽象)的生产器方法
}

// ======= 基础类模块 =========
type Apple struct {
	Fruit //为了易于理解显示继承(此行可以省略)
}

func (apple *Apple) Show() {
	fmt.Println("我是苹果")
}

type Banana struct {
	Fruit
}

func (banana *Banana) Show() {
	fmt.Println("我是香蕉")
}

type Pear struct {
	Fruit
}

func (pear *Pear) Show() {
	fmt.Println("我是梨")
}

// ========= 工厂模块  =========
// 具体的苹果工厂
type AppleFactory struct {
	AbstractFactory
}

func (fac *AppleFactory) CreateFruit() Fruit {
	//生产一个具体的苹果
	fruit := new(Apple)

	return fruit
}

// 具体的香蕉工厂
type BananaFactory struct {
	AbstractFactory
}

func (fac *BananaFactory) CreateFruit() Fruit {

	//生产一个具体的香蕉
	fruit := new(Banana)

	return fruit
}

// 具体的梨工厂
type PearFactory struct {
	AbstractFactory
}

func (fac *PearFactory) CreateFruit() Fruit {
	//生产一个具体的梨
	fruit := new(Pear)

	return fruit
}

// ======= 业务逻辑层 =======
func main() {
	//需求1：需要一个具体的苹果对象
	//1-先要一个具体的苹果工厂
	appleFac := new(AppleFactory)
	//2-生产相对应的具体水果
	apple := appleFac.CreateFruit()

	apple.Show()

	//需求2：需要一个具体的香蕉对象
	//1-先要一个具体的香蕉工厂
	bananaFac := new(BananaFactory)
	//2-生产相对应的具体水果
	banana := bananaFac.CreateFruit()

	banana.Show()

	//需求3：需要一个具体的梨对象
	//1-先要一个具体的梨工厂
	pearFac := new(PearFactory)
	//2-生产相对应的具体水果
	pear := pearFac.CreateFruit()

	pear.Show()

	//需求4：需要一个日本的苹果？
}
