package logx

import (
	"context"

	"github.com/rizalgowandy/gdk/pkg/converter"
	"github.com/rizalgowandy/gdk/pkg/tags"
	"github.com/segmentio/ksuid"
	"google.golang.org/grpc/metadata"
)

// Key for standardized context value.
const (
	// 	RequestID is a random generated string to identify each request.
	//	Example: qwerty1234.
	RequestID = tags.RequestID
	UserID    = tags.ActorID
)

type contextKey string

const (
	CtxKeyRequestID = contextKey(RequestID)
	CtxKeyUserID    = contextKey(UserID)
)

// GenRequestID returns a unique request id.
func GenRequestID() string {
	return ksuid.New().String()
}

// NewContext returns a context with built-in request id.
func NewContext(in ...context.Context) context.Context {
	ctx := context.Background()
	if len(in) > 0 {
		ctx = in[0]
	}

	return SetRequestID(ctx, GenRequestID())
}

// ContextWithRequestID checks if current context has request id already and assign one if none found.
// Use for middleware and intercept all incoming requests to assign request id to the context if it doesn't exist.
func ContextWithRequestID(ctx context.Context) context.Context {
	if ctx == nil {
		ctx = ContextWithRequestID(context.Background())
	}

	// Check if it's already exists, then just return.
	if GetRequestID(ctx) != "" {
		return ctx
	}

	// Try to assign from metadata.
	ctx = SetRequestIDFromMetadata(ctx)

	// Check if it's already exists, then just return.
	if GetRequestID(ctx) != "" {
		return ctx
	}

	// If request id is still missing, create manually.
	return NewContext(ctx)
}

// SetRequestID returns a context with assigned request id.
func SetRequestID(ctx context.Context, id string) context.Context {
	if ctx == nil {
		return SetRequestID(context.Background(), id)
	}
	return context.WithValue(ctx, CtxKeyRequestID, id)
}

// GetContextID returns a request id assigned inside a context.
func GetRequestID(ctx context.Context) string {
	if ctx == nil {
		return ""
	}

	id, ok := ctx.Value(CtxKeyRequestID).(string)
	if !ok {
		return ""
	}
	return id
}

// SetRequestIDFromMetadata returns a context
// with request id and context id value if they exists in metadata.
func SetRequestIDFromMetadata(ctx context.Context) context.Context {
	if ctx == nil {
		return nil
	}

	// Get metadata from context.
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return ctx
	}

	// Get metadata request id.
	if requestID, ok := md[RequestID]; ok {
		if len(requestID) > 0 {
			ctx = SetRequestID(ctx, requestID[0])
		}
	}

	return ctx
}

func SetActorID(ctx context.Context, id int) context.Context {
	if ctx == nil {
		return SetActorID(context.Background(), id)
	}
	return context.WithValue(ctx, CtxKeyUserID, id)
}

func GetActorID(ctx context.Context) int {
	if ctx == nil {
		return 0
	}

	id, ok := ctx.Value(CtxKeyUserID).(int)
	if !ok {
		return 0
	}
	return id
}

func GetActorIDFromMetadata(ctx context.Context) int {
	if ctx == nil {
		return 0
	}

	// Get metadata from context.
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return 0
	}

	// Get metadata request id.
	if userID, ok := md[UserID]; ok {
		if len(userID) > 0 {
			return converter.Int(userID[0])
		}
	}

	return 0
}
