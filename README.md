[![GoDoc](https://godoc.org/github.com/gobigbang/eventrouter?status.svg)](https://godoc.org/github.com/gobigbang/eventrouter)
[![GitHub release](https://img.shields.io/github/release/gobigbang/eventrouter.svg)](https://img.shields.io/github/release/gobigbang/eventrouter.svg)
[![GitHub license](https://img.shields.io/badge/license-MIT-blue.svg)](https://raw.githubusercontent.com/gobigbang/eventrouter/master/LICENSE)


# EventRouter

EventRouter is a simple event routing library for Go. It allows you to bind handlers to routes and hooks to events, making it easy to manage and handle events in your application.

## Installation

To install EventRouter, use go get:

```sh
go get github.com/gobigbang/eventrouter
```

## Usage

### Creating an EventRouter

To create a new EventRouter, use the NewEventRouter function:

```go
router := eventrouter.NewEventRouter()
```

### Binding Handlers

You can bind handlers to routes using the Bind method:

```go
handler := eventrouter.NewHandler("/test", func(e eventrouter.IEvent) error {
    return nil
})
router.Bind(handler)
```

### Unbinding Handlers

To unbind a handler from a route, use the Unbind method:

```go
router.Unbind(handler)
```

### Binding Hooks

You can bind hooks to handlers using the BindHook method:

```go
hook:= eventrouter.NewHook(func(next eventrouter.HandlerFunc) eventrouter.HandlerFunc {
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
handler.BindHook(hook)

// you can also hook to the whole router
router.BindHook(hook)
```

### Unbinding Hooks

To unbind a hook from an event, use the UnbindHook method:

```go 
router.UnbindHook(hook) 
```

### Handling Events

To handle an event, use the Handle method:

```go
e := eventrouter.NewEvent("/test.1")
router.Handle(e)
```

## Testing

To run the tests for this project, use the go test command:

```sh 
go test ./... 
```

## License

This project is licensed under the MIT License. See the LICENSE file for details.
