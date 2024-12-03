package eventrouter_test

import (
	"testing"

	"github.com/pablor21/eventrouter"
)

type MockEvent struct{}

func (e *MockEvent) Name() string {
	return "mockEvent"
}

type MockHook struct {
	eventrouter.Hook
	name string
}

func (h *MockHook) Name() string {
	return h.name
}

func TestRoute(t *testing.T) {
	handler := eventrouter.NewHandler("/test", nil)
	if handler.Route() != "/test" {
		t.Fatalf("Expected route to be /test, got %s", handler.Route())
	}
}

func TestHandle(t *testing.T) {
	handler := eventrouter.NewHandler("/test", func(e eventrouter.IEvent) error {
		return nil
	})
	event := &MockEvent{}
	if err := handler.Handle(event); err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
}

func TestHandleNilFunction(t *testing.T) {
	handler := &eventrouter.Handler{}
	event := &MockEvent{}
	if err := handler.Handle(event); err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
}

func TestHooks(t *testing.T) {
	hook := &MockHook{name: "testHook"}
	handler := eventrouter.NewHandler("/test", nil, hook)
	if len(handler.Hooks()) != 1 {
		t.Fatalf("Expected 1 hook, got %d", len(handler.Hooks()))
	}
}

func TestHanlderBind(t *testing.T) {
	hook := &MockHook{name: "testHook"}
	handler := &eventrouter.Handler{}
	handler.BindHook(hook)
	if len(handler.Hooks()) != 1 {
		t.Fatalf("Expected 1 hook, got %d", len(handler.Hooks()))
	}
}

func TestHandlerUnbind(t *testing.T) {
	hook := &MockHook{name: "testHook"}
	handler := &eventrouter.Handler{}
	handler.BindHook(hook)
	handler.UnbindHook(hook)
	if len(handler.Hooks()) != 0 {
		t.Fatalf("Expected 0 hooks, got %d", len(handler.Hooks()))
	}
}
