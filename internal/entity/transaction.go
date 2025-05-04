package entity

import (
	"github.com/uptrace/bun"
	"time"
)

type Transaction struct {
	bun.BaseModel `bun:"table:sentry_transactions"`

	TraceID      string                 `bun:"trace_id"`
	SpanID       string                 `bun:"span_id"`
	ParentSpanID string                 `bun:"parent_span_id"`
	Op           string                 `bun:"op"`
	Description  string                 `bun:"description"`
	StartTime    time.Time              `bun:"start_time"`
	EndTime      time.Time              `bun:"end_time"`
	Tags         map[string]interface{} `bun:"tags,type:jsonb"`
}
