package dp

/*Adaptee 适配者*/

// Adaptee 适配者接口
type Adaptee interface {
	SpecificRequest() string
}

// adapteeImpl 适配者接口具体实现
type adapteeImpl struct{}

func (a *adapteeImpl) SpecificRequest() string {
	return "I'm Adaptee method"
}

/*Target 目标*/

// Target 目标接口
type Target interface {
	Request() string
}

// adapter 转换Adaptee到Target
// adapter 是Target接口的具体实现
type adapter struct {
	Adaptee
}

func (a *adapter) Request() string {
	// 原本adaptee的方法被转换到target
	return a.SpecificRequest()
}

// NewAdapter 工厂函数，对外暴露adapter
func NewAdapter(adaptee Adaptee) Target {
	return &adapter{adaptee}
}
