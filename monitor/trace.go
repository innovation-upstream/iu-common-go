package monitor

import (
	"context"
	"os"

	"github.com/pkg/errors"

	texporter "github.com/GoogleCloudPlatform/opentelemetry-operations-go/exporter/trace"
	"go.opentelemetry.io/otel"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
)

func InitCloudTraceOpenTelemetry(ctx context.Context) (func(ctx context.Context) error, error) {
	// TODO: add support for alternate, open-source exporters
	projectID := os.Getenv("GOOGLE_CLOUD_PROJECT")
	exporter, err := texporter.NewExporter(texporter.WithProjectID(projectID))
	if err != nil {
		return func(ctx context.Context) error {
			return nil
		}, errors.WithStack(err)
	}

	tp := sdktrace.NewTracerProvider(
		sdktrace.WithBatcher(exporter),
		sdktrace.WithSampler(sdktrace.TraceIDRatioBased(0.0001)),
	)
	otel.SetTracerProvider(tp)

	return tp.ForceFlush, nil
}
