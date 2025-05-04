package schedule_mdoule

import (
	"github.com/go-co-op/gocron/v2"
	"github.com/google/uuid"
	"time"
)

type monitor struct{}

func (monitor) IncrementJob(id uuid.UUID, name string, tags []string, status gocron.JobStatus) {

}

func (monitor) RecordJobTiming(startTime, endTime time.Time, id uuid.UUID, name string, tags []string) {
}

func newMonitor() *monitor {
	return &monitor{}
}
