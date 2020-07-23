package mediator

import (
	"fmt"
)

// see https://blog.csdn.net/cloudUncle/article/details/83686683

// 中介者消息接口
type UnitedNations interface {
	ForwardMessage(message string, country Country)
}

// 中介者类,使实现中介者接口
type UnitedNationsSecurityCouncil struct {
	USA
	Iraq
}

func (unsc UnitedNationsSecurityCouncil) ForwardMessage(message string, country Country) {
	switch country.(type) {
	case USA:
		unsc.Iraq.GetMessage(message)
	case Iraq:
		unsc.USA.GetMessage(message)
	default:
		fmt.Println("The country is not a member of UNSC")
	}
}

// 同事类接口
type Country interface {
	SendMessage(message string)
	GetMessage(message string)
}

// 实现同事类接口
type USA struct {
	UnitedNations
}

func (usa USA) SendMessage(message string) {
	usa.UnitedNations.ForwardMessage(message, usa)
}
func (usa USA) GetMessage(message string) {
	fmt.Printf("美国收到对方消息: %s\n", message)
}

// 实现同事类接口
type Iraq struct {
	UnitedNations
}

func (iraq Iraq) SendMessage(message string) {
	iraq.UnitedNations.ForwardMessage(message, iraq)
}
func (iraq Iraq) GetMessage(message string) {
	fmt.Printf("伊拉克收到对方消息: %s\n", message)
}
