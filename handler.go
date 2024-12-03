package eventrouter

type HandlerFunc func(e IEvent) error
type HookFunc func(next HandlerFunc) HandlerFunc

type IHook interface {
	Handle(next HandlerFunc) HandlerFunc
}

type Hook struct {
	fn HookFunc
}

func NewHook(fn HookFunc) *Hook {
	return &Hook{
		fn: fn,
	}
}

func (h *Hook) Handle(next HandlerFunc) HandlerFunc {
	return h.fn(next)
}

type IHandler interface {
	Route() string
	Handle(e IEvent) error
	Hooks() []IHook
	BindHook(hook ...IHook) IHandler
	UnbindHook(hook IHook) IHandler
	BindHookFunc(hook ...HookFunc) IHandler
}

type Handler struct {
	route string
	fn    HandlerFunc
	hooks []IHook
}

func NewHandler(route string, fn HandlerFunc, hooks ...IHook) *Handler {
	return &Handler{
		route: route,
		fn:    fn,
		hooks: hooks,
	}
}

func (h *Handler) Route() string {
	return h.route
}

func (h *Handler) Handle(e IEvent) error {
	if h.fn == nil {
		return nil
	}
	return h.fn(e)
}

func (h *Handler) Hooks() []IHook {
	return h.hooks
}

func (h *Handler) BindHook(hook ...IHook) IHandler {
	h.hooks = append(h.hooks, hook...)
	return h
}

func (h *Handler) UnbindHook(hook IHook) IHandler {
	for i, ho := range h.hooks {
		if ho == hook {
			h.hooks = append(h.hooks[:i], h.hooks[i+1:]...)
			return h
		}
	}
	return h
}

func (h *Handler) BindHookFunc(hook ...HookFunc) IHandler {
	for _, fn := range hook {
		h.hooks = append(h.hooks, NewHook(fn))
	}
	return h
}
