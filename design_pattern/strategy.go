package dp

import "fmt"

// SortStrategy 排序策略
// 可用多种算法实现该接口
type SortStrategy interface {
	Sort(*SortCtx)
}

// SortCtx 排序上下文，定义sort用到的数据等
type SortCtx struct {
	arr string // 待排序数据
}

// -------------------冒泡排序&快排（SortStrategy接口实现）----------------------------//

type BubbleSort struct {
}

func (s *BubbleSort) Sort(ctx *SortCtx) {
	fmt.Printf("Sort %s use bubble sort.", ctx.arr)
}

type QuickSort struct{}

func (s *QuickSort) Sort(ctx *SortCtx) {
	fmt.Printf("Sort %s use quick sort.", ctx.arr)
}

// --------------------sort client--------------------------------//
type Sort struct {
	strategy SortStrategy
	ctx      *SortCtx
}

func (s *Sort) Sort() {
	s.strategy.Sort(s.ctx)
}

// NewSort api开放给外界
func NewSort(arr string, strategy SortStrategy) *Sort {
	return &Sort{
		strategy: strategy,
		ctx:      &SortCtx{arr: arr},
	}
}
