package pubsub_trigger

import (
	"context"
	errors_helpers "demo/helpers/errors"
	otel_helpers "demo/helpers/trace/pubsub"
	"os"

	"cloud.google.com/go/pubsub"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/sdk/resource"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
)

type PubsubClient struct {
	client pubsub.Client
	ctx    context.Context
}

func SetupPubsub(ctx context.Context) *PubsubClient {
	projecId := os.Getenv("GOOGLE_CLOUD_PROJECT")

	resources := resource.NewWithAttributes(
		semconv.SchemaURL,
		semconv.ServiceNameKey.String("subscriber"),
		attribute.String("collection", "marketplace"),
		attribute.String("context", "mongodb"),
	)

	otel_helpers.SetupTracer(ctx, resources)

	client, err := pubsub.NewClientWithConfig(ctx, projecId, &pubsub.ClientConfig{
		EnableOpenTelemetryTracing: true,
	})

	if err != nil {
		errors_helpers.ReportError(ctx, err)
		return nil
	}

	return &PubsubClient{
		client: *client,
		ctx:    ctx,
	}
}

func (p *PubsubClient) CloseClient() {
	defer p.client.Close()
}
