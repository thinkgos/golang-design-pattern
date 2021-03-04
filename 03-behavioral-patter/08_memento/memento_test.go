package memento

import (
	"fmt"
)

func ExampleGame() {
	game := &Game{
		hp: 10,
		mp: 10,
	}

	fmt.Println(game.Status())
	progress := game.Save() // 保存至备忘录

	game.Play(-2, -3)
	fmt.Println(game.Status())

	game.Load(progress) // 从备忘录恢复
	fmt.Println(game.Status())

	// Output:
	// Current HP:10, MP:10
	// Current HP:7, MP:8
	// Current HP:10, MP:10
}
