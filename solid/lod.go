/*迪米特原则（Law of Demeter）*/
/*
一个类需要对自己需要耦合或者调用的类知道的最少
*/

package solid

import "fmt"

// Teacher 老师
type Teacher struct {
}

func (t Teacher) Command(leader GroupLeader) {
	leader.countGirls()
}

type Girl struct {
}

type GroupLeader struct {
	Girls []Girl
}

func NewGroupLeader(girls []Girl) GroupLeader {
	return GroupLeader{
		Girls: girls,
	}
}

func (l GroupLeader) countGirls() {
	fmt.Printf("一共有%d个女生。", len(l.Girls))
}
