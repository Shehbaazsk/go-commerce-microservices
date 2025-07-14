package routes

import (
	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {

	// Public routes
	public := r.Group("/users")
	RegisterPublicRoutes(public)

	// Protected routes
	protected := r.Group("/users")
	RegisterProtectedRoutes(protected)

}
