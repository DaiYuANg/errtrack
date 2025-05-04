package entity

import (
	"github.com/uptrace/bun"
	"time"
)

// 事件主表（对应Sentry的Issue）
type Event struct {
	bun.BaseModel `bun:"table:sentry_events"`

	EventID     string                 `bun:"event_id,pk"`
	ProjectID   int                    `bun:"project_id"`
	Release     string                 `bun:"release"`
	Environment string                 `bun:"environment"`
	Message     string                 `bun:"message"`
	Exception   map[string]interface{} `bun:"type:jsonb"`
	Tags        map[string]interface{} `bun:"type:jsonb"`
	ReceivedAt  time.Time              `bun:"received_at,default:current_timestamp"`
}
