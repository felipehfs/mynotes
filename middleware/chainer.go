package middleware

import "net/http"

type Middleware func(http.Handler) http.Handler

// Adapt chains the middlewares
func Adapt(next http.Handler, handlers ...Middleware) http.Handler {
	for _, handle := range handlers {
		next = handle(next)
	}
	return next
}
