package event

import (
	"fmt"

	"github.com/modernice/goes/helper/pick"
)

type Handler interface {
	RegisterHandler(eventName string, handler func(Event))
}

// RegisterHandler registers an event handler for the given event name.
// The provided eventHandler should usually be an aggregate or projection that
// uses the registered handler to apply the events onto itself. The *Base
// aggregate implements Handler.
//
//	type Foo struct {
//		*aggregate.Base
//
//		Foo string
//		Bar string
//		Baz string
//	}
//
//	type FooEvent { Foo string }
//	type BarEvent { Bar string }
//	type BazEvent { Bar string }
//
//	func NewFoo(id uuid.UUID) *Foo  {
//		foo := &Foo{Base: aggregate.New("foo", id)}
//		aggregate.Register(foo, "foo", foo.foo)
//		aggregate.Register(foo, "bar", foo.bar)
//		aggregate.Register(foo, "baz", foo.baz)
//		return foo
//	}
//
//	func (f *Foo) foo(e event.EventOf[FooEvent]) {
//		f.Foo = e.Data().Foo
//	}
//
//	func (f *Foo) foo(e event.EventOf[BarEvent]) {
//		f.Bar = e.Data().Bar
//	}
//
//	func (f *Foo) foo(e event.EventOf[BazEvent]) {
//		f.Baz = e.Data().Baz
//	}
func RegisterHandler[D any](eh Handler, eventName string, handler func(Of[D])) {
	eh.RegisterHandler(eventName, func(evt Event) {
		if casted, ok := TryCast[D](evt); ok {
			handler(casted)
		} else {
			aggregateName := "<unknown>"
			if a, ok := eh.(pick.AggregateProvider); ok {
				aggregateName = pick.AggregateName(a)
			}
			panic(fmt.Errorf(
				"[goes/event.RegisterHandler] Cannot cast %T to %T. "+
					"You probably provided the wrong event name for this handler. "+
					"[event=%v, aggregate=%v]",
				evt, casted, eventName, aggregateName,
			))
		}
	})
}
