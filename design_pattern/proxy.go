package dp

type ISubject interface {
	Do() string
}

// 实际的对象
type Subject struct {
}

func (Subject) Do() string {
	return "do real"
}

// Proxy 代理对象
type Proxy struct {
	real Subject
}

func (p Proxy) Do() string {
	var res string

	// 调用之前可以检查缓存、权限啊等等
	res += "before:"

	// 代理真实对象
	res += p.real.Do()

	// 调用后的操作
	res += ":after"

	return res

}
