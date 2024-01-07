## 设计原则与设计模式

### 1、设计原则

#### 1.1、单一职责原则

单一职责原则：类的职责单一，对外只提供一种功能，而且引起类变化的原因只有一个

举个打电话的例子:拨号、通话、回应、挂机

可以设计一个这样的接口

```go
type PhoneCall interface{
    Dial(phoneNumber int64)
    Talk(s struct{})
    Hangup()
}
```

单一职责原则要求一个接口或一个类只能有一个原因引起变化，也就是一个接口或者类只能有一个职责，它就负责一件事情
而 PhoneCall 这个接口做了两件事情：
协议管理和数据传送。Dial 和 Hangup 这两个方法实现的是协议管理，分别负责拨号接通和挂机.
Talk 方法实现的是数据传送。
不管是协议接通的变化还是输出传送的变化，都会引起这个接口的变化。

在面向对象编程的过程中，设计一个类，建议对外提供的功能单一，接口单一
影响一个类的范围就只限定在这一个接口上，一个类的一个接口具备这个类的功能含义，职责单一不复杂

举个例子：

```go
package main

import "fmt"

type workdress struct{}

func (w *workdress)OnWorkDress(){
    fmt.Println("工作的装扮...")
}

type shoppingdress struct{}

func (s *shoppingdress)OnShoppingDress(){
    fmt.Println("逛街的装扮....")
}

```

#### 1.2、开闭原则

开闭原则: 开:对扩展开放 闭：对修改关闭

```go
// 开闭原则
// 对扩展开放
// 对修改关闭

// 平铺式的代码
// type Banker struct{}

// func (b *Banker) Save() { fmt.Println("银行职员进行了存款的业务...") }

// func (b *Banker) Trans() { fmt.Println("银行职员进行了转账的业务....") }

// func (b *Banker) Stack() { fmt.Println("银行职员进行了股票的业务....") }

// 这种平铺的设计  在增加职员的作用的时候 需要修改原有的类
// 这中修改是危险的

// 基于开闭原则的代码设计

type Banker interface {
	DoBuz()
}

type SaveBanker struct{}

func (s *SaveBanker) DoBuz() { fmt.Println("银行职员进行了存款的业务....") }

type TransBanker struct{}

func (t *TransBanker) DoBuz() { fmt.Println("银行职员进行了转账的业务....") }

type StackBanker struct{}

func (s *StackBanker) DoBuz() { fmt.Println("银行职员进行了股票的业务....") }


```

当增加车辆或者司机的时候，不会对原有的代码进行破坏

接口的意义:
实现多态 调用未来

#### 1.3、依赖倒转原则

依赖于抽象 而不是依赖于具体的类

高耦合度的代码:

```go
package main

import "fmt"

// === > 奔驰汽车 <===
type Benz struct {

}

func (this *Benz) Run() {
	fmt.Println("Benz is running...")
}

// === > 宝马汽车  <===
type BMW struct {

}

func (this *BMW) Run() {
	fmt.Println("BMW is running ...")
}


//===> 司机张三  <===
type Zhang3 struct {
	//...
}

func (zhang3 *Zhang3) DriveBenZ(benz *Benz) {
	fmt.Println("zhang3 Drive Benz")
	benz.Run()
}

func (zhang3 *Zhang3) DriveBMW(bmw *BMW) {
	fmt.Println("zhang3 drive BMW")
	bmw.Run()
}

//===> 司机李四 <===
type Li4 struct {
	//...
}

func (li4 *Li4) DriveBenZ(benz *Benz) {
	fmt.Println("li4 Drive Benz")
	benz.Run()
}

func (li4 *Li4) DriveBMW(bmw *BMW) {
	fmt.Println("li4 drive BMW")
	bmw.Run()
}

func main() {
	//业务1 张3开奔驰
	benz := &Benz{}
	zhang3 := &Zhang3{}
	zhang3.DriveBenZ(benz)

	//业务2 李四开宝马
	bmw := &BMW{}
	li4 := &Li4{}
	li4.DriveBMW(bmw)
}
```

我们将这种代码改成面对抽象层依赖倒转

```go
package main

import "fmt"

// ===== >   抽象层  < ========
type Car interface {
	Run()
}

type Driver interface {
	Drive(car Car)
}

// ===== >   实现层  < ========
type BenZ struct {
	//...
}

func (benz * BenZ) Run() {
	fmt.Println("Benz is running...")
}

type Bmw struct {
	//...
}

func (bmw * Bmw) Run() {
	fmt.Println("Bmw is running...")
}

type Zhang_3 struct {
	//...
}

func (zhang3 *Zhang_3) Drive(car Car) {
	fmt.Println("Zhang3 drive car")
	car.Run()
}

type Li_4 struct {
	//...
}

func (li4 *Li_4) Drive(car Car) {
	fmt.Println("li4 drive car")
	car.Run()
}


// ===== >   业务逻辑层  < ========
func main() {
	//张3 开 宝马
	var bmw Car
	bmw = &Bmw{}

	var zhang3 Driver
	zhang3 = &Zhang_3{}

	zhang3.Drive(bmw)

	//李4 开 奔驰
	var benz Car
	benz = &BenZ{}

	var li4 Driver
	li4 = &Li_4{}

	li4.Drive(benz)
}
```

#### 1.4、里氏代换原则

任何抽象类出现的地方都可以用它的实现类进行替换
基类适用的，子类一定适用（子类可以扩展父类的功能，但不能改变父类原有的功能）

```go
package main

import "fmt"

// 里氏代换原则
// 基类适用的，子类一定适用（子类可以扩展父类的功能，但不能改变父类原有的功能）
// 这个在go里面似乎不太需要这样考虑

type ClassA struct{}

func (ca *ClassA) Add(x int, y int) int {
	return x + y
}

type ClassB struct {
	ClassA
}

func (cb *ClassB) Add(a string, b string) string {
	return a + b
}

func main() {
	a := 10
	b := 20
	cb := ClassB{}
	fmt.Println(cb.ClassA.Add(a, b))
	x := "hello,"
	y := "hi"
	fmt.Println(cb.Add(x, y))
}
```

#### 1.5、接口隔离原则

不应该强迫用户的程序依赖他们不需要的接口方法。
一个接口应该只提供一种对外功能，不应该把所有操作都封装到一个接口中去。

```go
package main

// 接口隔离 强调 不要把所有的方法封装到一个接口上
// 将大的接口设计成小的接口
// 不把所有的操作封装到一个接口上

// 假设A类 需要func1 func2 func5
// 假设B类 需要func1 func3 func4

type SomeISP interface {
	Func1()
	Func2()
	Func3()
	Func4()
	Func5()
}

// 类A 需要的是1 2 5 但是它还需要实现3 4
// 类B 需要的是1 3 4 但是它还需要实现2 5

func main() {

}
```
#### 1.6、合成复用原则

如果使用继承 会导致父类的任何变换都可能影响到子类，使用组合，则降低了这种
依赖关系，推荐使用组合而不是继承

```go
package main

import "fmt"

// 合成复用原则

type Cat struct{}

func (c *Cat) Eat() {
	fmt.Println("小猫吃饭")
}

// 给小猫添加睡觉的方法
// 通过继承的方式

type CatB struct {
	Cat
}

func (cb *CatB) Sleep() {
	fmt.Println("小猫睡觉")
}

// 给小猫添加睡觉的方法
// 通过组合的方式

type CatC struct {
	Ca Cat
}

func (cc *CatC) Sleep() {
	fmt.Println("小猫睡觉")
}

func main() {
	c:=Cat{}
	c.Eat()

	cb:=CatB{c}
	cb.Sleep()
	cb.Eat() //继承

	// 组合
	cc:=CatC{Ca: c}
	cc.Sleep()
	cc.Ca.Eat() // 组合
}

```