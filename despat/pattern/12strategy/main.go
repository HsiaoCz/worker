package main

import "fmt"

// Context (环境类):环境类时使用算法的角色,它在解决某个问题(即实现某个方法)时可以采用多种策略，在环境类种维护一个对抽象策略类的引用实例，用于定义所采用的策略

// Strategy (抽象策略类)：它为所支持的算法声明了抽象方法，是所有策略类的父类，它可以是抽象类或者具体类，也可以是接口。环境类通过抽象策略类中声明的方法运行时调用具体的策略类中实现算法

// ConcreteStrategy(具体策略类):它实现了在抽象类中声明的算法，在运行时，具体策略类将覆盖在环境类中定义的抽象策略类对象，使用一种具体的算法实现某个业务处理

// weapon strategy
type WeaponStrategy interface {
	UseWeapon() // use weapon
}

// Ak47 strategy
type Ak47 struct{}

func (a *Ak47) UseWeapon() {
	fmt.Println("use ak47 to fight")
}

// knife strategy
type Knife struct{}

func (k *Knife) UseWeapon() {
	fmt.Println("use knife to fight")
}

// envirment
type Hero struct {
	strategy WeaponStrategy
}

// set a strategy
func (h *Hero) SetWeaponStrategy(s WeaponStrategy) {
	h.strategy = s
}

func (h *Hero) Fight() {
	h.strategy.UseWeapon()
}

func main() {
	hero := &Hero{}

	hero.SetWeaponStrategy(new(Ak47))
	hero.Fight()

	hero.SetWeaponStrategy(new(Knife))
	hero.Fight()
}
