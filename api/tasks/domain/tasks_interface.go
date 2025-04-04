package tasks_domain

import "github.com/gin-gonic/gin"

type TaskInterface interface {
}

// interfaz para crear un 'task'
type CreateTaskInterface interface {
	CreateTask(c *gin.Context)
}

// interfaz para obtener todos los 'tasks'
type ReadAllTaskInterface interface {
	ReadData(c *gin.Context)
}

// interfaz para actualizar un 'task'
type UpdateTaskInterface interface {
	UpdateData(c *gin.Context)
}

// interfaz para eliminar un 'task'
type DeleteTaskInterface interface {
	DeleteData(c *gin.Context)
}
