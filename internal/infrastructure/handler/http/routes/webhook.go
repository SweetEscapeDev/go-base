package routes

import (
	"github.com/gin-gonic/gin"
)

type WebhookRouteInterface interface {
	RegisterRoutes()
	VtWeb(c *gin.Context)
	Chat(c *gin.Context)
}

type WebhookbRoute struct {
	router *gin.Engine
}

func WebhookbRouteInit(router *gin.Engine) WebhookRouteInterface {
	route := &WebhookbRoute{
		router: router,
	}
	route.RegisterRoutes()
	return route
}

func (r *WebhookbRoute) RegisterRoutes() {
	r.router.POST("/webhook/vtweb", r.VtWeb)
	r.router.POST("/chat-webhook", r.Chat)
}

func (r *WebhookbRoute) VtWeb(c *gin.Context) {

}

func (r *WebhookbRoute) Chat(c *gin.Context) {

}
