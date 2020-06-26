/*里氏替换原则*/
/*
意思是只要父类出现的地方，子类就可以出现，反之则不行。
1. 子类必须完全实现父类的方法

*/
package solid

import "fmt"

type IGun interface {
	shoot()
}

// 手枪
type HandGun struct{}

func (HandGun) shoot() {
	fmt.Printf("手枪射击...")
}

// 步枪
type Rifle struct{}

func (Rifle) shoot() {
	fmt.Printf("步枪射击...")
}

// 玩具枪不能杀人，所以单独新建一个接口
// type ToyGun struct {
// }

type IToyGun interface {
	IGun
}

type ToyGun struct {
}

func (ToyGun) shoot() {
	fmt.Println("玩具枪兹水...")
}

// 士兵
type Soldier struct {
	Gun IGun // 发一把抽象枪
}

func (s *Soldier) KillEnemy() {
	fmt.Println("我要射击了：")
	s.Gun.shoot()
}
