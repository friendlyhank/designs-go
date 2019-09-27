package main

//抽象目标
type Subject interface {
	Attach(Observer) //注册观察者
	Detach(Observer) //释放观察者
	Notify()         //通知所有注册的观察者
}

//抽象观察者
type Observer interface {
	Update(Subject) //观察者进行更新状态
}

//implements Subject
type ConcreteSubject struct {
	observers []Observer
	value     int
}

func NewConcreteSubject() *ConcreteSubject {
	s := new(ConcreteSubject)
	s.observers = make([]Observer,0)
	return s
}

func (s *ConcreteSubject) Attach(observe Observer) { //注册观察者
	s.observers = append(s.observers,observe)
}

func (s *ConcreteSubject) Detach(observer Observer) { //释放观察者

	for i,ob :=range s.observers{
		if ob == observer{
			s.observers = append(s.observers[:i],s.observers[i+1:]...)
		}
	}
}

func (s *ConcreteSubject) Notify() { //通知所有观察者
	for _,ob :=range s.observers{
		ob.Update(s)
	}
}

func (s *ConcreteSubject) setValue(value int) {
	s.value = value
	s.Notify()
}

func (s *ConcreteSubject) getValue() int {
	return s.value
}

/**
 * 具体观察者 implements Observer
 *
 */
type ConcreteObserver1 struct {
}

func (c *ConcreteObserver1) Update(subject Subject) {
	println("ConcreteObserver1  value is ", subject.(*ConcreteSubject).getValue())
}

/**
 * 具体观察者 implements Observer
 *
 */
type ConcreteObserver2 struct {
}

func (c *ConcreteObserver2) Update(subject Subject) {
	println("ConcreteObserver2 value is ", subject.(*ConcreteSubject).getValue())
}

func main() {

	subject := NewConcreteSubject()
	observer1 := new(ConcreteObserver1)
	observer2 := new(ConcreteObserver2)
	subject.Attach(observer1)
	subject.Attach(observer2)
	//subject.Detach(observer1)
	subject.setValue(5)

}