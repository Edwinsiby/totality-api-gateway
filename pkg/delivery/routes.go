package delivery

import "github.com/gin-gonic/gin"

type UserRoutes struct {
	Handlers Handlers
}

func NewUserRoutes(handlers Handlers) UserRoutes {
	return UserRoutes{
		Handlers: handlers,
	}
}

func (h UserRoutes) StartUserRoutes(router *gin.Engine) {
	router.GET("health-check", h.Handlers.HealthCheck)
	router.POST("user/details", h.Handlers.UserDetailsByID)
	router.POST("user/list/details", h.Handlers.UserDetailsByList)
}
