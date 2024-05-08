package facade

import "fmt"

/**
外观模式
- 通过一个统一入口，去调用不同子系统组合的函数

例子：家庭影院
- 想唱歌时，打开电视，关闭灯光
- 想照明时，关闭电视，打开灯光
*/

type TV struct{}

func (tv *TV) On() {
	fmt.Println("TV is on")
}
func (tv *TV) Off() {
	fmt.Println("TV is off")
}

type Light struct{}

func (lt *Light) On() {
	fmt.Println("Light is on")
}
func (lt *Light) Off() {
	fmt.Println("Light is off")
}

type HomePlayerFacade struct {
	tv TV
	lt Light
}

func (hpf *HomePlayerFacade) DoKTV() {
	fmt.Println("进入 KTV 模式")
	hpf.tv.On()
	hpf.lt.Off()
}
func (hpf *HomePlayerFacade) DoLight() {
	fmt.Println("进入照明模式")
	hpf.tv.Off()
	hpf.lt.On()
}
