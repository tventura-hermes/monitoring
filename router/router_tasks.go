package router

import (
	tasks_infrastructure "demo/api/tasks/infrastructure"

	"github.com/gin-gonic/gin"
)

func (ro *Routes) TasksRoutes(r *gin.Engine) {
	tr := tasks_infrastructure.NewTaskHandler(ro.Context)
	r.POST("/tasks", tr.CreateTask)
}
