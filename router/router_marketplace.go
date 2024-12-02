package router

import (
	marketplace_infrastructure "demo/api/marketplace/infrastructure"

	"github.com/gin-gonic/gin"
)

func (ro *Routes) MarketplaceRoutes(r *gin.Engine) {
	mr := marketplace_infrastructure.NewMarketplaceHandler(ro.Context)
	r.GET("/marketplace", mr.SaveMessage)
}
