package monitor

import (
	"os"

	texporter "github.com/GoogleCloudPlatform/opentelemetry-operations-go/exporter/trace"
	"go.opentelemetry.io/otel/api/global"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	"unknwon.dev/clog/v2"
)

func InitCloudTraceOpenTelemetry() {
	projectID := os.Getenv("GOOGLE_CLOUD_PROJECT")
	exporter, err := texporter.NewExporter(texporter.WithProjectID(projectID))
	if err != nil {
		clog.Fatal("texporter.NewExporter: %+v", err)
	}

	tp, err := sdktrace.NewProvider(sdktrace.WithSyncer(exporter))
	if err != nil {
		clog.Fatal("%+v", err)
	}
	global.SetTraceProvider(tp)
}
