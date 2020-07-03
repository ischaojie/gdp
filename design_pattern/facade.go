package dp

import "fmt"

/*外观模式*/

// 外观角色
type Facade interface {
	method()
}

type facadeImpl struct {
	a ASubModel
	b BSubModel
}

func (f *facadeImpl) method() {
	f.a.methodA()
	f.b.methodB()
}

// --------------------------------子模块A-----------------------------//

type ASubModel interface {
	methodA()
}

type aSubModelImpl struct{}

func (aSubModelImpl) methodA() {
	fmt.Println("I'm A sub model.")
}

// ---------------------------------子模块B---------------------------- //

type BSubModel interface {
	methodB()
}

type bSubModelImpl struct{}

func (bSubModelImpl) methodB() {
	fmt.Println("I'm B sub model.")
}

// --------------------------------------------------------------- //

// 对外统一的接口
func NewFacade() Facade {
	return &facadeImpl{
		a: NewA(),
		b: NewB(),
	}
}

// 同时每个子模块也有独立对外的接口
func NewA() ASubModel {
	return &aSubModelImpl{}
}

func NewB() BSubModel {
	return &bSubModelImpl{}
}
