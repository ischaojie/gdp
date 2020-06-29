package dp

// Builder 代表被建造的对象
// 以汽车工厂为例，Builder代表汽车的组成
type Builder interface {
	Part1()
	Part2()
	Part3()
}

// Builder1 建造对象1
type Builder1 struct {
	result string
}

// Part1 Builder1 部分1
func (b *Builder1) Part1() {
	b.result += "1"
}

// Part2 Buider1 部分2
func (b *Builder1) Part2() {
	b.result += "2"
}

// Part3 buidler1
func (b *Builder1) Part3() {
	b.result += "3"
}

// GetResult Builder1
func (b *Builder1) GetResult() string {
	return b.result
}

// Builder2 建造对象2
type Builder2 struct {
	result int
}

// Part1 Builder2 部分1
func (b *Builder2) Part1() {
	b.result += 1
}

// Part2 Buider2 部分2
func (b *Builder2) Part2() {
	b.result += 2
}

// Part3 buidler2
func (b *Builder2) Part3() {
	b.result += 3
}

// GetResult Builder2
func (b *Builder2) GetResult() int {
	return b.result
}

// Director 代表建设者
type Director struct {
	builder Builder
}

// NewDirector 新建director
func NewDirector(builder Builder) *Director {
	return &Director{
		builder: builder,
	}
}

// Construct 组装
func (d *Director) Construct() {
	d.builder.Part1()
	d.builder.Part2()
	d.builder.Part3()
}
