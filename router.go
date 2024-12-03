package eventrouter

import (
	"regexp"
)

type EventRouter struct {
	hooks    []IHook
	handlers map[string][]IHandler
	matchers map[*regexp.Regexp]string
	// lock     sync.RWMutex
}

func NewEventRouter() *EventRouter {
	return &EventRouter{
		hooks:    []IHook{},
		handlers: make(map[string][]IHandler),
		matchers: make(map[*regexp.Regexp]string),
		// lock:     sync.RWMutex{},
	}
}

func (eb *EventRouter) Handlers() map[string][]IHandler {
	return eb.handlers
}

func (eb *EventRouter) Bind(handler IHandler) error {
	route := handler.Route()

	if _, ok := eb.handlers[route]; !ok {
		eb.handlers[route] = []IHandler{}
	}

	// check if the handler is already registered
	for _, h := range eb.handlers[route] {
		if h == handler {
			return nil
		}
	}
	matcher, err := regexp.Compile(route)
	if err != nil {
		return err

	}

	eb.matchers[matcher] = route
	eb.handlers[route] = append(eb.handlers[route], handler)
	return nil
}

func (eb *EventRouter) BindFunc(route string, handlerFunc HandlerFunc, hooks ...IHook) error {
	return eb.Bind(NewHandler(route, handlerFunc, hooks...))
}

func (eb *EventRouter) Unbind(handler IHandler) *EventRouter {
	route := handler.Route()
	if _, ok := eb.handlers[route]; !ok {
		return eb
	}

	for i, h := range eb.handlers[route] {
		if h == handler {
			eb.handlers[route] = append(eb.handlers[route][:i], eb.handlers[route][i+1:]...)
			// if there are no more handlers for the route, remove the route
			if len(eb.handlers[route]) == 0 {
				delete(eb.handlers, route)
				for matcher, r := range eb.matchers {
					if r == route {
						delete(eb.matchers, matcher)
					}
				}
			}

			return eb
		}
	}

	return eb
}

func (eb *EventRouter) BindHook(hooks ...IHook) *EventRouter {
	eb.hooks = append(eb.hooks, hooks...)
	return eb
}

func (eb *EventRouter) UnbindHook(hook IHook) *EventRouter {
	for i, h := range eb.hooks {
		if h == hook {
			eb.hooks = append(eb.hooks[:i], eb.hooks[i+1:]...)
			return eb
		}
	}
	return eb
}

func (eb *EventRouter) BindHookFunc(hookFunc ...HookFunc) *EventRouter {
	for _, fn := range hookFunc {
		eb.BindHook(NewHook(fn))
	}
	return eb
}

func (eb *EventRouter) run(e IEvent, h IHandler) error {

	handlerFn := func(e IEvent) error {
		return h.Handle(e)
	}

	if h.Hooks() != nil {
		// iterate over the hooks in reverse order
		for i := len(h.Hooks()) - 1; i >= 0; i-- {
			handlerFn = h.Hooks()[i].Handle(handlerFn)
		}
	}

	// global hooks
	if eb.hooks != nil {
		// iterate over the hooks in reverse order
		for i := len(eb.hooks) - 1; i >= 0; i-- {
			handlerFn = eb.hooks[i].Handle(handlerFn)
		}
	}

	return handlerFn(e)
}

func (eb *EventRouter) Handle(e IEvent) error {
	// iterate over the handlers and check if the regex matches
	toRun := []IHandler{}

	for matcher, route := range eb.matchers {
		if matcher.MatchString(e.Name()) {
			toRun = append(toRun, eb.handlers[route]...)
		}
	}

	for _, handler := range toRun {
		err := eb.run(e, handler)
		if err != nil {
			return err
		}
	}

	return nil
}
