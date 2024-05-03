package command

import "fmt"

// ----- 抽象层 -----
type Command interface {
	Treat()
}

// ----- 实现层 -----
type Doctor struct {
}

func (d *Doctor) treatEye() {
	fmt.Println("医生治疗眼睛")
}
func (d *Doctor) treatNose() {
	fmt.Println("医生治疗鼻子")
}

type EyeCommand struct {
	doctor *Doctor
}

func (eyeCmd *EyeCommand) Treat() {
	eyeCmd.doctor.treatEye()
}

type NoseCommand struct {
	doctor *Doctor
}

func (nsCmd *NoseCommand) Treat() {
	nsCmd.doctor.treatNose()
}

type Nurse struct {
	CmdList []Command
}

func (ns *Nurse) Notify() {
	if ns.CmdList == nil {
		return
	}

	for _, item := range ns.CmdList {
		item.Treat()
	}
}

// ----- 业务逻辑层 -----
