package prototype

// Cloneable 是原型对象需要实现的接口
type Cloneable interface {
	Name() string
	Clone() Cloneable
}

type PrototypeManager struct {
	prototypes map[string]Cloneable
}

func NewPrototypeManager() *PrototypeManager {
	return &PrototypeManager{
		prototypes: make(map[string]Cloneable),
	}
}

func (p *PrototypeManager) Get(name string) Cloneable {
	return p.prototypes[name]
}

func (p *PrototypeManager) Set(name string, prototype Cloneable) {
	p.prototypes[name] = prototype
}

//**************************** 对象 ********************************************

type Type1 struct {
	name string
}

func (t *Type1) Clone() Cloneable {
	tc := *t // 复制一个现成对象(浅拷贝),比new完后再赋值要快?
	return &tc
}

func (t *Type1) Name() string {
	return t.name
}
