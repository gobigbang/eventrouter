package main

import (
	"github.com/gobigbang/eventrouter"
)

type testHandler struct {
	eventrouter.Handler
}

func (h *testHandler) Route() string {
	return "/test.*"
}

func (h *testHandler) Handle(e eventrouter.IEvent) error {
	println("handler called")
	return nil
}

type testHook struct {
	name string
}

func newTestHook(name string) *testHook {
	return &testHook{name: name}
}

func (h *testHook) Handle(next eventrouter.HandlerFunc) eventrouter.HandlerFunc {
	return func(e eventrouter.IEvent) error {
		println("pre run called", h.name)
		err := next(e)
		if err != nil {
			return err
		}
		println("hook called", h.name)
		return nil
	}
}

func main() {
	eb := eventrouter.NewEventRouter()

	h := &testHandler{}
	hook := eventrouter.NewHook(func(next eventrouter.HandlerFunc) eventrouter.HandlerFunc {
		return func(e eventrouter.IEvent) error {
			println("pre run called")
			err := next(e)
			if err != nil {
				return err
			}
			println("hook called")
			return nil
		}
	})
	h.BindHook(newTestHook("test"), hook)

	eb.Bind(h)

	e := eventrouter.NewEvent("/test.1")

	eb.Handle(e)
}
