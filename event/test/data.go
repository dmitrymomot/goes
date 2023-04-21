package test

import (
	"github.com/modernice/goes/codec"
)

// FooEventData is a testing event data.
type FooEventData struct{ A string }

// BarEventData is a testing event data.
type BarEventData struct{ A string }

// BazEventData is a testing event data.
type BazEventData struct{ A string }

// FoobarEventData is a testing event data.
type FoobarEventData struct{ A int }

// UnregisteredEventData is a testing event data that's not registered in the
// Encoder returned by NewEncoder.
type UnregisteredEventData struct{ A string }

// NewEncoder returns a *codec.Registry that can be used to encode and decode
// events. The Encoder returned by NewEncoder is pre-registered with the types
// FooEventData, BarEventData, BazEventData, and FoobarEventData.
func NewEncoder() *codec.Registry {
	r := codec.New()
	codec.Register[FooEventData](r, "foo")
	codec.Register[BarEventData](r, "bar")
	codec.Register[BazEventData](r, "baz")
	codec.Register[FoobarEventData](r, "foobar")
	return r
}
