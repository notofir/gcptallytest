package main

import (
	"os"
	"time"

	"github.com/notofir/gcptallytest/reporter"
	"github.com/uber-go/tally/v4"
	"go.uber.org/zap"
)

func main() {
	logger := zap.NewExample().Sugar()
	rep, err := reporter.NewGCPStatsReporter(reporter.GCPStatsReporterIn{
		GCPConfiguration: &reporter.GCPConfiguration{
			ProjectID:  os.Getenv("GCP_PROJECT_ID"),
			MetricType: os.Getenv("GCP_METRIC_TYPE"),
		},
		Logger: logger,
	})
	if err != nil {
		logger.Fatal(err)
	}
	scope, closer := tally.NewRootScope(tally.ScopeOptions{
		Reporter: rep.GCPStatsReporter,
	}, 5*time.Second)
	defer closer.Close()
	scope.Gauge("foo").Update(1.0)
	scope.Counter("bar").Inc(1)
}
