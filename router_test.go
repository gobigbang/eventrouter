package eventrouter_test

import (
	"testing"

	"github.com/pablor21/eventrouter"
)

type MockHandler struct {
	eventrouter.Handler
	route string
}

func (m *MockHandler) Route() string {
	return m.route
}

func TestNewEventRouter(t *testing.T) {
	router := eventrouter.NewEventRouter()
	if router == nil {
		t.Fatal("Expected new router to be created")
	}
}

func TestBind(t *testing.T) {
	router := eventrouter.NewEventRouter()
	handler := &MockHandler{route: "/test"}

	err := router.Bind(handler)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if len(router.Handlers()["/test"]) != 1 {
		t.Fatalf("Expected 1 handler, got %d", len(router.Handlers()["/test"]))
	}
}

func TestBindDuplicate(t *testing.T) {
	router := eventrouter.NewEventRouter()
	handler := &MockHandler{route: "/test"}

	router.Bind(handler)
	err := router.Bind(handler)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if len(router.Handlers()["/test"]) != 1 {
		t.Fatalf("Expected 1 handler, got %d", len(router.Handlers()["/test"]))
	}
}

func TestBindInvalidRegex(t *testing.T) {
	router := eventrouter.NewEventRouter()
	handler := &MockHandler{route: "["}

	err := router.Bind(handler)
	if err == nil {
		t.Fatal("Expected error for invalid regex, got nil")
	}
}
