package router

import (
	"context"
	trace_helpers "demo/helpers/trace/api"
	"errors"
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

	r.Routes = router
}

func (r *Routes) StartServer() {
	ctx := context.Background()
	trace_helpers.SetupLogging()

	shutdown, err := trace_helpers.SetupOpenTelemetry(ctx)
	if err != nil {
		slog.ErrorContext(ctx, "error setting up OpenTelemetry", slog.Any("error", err))
		os.Exit(1)
	}

	go r.Routes.Run(":8080")

	if err = errors.Join(shutdown(ctx)); err != nil {
		slog.ErrorContext(ctx, "server exited with error", slog.Any("error", err))
		os.Exit(1)
	}
}
