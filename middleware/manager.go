package middleware

import "net/http"

type Middleware func(http.Handler) http.Handler
type Manager struct {
	globalMiddlewares []Middleware
}

func (manager *Manager) With(n http.Handler, middlewares ...Middleware) http.Handler {

	next := n
	for _, middleware := range middlewares {
		next = middleware(next)
	}
	for _, middleware := range manager.globalMiddlewares {
		next = middleware(next)
	}
	return next

}

func (manager *Manager) Use(middlewares ...Middleware) {
	manager.globalMiddlewares = append(manager.globalMiddlewares, middlewares...)
}

func NewManager() *Manager {

	return &Manager{
		globalMiddlewares: make([]Middleware, 0),
	}
}
