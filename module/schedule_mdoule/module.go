package schedule_mdoule

import (
	"context"
	"github.com/go-co-op/gocron/v2"
	"github.com/samber/lo"
	"go.uber.org/fx"
)

var Module = fx.Module("schedule_mdoule",
	fx.Provide(newScheduler),
	fx.Invoke(schedulerLifecycle),
)

func newScheduler() gocron.Scheduler {
	return lo.Must(
		gocron.NewScheduler(
			gocron.WithMonitor(newMonitor()),
			gocron.WithLogger(
				gocron.NewLogger(gocron.LogLevelDebug),
			),
		),
	)
}

func schedulerLifecycle(lc fx.Lifecycle, cron gocron.Scheduler) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go cron.Start()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return cron.Shutdown()
		},
	})
}
