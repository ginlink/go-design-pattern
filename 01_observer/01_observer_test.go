package observer

import (
	"fmt"
	"testing"
)

func TestObserver(t *testing.T) {
	// 初始化观察者
	zs := &Zhangsan{}
	ls := &Lisi{}

	// 初始化通知者
	// monitor := new(Monitor)
	monitor := &Monitor{}

	monitor.Attach(zs)
	monitor.Attach(ls)

	fmt.Println("正在上课中，但是老师没有来")
	zs.Doing()
	ls.Doing()

	fmt.Println("此时，老师来了")
	monitor.Notify()
}
