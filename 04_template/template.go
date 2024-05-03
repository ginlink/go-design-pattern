package template

import "fmt"

/*
模板模式
- 抽象类定义一系列操作
- 模板类中定义一系列算法（或者流程）

例子：煮茶、煮咖啡
基本操作有：煮开水、冲泡、倒入杯中、加料、是否加料
模板：先煮开水、再冲泡、再倒入杯中、根据是否加料决定加料
*/

// ----- 抽象层 -----

type Beverage interface {
	BoilWater() // 煮开水
	Brew()      // 冲泡
	PourInCup() // 倒入杯中

	WantToAddThings() bool // 是否加料
	AddThings()            // 加料
}

// ----- 实现层 -----

type template struct {
	b Beverage
}

func (t *template) MakeBeverage() {
	if t.b == nil {
		return
	}

	t.b.BoilWater()
	t.b.Brew()
	t.b.PourInCup()

	if t.b.WantToAddThings() {
		t.b.AddThings()
	}
}

type MakeCoffee struct {
	template
}

func (c *MakeCoffee) BoilWater() {
	fmt.Println("为了喝咖啡，煮水")
}
func (c *MakeCoffee) Brew() {
	fmt.Println("冲泡咖啡")
}
func (c *MakeCoffee) PourInCup() {
	fmt.Println("将咖啡导入杯中")
}
func (c *MakeCoffee) WantToAddThings() bool {
	return false // 不想加东西，不加糖的咖啡好喝
}
func (c *MakeCoffee) AddThings() {
	fmt.Println("加入糖")
}

type MakeTea struct {
	template
}

func (c *MakeTea) BoilWater() {
	fmt.Println("为了喝茶，煮水")
}
func (c *MakeTea) Brew() {
	fmt.Println("冲泡茶叶")
}
func (c *MakeTea) PourInCup() {
	fmt.Println("将茶水导入杯中")
}
func (c *MakeTea) WantToAddThings() bool {
	return true // 喝茶就要加牛奶
}
func (c *MakeTea) AddThings() {
	fmt.Println("加牛奶")
}
