/*接口隔离原则:
意思是接口应该尽量细化，接口中的方法尽可能少。
*/
package solid

import "fmt"

// 星探接口
type ISearcher interface {
	show()
}

type Searcher struct {
	PrettyGirl PrettyGirl
}

func (s Searcher) show() {
	s.PrettyGirl.GoodLooking()
	s.PrettyGirl.NiceFigure()
	s.PrettyGirl.GreatTemperament()
}

// 美女接口
type IGoodBody interface {
	GoodLooking()
	NiceFigure()
}

type IGreateTemperament interface {
	GreatTemperament()
}

// 美女实现
type PrettyGirl struct {
	name string
}

func (g *PrettyGirl) GoodLooking() {
	fmt.Printf("%s 脸蛋很漂亮", g.name)
}

func (g *PrettyGirl) NiceFigure() {
	fmt.Printf("%s 身材很好", g.name)
}

func (g *PrettyGirl) GreatTemperament() {
	fmt.Printf("%s 气质很好", g.name)
}
