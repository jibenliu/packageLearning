package workflowPattern

import (
	"fmt"
)

type observer interface {
	update(state string)
}

// TextObserver is ConcreteObserver
type TextObserver struct {
	_name string
}

func (t TextObserver) update(state string) {
	fmt.Println(t._name + ": " + state)
}

// TestSubject is Subject
type TestSubject struct {
	_observers []observer
}

// Attach adds an observer
func (ts *TestSubject) Attach(observer observer) {
	ts._observers = append(ts._observers, observer)
}

// Detach removes an observer
func (ts *TestSubject) Detach(observer observer) {
	index := 0
	for i := range ts._observers {
		if ts._observers[i] == observer {
			index = i
			break
		}
	}
	ts._observers = append(ts._observers[0:index], ts._observers[index+1:]...)
}

func (ts TestSubject) notify(state string) {
	for _, observer := range ts._observers {
		observer.update(state)
	}
}

// TextEdit is ConcreteSubject
type TextEdit struct {
	TestSubject
	Text string
}

// SetText changes the text and informs observers
func (te TextEdit) SetText(text string) {
	te.Text = text
	te.TestSubject.notify(text)
}

func ObserverDemo() {
	//Client
	observer1 := TextObserver{"IObserver #1"}
	observer2 := TextObserver{"IObserver #2"}

	textEdit := TextEdit{}
	textEdit.Attach(observer1)
	textEdit.Attach(observer2)

	textEdit.SetText("test text")
	//printed:
	//IObserver #1: test text
	//IObserver #2: test text
}
