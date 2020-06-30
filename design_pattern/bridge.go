package dp

import "fmt"

/*桥接模式*/

/*Shape 形状*/

// Shape 形状接口
type Shape interface {
	Draw() string
}

// ShapeCircle 圆形
type ShapeCircle struct {
	color Color
}

func (s *ShapeCircle) Draw() string {
	return s.color.DrawColor(fmt.Sprintf("circle"))
}

// ShapeSquare 方形
type ShapeSquare struct {
	color Color
}

func (s *ShapeSquare) Draw() string {
	return s.color.DrawColor(fmt.Sprintf("square"))
}

/*Color 颜色接口*/

// Color 颜色接口
type Color interface {
	DrawColor(shape string) string
}

// ColorRed 红色
type ColorRed struct{}

func (c *ColorRed) DrawColor(shape string) string {
	return fmt.Sprintf("Draw %s shape with red color.", shape)
}

// ColorBlue 蓝色
type ColorBlue struct{}

func (c *ColorBlue) DrawColor(shape string) string {
	return fmt.Sprintf("Draw %s shape with blue color.", shape)
}

// 工厂函数

func NewCircle(color Color) *ShapeCircle {
	return &ShapeCircle{color: color}
}

func NewSquare(color Color) *ShapeSquare {
	return &ShapeSquare{color: color}
}

func NewRed() Color {
	return &ColorRed{}
}

func NewBlue() Color {
	return &ColorBlue{}
}
