package strategy

// 策略接口
type Strategy interface {
	Do(*ComputerContext) int
}

// 策略上下文
type ComputerContext struct {
	Num1 int
	Num2 int
}

type Computer struct {
	context  *ComputerContext
	strategy Strategy
}

// 根据不同的策创建实例方法
func NewComputer(num1, num2 int, strategy Strategy) *Computer {
	return &Computer{
		&ComputerContext{
			num1,
			num2,
		},
		strategy,
	}
}

func (p *Computer) SetStrategy(strategy Strategy) {
	p.strategy = strategy
}

// 调策略
func (p *Computer) Do() int {
	return p.strategy.Do(p.context)
}

// 加法 策略实例
type Add struct{}

func (*Add) Do(ctx *ComputerContext) int {
	return ctx.Num1 + ctx.Num2
}

// 乘法 策略实例
type Mul struct{}

func (*Mul) Do(ctx *ComputerContext) int {
	return ctx.Num1 * ctx.Num2
}

// 减法 策略实例
type Sub struct{}

func (*Sub) Do(ctx *ComputerContext) int {
	return ctx.Num1 - ctx.Num2
}
