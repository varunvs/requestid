package requestid

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"net/http"

	http2 "github.com/go-kit/kit/transport/http"
)

const requestIDHeader = "X-Request-Id"
const requestIDKey = "requestID"

func generateShortID() string {
	b := make([]byte, 6)
	rand.Read(b)

	return hex.EncodeToString(b)
}

func RequestIDToContext() http2.RequestFunc {
	return func(ctx context.Context, r *http.Request) context.Context {
		var id string

		if r.Header.Get(requestIDHeader) == "" {
			id = generateShortID()
		}

		return context.WithValue(ctx, requestIDKey, id)
	}
}

func RequestIDToResponse() http2.ServerResponseFunc {
	return func(ctx context.Context, writer http.ResponseWriter) context.Context {
		id := ctx.Value(requestIDKey)
		writer.Header().Set(requestIDHeader, id.(string))

		return ctx
	}
}
