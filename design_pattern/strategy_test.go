package dp

import "testing"

func TestSortStrategy(t *testing.T) {

	arr := "[3,2,6,4,8,1]"

	// 选择冒泡排序策略
	strategy := &BubbleSort{}

	sort := NewSort(arr, strategy)

	sort.Sort()
	// Sort [3,2,6,4,8,1] use bubble sort.

}
