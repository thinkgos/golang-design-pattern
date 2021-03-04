package observer

import "fmt"

type Subject interface {
	Attach(Observer) // 注册观察者
	Detach(Observer) // 释放观察者
	Notify()         // 通知所有注册的观察者
}

// Observer 观察者
type Observer interface {
	Update(Subject)
}

// 被观察者
type ConcreteSubject struct {
	observers []Observer
	context   string
}

func NewConcreteSubject() *ConcreteSubject {
	return &ConcreteSubject{
		observers: make([]Observer, 0),
	}
}

func (s *ConcreteSubject) Attach(o Observer) {
	s.observers = append(s.observers, o)
}

func (s *ConcreteSubject) Detach(o Observer) {
	for i := 0; i < len(s.observers); i++ {
		if s.observers[i] == o {
			s.observers = append(s.observers[:i], s.observers[i+1:]...)
			return
		}
	}
}

func (s *ConcreteSubject) Notify() {
	for _, o := range s.observers {
		o.Update(s)
	}
}

func (s *ConcreteSubject) UpdateContext(context string) {
	s.context = context
	s.Notify()
}

// 观察者实例
type Reader struct {
	name string
}

func NewReader(name string) *Reader {
	return &Reader{name}
}

func (r *Reader) Update(s Subject) {
	fmt.Printf("%s receive %s\n", r.name, s.(*ConcreteSubject).context)
}
