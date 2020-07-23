package memento

import "fmt"

// 备忘录, 保存对象的状态
type Memento struct {
	hp, mp int
}

// 状态
func (m *Memento) Status() string {
	return fmt.Sprintf("Current HP:%d, MP:%d", m.hp, m.mp)
}

// 对象
type Game struct {
	hp, mp int
}

func (g *Game) Play(mpDelta, hpDelta int) {
	g.mp += mpDelta
	g.hp += hpDelta
}

// 保存至备忘录
func (g *Game) Save() *Memento {
	return &Memento{
		hp: g.hp,
		mp: g.mp,
	}
}

// 从备忘录恢复
func (g *Game) Load(m *Memento) {
	g.mp = m.mp
	g.hp = m.hp
}

// 状态
func (g *Game) Status() string {
	return fmt.Sprintf("Current HP:%d, MP:%d", g.hp, g.mp)
}
