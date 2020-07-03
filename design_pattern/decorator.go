package dp

/*装饰模式*/

// 利用装饰模式实现加减乘除计算

type Component interface {
	Calc() int
}

type ConcreteComp struct{}

func (c *ConcreteComp) Calc() int {
	return 0
}

// MulDecorator 乘
type MulDecorator struct {
	Component
	num int
}

func (d *MulDecorator) Calc() int {
	return d.Component.Calc() * d.num
}

// AddDecorator 加
type AddDecorator struct {
	Component
	num int
}

func (d *AddDecorator) Calc() int {
	return d.Component.Calc() + d.num
}

// 工厂函数
func NewAdd(c Component, num int) Component {
	return &AddDecorator{
		Component: c,
		num:       num,
	}
}

func NewMul(c Component, num int) Component {
	return &MulDecorator{
		Component: c,
		num:       num,
	}
}
