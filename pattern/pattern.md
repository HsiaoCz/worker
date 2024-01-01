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

#### 1.3、依赖倒转原则
