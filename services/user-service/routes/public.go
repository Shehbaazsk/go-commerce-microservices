package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/shehbaazsk/go-commerce-microservices/user-service/handlers"
)

func RegisterUserProtectedRoutes(r *gin.RouterGroup) {

}

func RegisterRoleProtectedRoutes(r *gin.RouterGroup, dbPool *pgxpool.Pool) {
	h := handlers.NewRoleHandler(dbPool)
	r.POST("/", h.ListRoles)
	r.PATCH("/:id", h.UpdateRole)
	r.DELETE("/:id", h.DeleteRole)
	r.GET("/:id", h.GetRoleByID)
	r.POST("/create", h.CreateRole)

}
