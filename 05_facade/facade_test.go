package facade

import (
	"fmt"
	"testing"
)

func TestFacade(t *testing.T) {
	hpf := new(HomePlayerFacade)

	hpf.DoKTV()
	fmt.Println("-----")
	hpf.DoLight()
}
