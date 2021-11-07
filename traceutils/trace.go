package traceutils

import (
	"context"
	"fmt"
	"time"

	uuid "github.com/satori/go.uuid"
)

const (
	// <timestamp>:<uuid_v4>
	traceIdTemplate = "%d:%s"
	TraceIdKey      = "trace-id"
)

func GenerateTraceId() string {
	u, err := uuid.NewV4()
	if err != nil {
		// try yet
		u, err = uuid.NewV4()
		if err != nil {
			panic(fmt.Errorf("error generate traceId %w", err))
		}
	}
	t := time.Now().Unix()
	return fmt.Sprintf(traceIdTemplate, t, u)
}

// FromCtx return traceId from context is exist
func FromCtx(ctx context.Context) *string {
	if traceId, ok := ctx.Value(TraceIdKey).(string); ok {
		return &traceId
	}
	return nil
}
