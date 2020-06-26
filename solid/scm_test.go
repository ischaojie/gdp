package solid

import (
	"fmt"
	"testing"
)

func TestSCM(t *testing.T) {
	// 打折后
	book := OffNovelBook{NovelBook{
		Name:   "射雕英雄传",
		Price:  200,
		Author: "金庸",
	}}

	fmt.Printf("name:%s, Price:%d, Author:%s", book.GetName(), book.GetPrice(), book.GetAuthor())

}
