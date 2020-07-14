package bridge

func ExampleCommonSMS() {
	m := NewCommonMessage(NewViaSMS())
	m.SendMessage("have a drink?", "bob")
	// Output:
	// send have a drink? to bob via SMS
}

func ExampleCommonEmail() {
	m := NewCommonMessage(NewViaEmail())
	m.SendMessage("have a drink?", "bob")
	// Output:
	// send have a drink? to bob via Email
}

func ExampleUrgencySMS() {
	m := NewUrgencyMessage(NewViaSMS())
	m.SendMessage("have a drink?", "bob")
	// Output:
	// send [Urgency] have a drink? to bob via SMS
}

func ExampleUrgencyEmail() {
	m := NewUrgencyMessage(NewViaEmail())
	m.SendMessage("have a drink?", "bob")
	// Output:
	// send [Urgency] have a drink? to bob via Email
}
