package adapter

import "testing"

func TestAdapter(t *testing.T) {
	phone := Phone{v: &Adapter{v220: new(V220)}}
	phone.Charge()
}
