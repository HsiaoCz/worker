package main

import "fmt"

// 装饰模式(Decorator Pattern)：动态地给一个对象增加一些额外的职责，就增加对象功能来说，装饰模式比生成子类实现更为灵活。装饰模式是一种对象结构型模式。
// Component（抽象构件）：它是具体构件和抽象装饰类的共同父类，声明了在具体构件中实现的业务方法，
// 它的引入可以使客户端以一致的方式处理未被装饰的对象以及装饰之后的对象，实现客户端的透明操作。
// ConcreteComponent（具体构件）：它是抽象构件类的子类，用于定义具体的构件对象，实现了在抽象构件中声明的方法，装饰器可以给它增加额外的职责（方法）

// 抽象的构件
type Phone interface {
	Show() // 功能
}

// 装饰器的基础类
type Decorator struct {
	phone Phone
}

func (d *Decorator) Show() {}

// 实现层
// 具体的构件
type HuaWei struct{}

func (hw *HuaWei) Show() {
	fmt.Println("秀出了HuaWei手机")
}

type XiaoMi struct{}

func (xm *XiaoMi) Show() {
	fmt.Println("秀出了小米手机")
}

// 具体的装饰器类
type MoDecorator struct {
	Decorator //继承基础装饰器类(主要继承Phone成员属性)
}

func (md *MoDecorator) Show() {
	md.phone.Show()      //调用被装饰构件的原方法
	fmt.Println("贴膜的手机") //装饰额外的方法
}

func NewMoDecorator(phone Phone) Phone {
	return &MoDecorator{Decorator{phone}}
}

type KeDecorator struct {
	Decorator //继承基础装饰器类(主要继承Phone成员属性)
}

func (kd *KeDecorator) Show() {
	kd.phone.Show()
	fmt.Println("手机壳的手机") //装饰额外的方法
}

func NewKeDecorator(phone Phone) Phone {
	return &KeDecorator{Decorator{phone}}
}

// ------------ 业务逻辑层 ---------
func main() {
	huawei := new(HuaWei)
	huawei.Show() //调用原构件方法

	fmt.Println("---------")
	//用贴膜装饰器装饰，得到新功能构件
	moHuawei := NewMoDecorator(huawei) //通过HueWei ---> MoHuaWei
	moHuawei.Show()                    //调用装饰后新构件的方法

	fmt.Println("---------")
	keHuawei := NewKeDecorator(huawei) //通过HueWei ---> KeHuaWei
	keHuawei.Show()

	fmt.Println("---------")
	keMoHuaWei := NewMoDecorator(keHuawei) //通过KeHuaWei ---> KeMoHuaWei
	keMoHuaWei.Show()
}
