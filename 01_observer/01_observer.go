package observer

import (
	"fmt"
)

/* 两个角色
- 被观察的目标（通知者）
  - Attach(obs Observer)
  - Detach(obs Observer)
	- Notify()
- 观察者
	- Doing()
	- WhatToDo()
*/

// 抽象出层
type Notifier interface {
	Attach(obs Observer)
	Detach(obs Observer)
	Notify()
}

type Observer interface {
	WhatToDo()
}

// 实现层
type Monitor struct {
	listeners []Observer
}

func (m *Monitor) Attach(obs Observer) {
	m.listeners = append(m.listeners, obs)
}

func (m *Monitor) Detach(obs Observer) {
	for index, item := range m.listeners {
		if item == obs {
			m.listeners = append(m.listeners[:index], m.listeners[index+1:]...)
		}
	}
}

func (m *Monitor) Notify() {
	for _, item := range m.listeners {
		item.WhatToDo()
	}
}

type Zhangsan struct{}

func (zs *Zhangsan) Doing() {
	fmt.Println("zhangsan 正在抽烟")
}
func (zs *Zhangsan) WhatToDo() {
	fmt.Println("zhangsan 把烟扔了")
}

type Lisi struct{}

func (ls *Lisi) Doing() {
	fmt.Println("lisi 正在抄作业")
}
func (ls *Lisi) WhatToDo() {
	fmt.Println("lisi 停止抄作业")
}
