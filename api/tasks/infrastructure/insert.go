package tasks_infrastructure

import (
	tasks_domain "demo/api/tasks/domain"
	"demo/db"
	errors_helpers "demo/helpers/errors"
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
)

var tracer = otel.Tracer("create-tasks")

type TaskHandler struct {
}

func NewTaskHandler(server *gin.Context) tasks_domain.CreateTaskInterface {
	return &TaskHandler{}
}

func (m *TaskHandler) CreateTask(c *gin.Context) {
	ctx := c.Request.Context()
	var post tasks_domain.Task

	ctx, span := tracer.Start(ctx, "CreateTask")
	defer span.End()

	if err := c.BindJSON(&post); err != nil {
		errors_helpers.ReportError(ctx, err)
		span.RecordError(err)
		span.SetStatus(codes.Error, "Failed to bind JSON")
		c.JSON(http.StatusExpectationFailed, gin.H{"error": "Internal error request"})
		return
	}

	data := post.ToMongo()

	database, err := db.NewDatabase(ctx)
	if err != nil {
		errors_helpers.ReportError(ctx, err)
		span.RecordError(err)
		span.SetStatus(codes.Error, "Database setup failed")
		c.JSON(http.StatusExpectationFailed, gin.H{"error": "Internal Setup error"})
		return
	}

	dberr := database.Insert(data, "tasks")
	if dberr != nil {
		errors_helpers.ReportError(ctx, err)
		span.RecordError(dberr)
		span.SetStatus(codes.Error, "Database insert failed")
		c.JSON(http.StatusExpectationFailed, gin.H{"error": "Internal Database  Insert error"})
		errors_helpers.ReportError(ctx, err)
		return
	}

	span.SetAttributes(
		attribute.String("db.collection", "tasks"),
		attribute.Bool("cache", false),
	)

	slog.InfoContext(ctx, "POST handle /tasks query", slog.String("db.collection", "tasks"))
	c.JSON(http.StatusOK, gin.H{"message": "Data created"})
}
