package command

import "testing"

// ----- 业务逻辑层 -----
func TestCommand(t *testing.T) {
	doctor := Doctor{}
	eyeCmd := EyeCommand{doctor: &doctor}   // 治疗眼睛的命令
	noseCmd := NoseCommand{doctor: &doctor} // 治疗鼻子的命令

	// 护士
	nurse := Nurse{}
	nurse.CmdList = append(nurse.CmdList, &eyeCmd)
	nurse.CmdList = append(nurse.CmdList, &noseCmd)

	// 护士通知医生
	nurse.Notify()
}
