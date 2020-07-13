package observer

import "fmt"

// Observer 观察者
type Observer interface {
	Update(*Subject)
}

// 被观察者
type Subject struct {
	observers []Observer
	context   string
}

func NewSubject() *Subject {
	return &Subject{
		observers: make([]Observer, 0),
	}
}

func (s *Subject) UpdateContext(context string) {
	s.context = context
	s.notify()
}

func (s *Subject) Attach(o Observer) {
	s.observers = append(s.observers, o)
}

func (s *Subject) notify() {
	for _, o := range s.observers {
		o.Update(s)
	}
}

// 观察者实例
type Reader struct {
	name string
}

func NewReader(name string) *Reader {
	return &Reader{name}
}

func (r *Reader) Update(s *Subject) {
	fmt.Printf("%s receive %s\n", r.name, s.context)
}
