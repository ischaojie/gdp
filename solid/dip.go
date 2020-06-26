/*依赖倒置原则（Dependence Inversion Principle）DIP*/

/*
依赖倒置原则的本质是要求面向接口编程，
也就是模块之间的依赖通过抽象来发生，具体的实现间没有依赖关系
接口不应该依赖实现
同样的实现也不应该依赖接口
*/
package solid

import "fmt"

// IDriver 司机抽象接口
type IDriver interface {
	// 可以看到，我们在接口中实现依赖
	// 抽象司机依赖抽象的汽车
	// 这样司机就可以开不同的汽车
	Drive(car ICar) // 开车
}

// Driver 具体的司机
type Driver struct{}

// Drive 实现IDriver接口
func (Driver) Drive(car ICar) {
	car.Run()
}

// ICar 汽车抽象接口
type ICar interface {
	Run()
}

// 不同汽车的实现

// Benz 奔驰汽车
type Benz struct{}

func (Benz) Run() {
	fmt.Println("奔驰汽车开始运行...")
}

// BMW 宝马汽车
type BMW struct{}

func (BMW) Run() {
	fmt.Println("宝马汽车开始运行...")
}
