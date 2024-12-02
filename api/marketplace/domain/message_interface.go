package marketplace_domain

import "github.com/gin-gonic/gin"

type SaveMessageInterface interface {
	SaveMessage(c *gin.Context)
}
