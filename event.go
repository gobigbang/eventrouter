package eventrouter

type IEvent interface {
	Name() string
}

type Event struct {
	name string
}

func NewEvent(name string) *Event {
	return &Event{
		name: name,
	}
}

func (e *Event) Name() string {
	if e.name == "" {
		e.name = RandomString(10)
	}
	return e.name
}
