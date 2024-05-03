package strategy

import (
	"testing"
)

func TestStrategy(t *testing.T) {
	// ----- 业务逻辑层 -----

	hero := Hero{}

	// 使用小刀
	hero.SetStrategy(new(Knife))
	hero.Fight()

	// 使用 ak
	hero.SetStrategy(new(Ak47))
	hero.Fight()
}
