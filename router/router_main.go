package router

import (
	"context"
	errors_helpers "demo/helpers/errors"
	trace_helpers "demo/helpers/trace/api"
	"log/slog"
	"os"

	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
)

type Routes struct {
	Routes  *gin.Engine
	Context *gin.Context
}

func (r *Routes) SetupRouter() *gin.Engine {
	router := gin.Default()
	return router
}

func (r *Routes) CreateRoutes() {
	router := r.SetupRouter()
	router.Use(otelgin.Middleware("demo"))

	r.MarketplaceRoutes(router)
	r.TasksRoutes(router)

	r.Routes = router
}

func (r *Routes) StartServer() {
	ctx := context.Background()
	trace_helpers.SetupLogging()

	shutdown, err := trace_helpers.SetupOpenTelemetry(ctx)
	if err != nil {
		errors_helpers.ReportError(ctx, err)
		slog.ErrorContext(ctx, "error setting up OpenTelemetry", slog.Any("error", err))
		os.Exit(1)
	}

	defer func() {
		if err := shutdown(ctx); err != nil {
			errors_helpers.ReportError(ctx, err)
			slog.ErrorContext(ctx, "failed to shutdown OpenTelemetry", slog.Any("error", err))
		}
	}()

	slog.InfoContext(ctx, "start opentelemetry")

	errRun := r.Routes.Run(":8080")
	if errRun != nil {
		errors_helpers.ReportError(ctx, err)
		slog.ErrorContext(ctx, "Gin server exited with error", slog.Any("error", err))
	}

}
