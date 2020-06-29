package dp

import (
	"sync"
	"testing"
)

func TestInstance(t *testing.T) {
	ins1 := Instance()
	ins2 := Instance()
	if ins1 != ins2 {
		t.Fatal("instance not equal")
	}
}

const insCount = 100

// 测一下创建多个实例，是否都相等
func TestMultiInstance(t *testing.T) {
	wg := sync.WaitGroup{}
	wg.Add(insCount)

	inss := [insCount]*Singleton{}
	for i := 0; i < insCount; i++ {
		go func(item int) {
			inss[item] = Instance()
			wg.Done()
		}(i)
	}
	wg.Wait()

	for i := 1; i < insCount; i++ {
		if inss[i] != inss[i-1] {
			t.Fatal("there more instances in system")
		}
	}

}
