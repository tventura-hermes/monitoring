package otel_helpers

import (
	"context"
	errors_helpers "demo/helpers/errors"
	"os"

	texporter "github.com/GoogleCloudPlatform/opentelemetry-operations-go/exporter/trace"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	"google.golang.org/api/option"
)

func SetupTracer(ctx context.Context, resources *resource.Resource) {
	projecId := os.Getenv("GOOGLE_CLOUD_PROJECT")
	exporter, err := texporter.New(texporter.WithProjectID(projecId),
		texporter.WithTraceClientOptions(
			[]option.ClientOption{option.WithTelemetryDisabled()},
		),
	)

	if err != nil {
		errors_helpers.ReportError(ctx, err)
		return
	}

	tp := sdktrace.NewTracerProvider(
		sdktrace.WithBatcher(exporter),
		sdktrace.WithResource(resources),
		sdktrace.WithSampler(
			sdktrace.ParentBased(sdktrace.TraceIDRatioBased(0.5)),
		),
	)

	defer tp.ForceFlush(ctx)
	otel.SetTracerProvider(tp)
}
