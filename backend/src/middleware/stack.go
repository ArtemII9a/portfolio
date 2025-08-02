package middleware

import (
	"net/http"
)

// How the heck they even chained/timed in the proper way?
func NewMiddlewareStack(ms ...func(http.Handler) http.Handler) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		for i := len(ms) - 1; i >= 0; i-- {
			m := ms[i]
			next = m(next)
		}
		return next
	}
}
