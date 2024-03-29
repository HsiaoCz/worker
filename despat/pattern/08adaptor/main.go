package main

import "fmt"

// 适配器模式
// 将一个类的接口转换成客户希望的另外一个接口。使得原本由于接口不兼容而不能一起工作的那些类可以一起工作。
// Target（目标抽象类）：目标抽象类定义客户所需接口，可以是一个抽象类或接口，也可以是具体类。
// Adapter（适配器类）：适配器可以调用另一个接口，作为一个转换器，对 Adaptee 和 Target 进行适配，适配器类是适配器模式的核心，在对象适配器中，它通过继承 Target 并关联一个 Adaptee 对象使二者产生联系。
// Adaptee（适配者类）：适配者即被适配的角色，它定义了一个已经存在的接口，这个接口需要适配，适配者类一般是一个具体类，包含了客户希望使用的业务方法，在某些情况下可能没有适配者类的源代码。

// 适配的目标
type V5 interface {
	Use5V()
}

// 业务类 依赖V5接口
type Phone struct {
	v V5
}

func NewPhone(v V5) *Phone {
	return &Phone{
		v: v,
	}
}

func (p *Phone) Charge() {
	fmt.Println("phone 进行充电...")
	p.v.Use5V()
}

// 被适配的角色
type V220 struct{}

func (v *V220) Use220V() {
	fmt.Println("使用220V的电压")
}

// 电源适配器
type Adapter struct {
	v220 *V220
}

func (a *Adapter) Use5V() {
	fmt.Println("使用电源适配器进行充电")
	// 调用适配者的方法
	a.v220.Use220V()
}
func NewAdapter(v220 *V220) *Adapter {
	return &Adapter{v220}
}

// ------- 业务逻辑层 -------
func main() {
	iphone := NewPhone(NewAdapter(new(V220)))

	iphone.Charge()
}
