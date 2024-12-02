package marketplace_infrastructure

import (
	marketplace_domain "demo/api/marketplace/domain"
	"demo/db"
	cache_helpers "demo/helpers/cache"
	errors_helpers "demo/helpers/errors"
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	oteltrace "go.opentelemetry.io/otel/trace"
)

var tracer = otel.Tracer("get-messages")

type MarketplaceHandler struct {
}

func NewMarketplaceHandler(server *gin.Context) marketplace_domain.SaveMessageInterface {
	return &MarketplaceHandler{}
}

func (m *MarketplaceHandler) SaveMessage(c *gin.Context) {
	ctx := c.Request.Context()

	cacheKey := "getMarketplace"
	cache, setupError := cache_helpers.NewCache(ctx)

	if setupError != nil {
		errors_helpers.ReportError(ctx, setupError)
		c.JSON(http.StatusExpectationFailed, gin.H{"error": "Internal using cache server"})
	}

	cacheData, err := cache.Get(cacheKey)

	if err != nil {
		database, err := db.NewDatabase(ctx)
		if err != nil {
			errors_helpers.ReportError(ctx, err)
			c.JSON(http.StatusExpectationFailed, gin.H{"error": "Internal Setup error"})
			return
		}

		result, err := database.Select("marketplace")

		if err != nil {
			c.JSON(http.StatusExpectationFailed, gin.H{"error": "Internal error"})
			errors_helpers.ReportError(ctx, err)
			return
		}

		slog.InfoContext(c.Request.Context(), "GET handle /marketplace query", slog.String("db.collection", "marketplace"))

		_, span := tracer.Start(c.Request.Context(), "read", oteltrace.WithAttributes(
			attribute.Bool("trace", true),
			attribute.String("db.collection", "marketplace"),
			attribute.Bool("cache", false),
		))
		defer span.End()
		c.JSON(http.StatusOK, gin.H{"data": result})
	}

	cache.Close()
	slog.InfoContext(c.Request.Context(), "GET handle /marketplace query", slog.String("db.collection", "marketplace"))

	_, span := tracer.Start(c.Request.Context(), "read", oteltrace.WithAttributes(
		attribute.Bool("trace", true),
		attribute.String("db.collection", "marketplace"),
		attribute.Bool("cache", true),
	))
	defer span.End()
	c.JSON(http.StatusOK, gin.H{"data": cacheData})
}
