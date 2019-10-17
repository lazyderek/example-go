package prototype

import "testing"

func TestNewDog(t *testing.T) {
	d := NewDog()
	d.Run()

	c := NewDog()
	c.Run()
}
