package contextbuilder

import "context"

type ContextKey string

// NewContextWithRequestID returns a new Context that carries value u.
func NewContextWithRequestID(ctx context.Context, requestID string) context.Context {
	return context.WithValue(ctx, requestIDContextKey(), requestID)
}

func requestIDContextKey() ContextKey {
	return ContextKey("request_id")
}

// RequestIDFromContext returns the requestID value stored in ctx, if any.
func RequestIDFromContext(ctx context.Context) (string, bool) {
	u, ok := ctx.Value(requestIDContextKey()).(string)

	return u, ok
}

// CustomValues returns a list of all custom values, if any.
func CustomValues(ctx context.Context) map[string]string {
	m := map[string]string{}
	rid, ok := ctx.Value(requestIDContextKey()).(string)

	if ok {
		m["request_id"] = rid
	}

	return m
}
