package adapter

import "fmt"

/**
适配器模式
- 适配目标
- 适配器
- 适配者

示例：手机通过适配器使用220v充电 (实际上手机通过5v充电)
- 适配目标：5v充电
- 适配器：使用220v进行5v充电
- 适配者：220v充电
*/

type V5 interface {
	Use5V()
}

type V220 struct{}

func (v *V220) Use220V() {
	fmt.Println("使用220v电压")
}

type Adapter struct {
	v220 *V220
}

func (a *Adapter) Use5V() {
	fmt.Println("使用适配充电")
	a.v220.Use220V()
}

// 手机
type Phone struct {
	v V5
}

func (p *Phone) Charge() {
	fmt.Println("充电")
	p.v.Use5V()
}
