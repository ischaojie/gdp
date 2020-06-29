package dp

import (
	"fmt"
)

/*简单工厂模式*/

/*
简单工厂模式说的是，应该用一个专门的类或者方法来产生其他的类。
在例子中，只有Animal接口和NewAnimal对外可见，封装了实现的细节。
*/

// Animal 动物接口
type Animal interface {
	Say()
}

type Dog struct {
}

func (d Dog) Say() {
	fmt.Println("Hello, I'm a Dog.")
}

type Cat struct {
}

func (c Cat) Say() {
	fmt.Println("Hello, I'm a Cat.")
}

type AnimalType int

const (
	CatType AnimalType = 1 << iota
	DogType
)

/*这就是我们的工厂*/
// 工厂根据动物类型产生不同的动物
func NewAnimal(t AnimalType) Animal {
	switch t {
	case CatType:
		return Cat{}
	case DogType:
		return Dog{}
	default:
		return Cat{}
	}
}
