/* 单例模式 */

/* 单例模式满足的是某些类在整个系统中只存在
	一个实例，比如数据库，或者唯一序列号的生成。
*/

package dp

import "sync"

type Singleton struct {

}

var (
	instance *Singleton
	once sync.Once
	mux sync.Mutex
)

// Go语言中可以使用Once实现
// Once保证只执行一次

func Instance() *Singleton{
	once.Do(func ()  {
		instance = &Singleton{}
	})
	return instance
}

// 不使用Once实现
// 一种常见的说法叫做懒汉模式
// 懒汉模式指的是假设对象已经创建，如果未创建，通过双重加锁创建对象
func Instace2() *Singleton{
	if instance == nil{
		mux.Lock()
		defer mux.Unlock()
		if instance == nil {
			instance = &Singleton{}
		}
	}
	return instance
}

// 另一种写法叫做饿汉模式
// 饿汉模式的做法是预先初始化
// go语言中直接在init初始化
func init() {
	instance = new(Singleton)
}

func GetInstance() *Singleton {
	return instance
}
