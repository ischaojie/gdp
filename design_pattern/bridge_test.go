package dp

import (
	"fmt"
	"testing"
)

/*我们可以组装（桥接）不同的形状和颜色*/

func TestShapeCircle(t *testing.T) {
	red := NewRed()
	blue := NewBlue()

	circleWithRed := NewCircle(red)
	circleWithBlue := NewCircle(blue)

	res1 := circleWithRed.Draw()
	res2 := circleWithBlue.Draw()
	fmt.Println(res1) // Draw circle shape with red color.
	fmt.Println(res2) // Draw circle shape with blue color.

}

func TestShapeSquare(t *testing.T) {
	red := NewRed()
	blue := NewBlue()

	SquareWithRed := NewSquare(red)
	SquareWithBlue := NewSquare(blue)

	res1 := SquareWithRed.Draw()
	res2 := SquareWithBlue.Draw()
	fmt.Println(res1) // Draw Square shape with red color.
	fmt.Println(res2) // Draw Square shape with blue color.

}
