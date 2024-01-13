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

#### 1.7、迪米特法则

一个对象应当对其他对象尽可能少的了解，从而降低各个对象之间的耦合，提高系统的可维护性。例如在一个程序中，各个模块之间相互调用时，通常会提供一个统一的接口来实现。这样其他模块不需要了解另外一个模块的内部实现细节，这样当一个模块内部的实现发生改变时，不会影响其他模块的使用。（黑盒原理）

```go
package main

import "fmt"

// 迪米特法则
// 一个对象应该对其他对象尽可能少的了解
// 从而降低各个系统之间的耦合 提高系统的可维护性
// 迪米特法则 又叫最少知道原则
// 如果两个软件实体无须直接通信，那么就不应当发生直接的相互调用，可以通过第三方转发该调用

type Student struct {
	Id   string
	Name string
}

type Class struct {
	Id       string
	Name     string
	Students []Student
}

func (c Class) PrintStudents() {
	for _, student := range c.Students {
		fmt.Println(student)
	}
}

type School struct {
	Id      string
	Name    string
	Classes []Class
}

// 这里 school 与 student 没有直接的关系
func (s School) PrintAllStudents() {

	// 高耦合度的写法
	// for _, class := range s.Classes {
	// 	for _, student := range class.Students {
	// 		fmt.Println(student)
	// 	}
	// }

	// 降低了依赖关系的写法
	for _, class := range s.Classes {
		class.PrintStudents()
	}
}
```

### 2、设计模式

#### 2.1、简单工厂模式

如果没有工厂类

```go
package main

import "fmt"

//水果类
type Fruit struct {
	//...
	//...
	//...
}

func (f *Fruit) Show(name string) {
	if name == "apple" {
		fmt.Println("我是苹果")
	} else if name == "banana" {
		fmt.Println("我是香蕉")
	} else if name == "pear" {
		fmt.Println("我是梨")
	}
}

//创建一个Fruit对象
func NewFruit(name string) *Fruit {
	fruit := new(Fruit)

	if name == "apple" {
		//创建apple逻辑
	} else if name == "banana" {
		//创建banana逻辑
	} else if name == "pear" {
		//创建pear逻辑
	}

	return fruit
}

func main() {
	apple := NewFruit("apple")
	apple.Show("apple")

	banana := NewFruit("banana")
	banana.Show("banana")

	pear := NewFruit("pear")
	pear.Show("pear")
}
```

不难看出，Fruit 类是一个“巨大的”类，在该类的设计中存在如下几个问题：
(1) 在 Fruit 类中包含很多“if…else…”代码块，整个类的代码相当冗长，代码越长，阅读难度、维护难度和测试难度也越大；而且大量条件语句的存在还将影响系统的性能，程序在执行过程中需要做大量的条件判断。
(2) Fruit 类的职责过重，它负责初始化和显示所有的水果对象，将各种水果对象的初始化代码和显示代码集中在一个类中实现，违反了“单一职责原则”，不利于类的重用和维护；  
(3) 当需要增加新类型的水果时，必须修改 Fruit 类的构造函数 NewFruit()和其他相关方法源代码，违反了“开闭原则”

简单工厂模式

```go
package main

import "fmt"

// 简单工厂模式
// 工厂（Factory）角色：简单工厂模式的核心，它负责实现创建所有实例的内部逻辑。
// 工厂类可以被外界直接调用，创建所需的产品对象。
// 抽象产品（AbstractProduct）角色：简单工厂模式所创建的所有对象的父类，它负责描述所有实例所共有的公共接口。
// 具体产品（Concrete Product）角色：简单工厂模式所创建的具体实例对象

// ======= 抽象层 =========

// 水果类(抽象接口)
type Fruit interface {
	Show() //接口的某方法
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
// 一个工厂， 有一个生产水果的机器，返回一个抽象水果的指针
type Factory struct{}

func (fac *Factory) CreateFruit(kind string) Fruit {
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

// ==========业务逻辑层==============
func main() {
	factory := new(Factory)

	apple := factory.CreateFruit("apple")
	apple.Show()

	banana := factory.CreateFruit("banana")
	banana.Show()

	pear := factory.CreateFruit("pear")
	pear.Show()
}
```

简单工厂模式的优缺点:

优点：

1. 实现了对象创建和使用的分离。
2. 不需要记住具体类名，记住参数即可，减少使用者记忆量。

缺点：

1. 对工厂类职责过重，一旦不能工作，系统受到影响。
2. 增加系统中类的个数，复杂度和理解度增加。
3. 违反“开闭原则”，添加新产品需要修改工厂逻辑，工厂越来越复杂。

适用场景：

1.  工厂类负责创建的对象比较少，由于创建的对象较少，不会造成工厂方法中的业务逻辑太过复杂。
2.  客户端只知道传入工厂类的参数，对于如何创建对象并不关心。

#### 2.2、工厂模式

工厂模式里的角色和职责

抽象工厂（Abstract Factory）角色：工厂方法模式的核心，任何工厂类都必须实现这个接口。
工厂（Concrete Factory）角色：具体工厂类是抽象工厂的一个实现，负责实例化产品对象。
抽象产品（Abstract Product）角色：工厂方法模式所创建的所有对象的父类，它负责描述所有实例所共有的公共接口。
具体产品（Concrete Product）角色：工厂方法模式所创建的具体实例对象。

```go
package main

import "fmt"

// ======= 抽象层 =========

//水果类(抽象接口)
type Fruit interface {
	Show()	 //接口的某方法
}

//工厂类(抽象接口)
type AbstractFactory interface {
	CreateFruit() Fruit //生产水果类(抽象)的生产器方法
}

// ======= 基础类模块 =========
type Apple struct {
	Fruit  //为了易于理解显示继承(此行可以省略)
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
//具体的苹果工厂
type AppleFactory struct {
	AbstractFactory
}

func (fac *AppleFactory) CreateFruit() Fruit {
	var fruit Fruit

	//生产一个具体的苹果
	fruit = new(Apple)

	return fruit
}

//具体的香蕉工厂
type BananaFactory struct {
	AbstractFactory
}

func (fac *BananaFactory) CreateFruit() Fruit {
	var fruit Fruit

	//生产一个具体的香蕉
	fruit = new(Banana)

	return fruit
}


//具体的梨工厂
type PearFactory struct {
	AbstractFactory
}

func (fac *PearFactory) CreateFruit() Fruit {
	var fruit Fruit

	//生产一个具体的梨
	fruit = new(Pear)

	return fruit
}

//======= 业务逻辑层 =======
func main() {
	//需求1：需要一个具体的苹果对象
	//1-先要一个具体的苹果工厂
	var appleFac AbstractFactory
	appleFac = new(AppleFactory)
	//2-生产相对应的具体水果
	var apple Fruit
	apple = appleFac.CreateFruit()

	apple.Show()


	//需求2：需要一个具体的香蕉对象
	//1-先要一个具体的香蕉工厂
	var bananaFac AbstractFactory
	bananaFac = new(BananaFactory)
	//2-生产相对应的具体水果
	var banana Fruit
	banana = bananaFac.CreateFruit()

	banana.Show()

	//需求3：需要一个具体的梨对象
	//1-先要一个具体的梨工厂
	var pearFac AbstractFactory
	pearFac = new(PearFactory)
	//2-生产相对应的具体水果
	var pear Fruit
	pear = pearFac.CreateFruit()

	pear.Show()

	//需求4：需要一个日本的苹果？
}
```

当我们需要一个日本的苹果的时候 可以对代码作如下修改:

```go
package main

import "fmt"

// ======= 抽象层 =========

//水果类(抽象接口)
type Fruit interface {
	Show()	 //接口的某方法
}

//工厂类(抽象接口)
type AbstractFactory interface {
	CreateFruit() Fruit //生产水果类(抽象)的生产器方法
}

// ======= 基础类模块 =========
type Apple struct {
	Fruit  //为了易于理解显示继承(此行可以省略)
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

//(+) 新增一个"日本苹果"
type JapanApple struct {
	Fruit
}

func (jp *JapanApple) Show() {
	fmt.Println("我是日本苹果")
}

// ========= 工厂模块  =========
//具体的苹果工厂
type AppleFactory struct {
	AbstractFactory
}

func (fac *AppleFactory) CreateFruit() Fruit {
	var fruit Fruit

	//生产一个具体的苹果
	fruit = new(Apple)

	return fruit
}

//具体的香蕉工厂
type BananaFactory struct {
	AbstractFactory
}

func (fac *BananaFactory) CreateFruit() Fruit {
	var fruit Fruit

	//生产一个具体的香蕉
	fruit = new(Banana)

	return fruit
}


//具体的梨工厂
type PearFactory struct {
	AbstractFactory
}

func (fac *PearFactory) CreateFruit() Fruit {
	var fruit Fruit

	//生产一个具体的梨
	fruit = new(Pear)

	return fruit
}

//具体的日本工厂
type JapanAppleFactory struct {
	AbstractFactory
}

func (fac *JapanAppleFactory) CreateFruit() Fruit {
	var fruit Fruit

	//生产一个具体的日本苹果
	fruit = new(JapanApple)

	return fruit
}

// ========= 业务逻辑层  =========
func main() {
	/*
		本案例为了突出根据依赖倒转原则与面向接口编程特性。
	    一些变量的定义将使用显示类型声明方式
	*/

	//需求1：需要一个具体的苹果对象
	//1-先要一个具体的苹果工厂
	var appleFac AbstractFactory
	appleFac = new(AppleFactory)
	//2-生产相对应的具体水果
	var apple Fruit
	apple = appleFac.CreateFruit()

	apple.Show()


	//需求2：需要一个具体的香蕉对象
	//1-先要一个具体的香蕉工厂
	var bananaFac AbstractFactory
	bananaFac = new(BananaFactory)
	//2-生产相对应的具体水果
	var banana Fruit
	banana = bananaFac.CreateFruit()

	banana.Show()

	//需求3：需要一个具体的梨对象
	//1-先要一个具体的梨工厂
	var pearFac AbstractFactory
	pearFac = new(PearFactory)
	//2-生产相对应的具体水果
	var pear Fruit
	pear = pearFac.CreateFruit()

	pear.Show()

	//需求4：需要一个日本的苹果？
	//1-先要一个具体的日本评估工厂
	var japanAppleFac AbstractFactory
	japanAppleFac = new(JapanAppleFactory)
	//2-生产相对应的具体水果
	var japanApple Fruit
	japanApple = japanAppleFac.CreateFruit()

	japanApple.Show()
}

```

工厂方法模式的优缺点:

优点：

1. 不需要记住具体类名，甚至连具体参数都不用记忆。
2. 实现了对象创建和使用的分离。
3. 系统的可扩展性也就变得非常好，无需修改接口和原类。 4.对于新产品的创建，符合开闭原则。

缺点：

1. 增加系统中类的个数，复杂度和理解度增加。
2. 增加了系统的抽象性和理解难度。

适用场景：

1. 客户端不知道它所需要的对象的类。
2. 抽象工厂类通过其子类来指定创建哪个对象。

#### 2.3、抽象工厂模式

（1）当添加一个新产品的时候，比如葡萄，虽然不用修改代码，但是需要添加大量的类，而且还需要添加相对的工厂。（系统开销，维护成本）
（2）如果使用同一地域的水果（日本苹果，日本香蕉，日本梨），那么需要分别创建具体的工厂，如果选择出现失误，将会造成混乱，虽然可以加一些约束，但是代码实现变得复杂。
所以“抽象工厂方法模式”引出了“产品族”和“产品等级结构”概念，其目的是为了更加高效的生产同一个产品组产品。

抽象工厂的角色和职责:

抽象工厂（Abstract Factory）角色：它声明了一组用于创建一族产品的方法，每一个方法对应一种产品。
具体工厂（Concrete Factory）角色：它实现了在抽象工厂中声明的创建产品的方法，生成一组具体产品，这些产品构成了一个产品族，每一个产品都位于某个产品等级结构中。
抽象产品（Abstract Product）角色：它为每种产品声明接口，在抽象产品中声明了产品所具有的业务方法。
具体产品（Concrete Product）角色：它定义具体工厂生产的具体产品对象，实现抽象产品接口中声明的业务方法。

```go
package main

import "fmt"

// ======= 抽象层 =========
type AbstractApple interface {
	ShowApple()
}

type AbstractBanana interface {
	ShowBanana()
}

type AbstractPear interface {
	ShowPear()
}

//抽象工厂
type AbstractFactory interface {
	CreateApple() AbstractApple
	CreateBanana() AbstractBanana
	CreatePear() AbstractPear
}

// ======== 实现层 =========
/*  中国产品族 */
type ChinaApple struct {}

func (ca *ChinaApple) ShowApple() {
	fmt.Println("中国苹果")
}

type ChinaBanana struct {}

func (cb *ChinaBanana) ShowBanana() {
	fmt.Println("中国香蕉")
}

type ChinaPear struct {}

func (cp *ChinaPear) ShowPear() {
	fmt.Println("中国梨")
}

type ChinaFactory struct {}

func (cf *ChinaFactory) CreateApple() AbstractApple {
	var apple AbstractApple

	apple = new(ChinaApple)

	return apple
}

func (cf *ChinaFactory) CreateBanana() AbstractBanana {
	var banana AbstractBanana

	banana = new(ChinaBanana)

	return banana
}

func (cf *ChinaFactory) CreatePear() AbstractPear {
	var pear AbstractPear

	pear = new(ChinaPear)

	return pear
}

/*  日本产品族 */
type JapanApple struct {}

func (ja *JapanApple) ShowApple() {
	fmt.Println("日本苹果")
}

type JapanBanana struct {}

func (jb *JapanBanana) ShowBanana() {
	fmt.Println("日本香蕉")
}

type JapanPear struct {}

func (cp *JapanPear) ShowPear() {
	fmt.Println("日本梨")
}

type JapanFactory struct {}

func (jf *JapanFactory) CreateApple() AbstractApple {
	var apple AbstractApple

	apple = new(JapanApple)

	return apple
}

func (jf *JapanFactory) CreateBanana() AbstractBanana {
	var banana AbstractBanana

	banana = new(JapanBanana)

	return banana
}

func (cf *JapanFactory) CreatePear() AbstractPear {
	var pear AbstractPear

	pear = new(JapanPear)

	return pear
}

/*  美国产品族 */
type AmericanApple struct {}

func (aa *AmericanApple) ShowApple() {
	fmt.Println("美国苹果")
}

type AmericanBanana struct {}

func (ab *AmericanBanana) ShowBanana() {
	fmt.Println("美国香蕉")
}

type AmericanPear struct {}

func (ap *AmericanPear) ShowPear() {
	fmt.Println("美国梨")
}

type AmericanFactory struct {}

func (af *AmericanFactory) CreateApple() AbstractApple {
	var apple AbstractApple

	apple = new(AmericanApple)

	return apple
}

func (af *AmericanFactory) CreateBanana() AbstractBanana {
	var banana AbstractBanana

	banana = new(AmericanBanana)

	return banana
}

func (af *AmericanFactory) CreatePear() AbstractPear {
	var pear AbstractPear

	pear = new(AmericanPear)

	return pear
}

// ======== 业务逻辑层 =======
func main() {
	//需求1: 需要美国的苹果、香蕉、梨 等对象
	//1-创建一个美国工厂
	var aFac AbstractFactory
	aFac = new(AmericanFactory)

	//2-生产美国苹果
	var aApple AbstractApple
	aApple = aFac.CreateApple()
	aApple.ShowApple()

	//3-生产美国香蕉
	var aBanana AbstractBanana
	aBanana = aFac.CreateBanana()
	aBanana.ShowBanana()

	//4-生产美国梨
	var aPear AbstractPear
	aPear = aFac.CreatePear()
	aPear.ShowPear()

	//需求2: 需要中国的苹果、香蕉
	//1-创建一个中国工厂
	cFac := new(ChinaFactory)

	//2-生产中国苹果
	cApple := cFac.CreateApple()
	cApple.ShowApple()

	//3-生产中国香蕉
	cBanana := cFac.CreateBanana()
	cBanana.ShowBanana()
}

```

#### 2.4、单例模式

在单例类的内部实现只生成一个实例，同时它提供一个静态的 getInstance()工厂方法，
让客户可以访问它的唯一实例；为了防止在外部对其实例化，将其构造函数设计为私有；
在单例类内部定义了一个 Singleton 类型的静态对象，作为外部共享的唯一实例。

单例解决的问题:

保证一个类永远只能有一个对象，且该对象的功能依然能被其他模块使用。

```go
package main

import "fmt"

/*
三个要点：
		一是某个类只能有一个实例；
		二是它必须自行创建这个实例；
		三是它必须自行向整个系统提供这个实例。
*/

/*
	保证一个类永远只能有一个对象
*/


//1、保证这个类非公有化，外界不能通过这个类直接创建一个对象
//   那么这个类就应该变得非公有访问 类名称首字母要小写
type singelton struct {}

//2、但是还要有一个指针可以指向这个唯一对象，但是这个指针永远不能改变方向
//   Golang中没有常指针概念，所以只能通过将这个指针私有化不让外部模块访问
var instance *singelton = new(singelton)

//3、如果全部为私有化，那么外部模块将永远无法访问到这个类和对象，
//   所以需要对外提供一个方法来获取这个唯一实例对象
//   注意：这个方法是否可以定义为singelton的一个成员方法呢？
//       答案是不能，因为如果为成员方法就必须要先访问对象、再访问函数
//        但是类和对象目前都已经私有化，外界无法访问，所以这个方法一定是一个全局普通函数
func GetInstance() *singelton {
	return instance
}

func (s *singelton) SomeThing() {
	fmt.Println("单例对象的某方法")
}

func main() {
	s := GetInstance()
	s.SomeThing()
}
```
