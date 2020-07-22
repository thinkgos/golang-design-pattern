package bridge

import "fmt"

// 消息接口 桥接抽象接口
type AbstractMessage interface {
	SendMessage(text, to string)
}

// 底层消息实现者 抽象接口
type MessageImplementer interface {
	Send(text, to string)
}

/************************ 抽象化实现者 ******************************/
// 实现SMS
type MessageSMS struct{}

func NewViaSMS() MessageImplementer {
	return &MessageSMS{}
}

func (*MessageSMS) Send(text, to string) {
	fmt.Printf("send %s to %s via SMS", text, to)
}

// 实现Email
type MessageEmail struct{}

func NewViaEmail() MessageImplementer {
	return &MessageEmail{}
}

func (*MessageEmail) Send(text, to string) {
	fmt.Printf("send %s to %s via Email", text, to)
}

/****************************  桥接实现者 ********************************/

// 桥接common
type CommonMessage struct {
	method MessageImplementer
}

func NewCommonMessage(method MessageImplementer) *CommonMessage {
	return &CommonMessage{method: method}
}

func (m *CommonMessage) SendMessage(text, to string) {
	m.method.Send(text, to)
}

// 桥接urgency
type UrgencyMessage struct {
	method MessageImplementer
}

func NewUrgencyMessage(method MessageImplementer) *UrgencyMessage {
	return &UrgencyMessage{method: method}
}

func (m *UrgencyMessage) SendMessage(text, to string) {
	m.method.Send(fmt.Sprintf("[Urgency] %s", text), to)
}
