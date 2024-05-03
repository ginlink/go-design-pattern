package strategy

import "fmt"

/*
两个类：1.策略类 2.环境类
- 策略是环境类的一个属性
- 策略可以被设置
*/

// ----- 抽象层 -----
type Strategy interface {
	UseWeapon()
}

// ----- 实现层 -----
type Knife struct{}

func (nf *Knife) UseWeapon() {
	fmt.Println("使用小刀")
}

type Ak47 struct{}

func (ak *Ak47) UseWeapon() {
	fmt.Println("使用 ak")
}

type Hero struct {
	s Strategy
}

func (h *Hero) SetStrategy(s Strategy) {
	h.s = s
}
func (h *Hero) Fight() {
	h.s.UseWeapon()
}

// ----- 业务逻辑层 -----
// 在测试文件中
