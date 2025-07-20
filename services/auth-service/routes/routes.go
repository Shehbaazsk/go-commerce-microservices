package routes

import (
	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {

	// Public routes
	public := r.Group("/auth")
	RegisterPublicRoutes(public)

	// Protected routes
	protected := r.Group("/auth")
	RegisterProtectedRoutes(protected)

}
