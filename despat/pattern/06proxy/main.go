package main

import "fmt"

// - subject（抽象主题角色）：真实主题与代理主题的共同接口。
// - RealSubject（真实主题角色）：定义了代理角色所代表的真实对象。
// - Proxy（代理主题角色）：含有对真实主题角色的引用，代理角色通常在将客户端调用传递给真是主题对象之前或者之后执行某些操作，而不是单纯返回真实的对象。

type Goods struct {
	Kind string
	Fact bool
}

// 抽象的购物主题
type Shopping interface {
	Buy(goods *Goods) // 某任务
}

// 实现层
// 具体的购物主题 subject

type KoreanShopping struct{}

func (ks *KoreanShopping) Buy(goods *Goods) {
	fmt.Println("去韩国进行了购物，买了:", goods.Kind)
}

// 具体的购物主题 实现了shopping 去非洲购物
type AfrikaShopping struct{}

func (as *AfrikaShopping) Buy(goods *Goods) {
	fmt.Println("去非洲进行了购物,买了:", goods.Kind)
}

// 海外的代理
type OverseasProxy struct {
	shopping Shopping
}

func (op *OverseasProxy) Buy(goods *Goods) {
	// 1. 先验货
	if op.distinguish(goods) {
		//2. 进行购买
		op.shopping.Buy(goods) //调用原被代理的具体主题任务
		//3 海关安检
		op.check(goods)
	}
}

// 创建一个代理,并且配置关联被代理的主题
func NewProxy(shopping Shopping) Shopping {
	return &OverseasProxy{shopping}
}

// 验货流程
func (op *OverseasProxy) distinguish(goods *Goods) bool {
	fmt.Println("对[", goods.Kind, "]进行了辨别真伪.")
	if !goods.Fact {
		fmt.Println("发现假货", goods.Kind, ", 不应该购买。")
	}
	return goods.Fact
}

// 安检流程
func (op *OverseasProxy) check(goods *Goods) {
	fmt.Println("对[", goods.Kind, "] 进行了海关检查， 成功的带回祖国")
}

func main() {
	g1 := Goods{
		Kind: "韩国面膜",
		Fact: true,
	}

	g2 := Goods{
		Kind: "CET4证书",
		Fact: false,
	}

	//如果不使用代理来完成从韩国购买任务
	shopping := new(KoreanShopping) //具体的购买主题

	//1-先验货
	if g1.Fact {
		fmt.Println("对[", g1.Kind, "]进行了辨别真伪.")
		//2-去韩国购买
		shopping.Buy(&g1)
		//3-海关安检
		fmt.Println("对[", g1.Kind, "] 进行了海关检查， 成功的带回祖国")
	}

	fmt.Println("---------------以下是 使用 代理模式-------")
	overseasProxy := NewProxy(shopping)
	overseasProxy.Buy(&g1)
	overseasProxy.Buy(&g2)
}
