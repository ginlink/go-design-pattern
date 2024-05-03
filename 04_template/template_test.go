package template

import "testing"

func TestTemplate(t *testing.T) {
	makeTea := new(MakeTea)
	makeTea.b = makeTea
	makeTea.MakeBeverage()

	makeCoffee := new(MakeCoffee)
	makeCoffee.b = makeCoffee
	makeCoffee.MakeBeverage()
}
