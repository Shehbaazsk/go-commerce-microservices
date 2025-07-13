package routes

import (
	"github.com/gin-gonic/gin"
	middlewares "github.com/shehbaazsk/go-commerce-micorservices/api-gateway/middleware"
)

func SetupRoutes(r *gin.Engine) {

	// Public routes
	public := r.Group("/api/v1")
	RegisterPublicRoutes(public)

	// Protected routes
	protected := r.Group("/api/v1")
	protected.Use(middlewares.AuthMiddleware())
	RegisterProtectedRoutes(protected)

}
