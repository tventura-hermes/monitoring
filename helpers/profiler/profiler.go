package profiler_helpers

import (
	"context"
	errors_helpers "demo/helpers/errors"
	"log/slog"

	"cloud.google.com/go/profiler"
)

func StartupProfiler(ctx context.Context) {
	cfg := profiler.Config{
		Service:        "test",
		ServiceVersion: "1.0.0",
		DebugLogging:   true,
	}

	if err := profiler.Start(cfg); err != nil {
		errors_helpers.ReportError(ctx, err)
		slog.ErrorContext(ctx, "error staring profiler")
	}
}
