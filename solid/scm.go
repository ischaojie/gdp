/*开闭原则*/
/*开闭原则讲的是如果软件实体发生变化，应该利用扩展，而不是修改已有的代码。*/
package solid

type IBook interface {
	GetName() string
	GetPrice() int
	GetAuthor() string
}

// NovelBook 小说
type NovelBook struct {
	Name   string
	Price  int
	Author string
}

func (n *NovelBook) GetName() string {
	return n.Name
}

func (n *NovelBook) GetPrice() int {
	return n.Price
}

func (n *NovelBook) GetAuthor() string {
	return n.Author
}

/*现在有新的需求，对所有小说打折，根据开闭原则，
打折相关的功能应该利用扩展实现，而不是在原有代码上修改，
所以，新增一个OffNovelBook接口，继承NovelBook，并重写GetPrice方法
*/

// OffNovelBook 组合(继承)了NovelBook
type OffNovelBook struct {
	NovelBook
}

// 重写GetPrice方法
func (n *OffNovelBook) GetPrice() int {
	beforePrice := n.NovelBook.GetPrice()
	offPrice := 0
	// 降价
	if beforePrice > 100 {
		offPrice = beforePrice * 90 / 100
	} else {
		offPrice = beforePrice * 80 / 100
	}
	return offPrice
}
