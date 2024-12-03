package eventrouter_test

import (
	"testing"

	"github.com/gobigbang/eventrouter"
)

func TestNewEvent(t *testing.T) {
	event := eventrouter.NewEvent("testEvent")
	if event == nil {
		t.Fatal("Expected new event to be created")
	}
	if event.Name() != "testEvent" {
		t.Fatalf("Expected event name to be 'testEvent', got %s", event.Name())
	}
}

func TestEventName(t *testing.T) {
	event := eventrouter.NewEvent("")
	name := event.Name()
	if name == "" {
		t.Fatal("Expected event name to be generated, got empty string")
	}
	if len(name) != 10 {
		t.Fatalf("Expected event name length to be 10, got %d", len(name))
	}
}
