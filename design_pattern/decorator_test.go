package dp

import (
	"fmt"
	"testing"
)

func TestDecorator(t *testing.T) {

	var c Component = &ConcreteComp{}

	c = NewAdd(c, 100)
	c = NewMul(c, 2)

	res := c.Calc()
	fmt.Printf("res is : %d", res)

}
