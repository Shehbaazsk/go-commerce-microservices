package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

func SetupRoutes(r *gin.Engine, dbPool *pgxpool.Pool) {

	// Public routes
	userPublic := r.Group("/users")
	RegisterUserPublicRoutes(userPublic)

	// Protected routes
	userProtected := r.Group("/users")
	RegisterUserProtectedRoutes(userProtected)

	roleProtected := r.Group("/api/v1/roles")
	RegisterRoleProtectedRoutes(roleProtected, dbPool)

}
