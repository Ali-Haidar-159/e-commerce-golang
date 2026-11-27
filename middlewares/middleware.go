package middlewares

import "net/http"

type MiddlewareType func(http.Handler) http.Handler

type Middleware struct {
	MiddlewareList []MiddlewareType
}

// Execution order when request arrives: allMiddlewares -> MiddlewareList -> next
func (m *Middleware) With(next http.Handler, extraMiddlewares ...MiddlewareType) http.Handler {
	n := next

	// for i := len(extraMiddlewares) - 1; i >= 0; i-- {
	// 	n = extraMiddlewares[i](n)
	// }

	for _, item := range extraMiddlewares {
		n = item(n)
	}

	// for i := len(m.MiddlewareList) - 1; i >= 0; i-- {
	// 	n = m.MiddlewareList[i](n)
	// }

	for _, item := range m.MiddlewareList {
		n = item(n)
	}

	return n
}

func (m *Middleware) Use(middlewares ...MiddlewareType) *Middleware {
	m.MiddlewareList = append(m.MiddlewareList, middlewares...)
	return m
}

func NewMiddleware() *Middleware {
	return &Middleware{
		MiddlewareList: make([]MiddlewareType, 0),
	}
}
