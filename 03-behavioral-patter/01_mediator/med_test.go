package mediator

func ExampleUnitedNationsSecurityCouncil_ForwardMessage() {
	//创建一个具体中介者
	tMediator := UnitedNationsSecurityCouncil{}

	//创建具体同事,并且让他认识中介者
	tColleageA := USA{tMediator}
	tColleageB := Iraq{tMediator}

	//让中介者认识每一个具体同事
	tMediator.USA = tColleageA
	tMediator.Iraq = tColleageB

	//A同事发送消息
	tColleageA.SendMessage("停止核武器研发，否则发动战争")
	tColleageB.SendMessage("我们没有研发核武器，也不怕战争")

	// Output:
	// 伊拉克收到对方消息: 停止核武器研发，否则发动战争
	// 美国收到对方消息: 我们没有研发核武器，也不怕战争
}
