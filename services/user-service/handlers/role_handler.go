package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/shehbaazsk/go-commerce-microservices/pkg/response"
	"github.com/shehbaazsk/go-commerce-microservices/user-service/models"
	"github.com/shehbaazsk/go-commerce-microservices/user-service/services"
)

type RoleHandler struct {
	RoleService services.RoleService
}

func NewRoleHandler(dbPool *pgxpool.Pool) *RoleHandler {
	return &RoleHandler{RoleService: services.NewRoleService(dbPool)}
}

func (h *RoleHandler) CreateRole(c *gin.Context) {
	var req models.CreateRoleRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequestError(c, "Invalid payload", err)
		return
	}
	role, err := h.RoleService.CreateRole(c, req)
	if err != nil {
		response.InternalServerError(c, "Could not create role", err)
		return
	}

	response.Success(c, http.StatusCreated, "Role created scuccessfully", role)
}

func (h *RoleHandler) GetRoleByID(c *gin.Context) {
	id := c.Param("id")
	roleID, err := strconv.Atoi(id)
	if err != nil {
		response.BadRequestError(c, "Invalid ID", err)
		return
	}

	role, err := h.RoleService.GetRoleByID(c, int32(roleID))
	if err != nil {
		response.NotFoundError(c, "Role not found")
		return
	}

	response.Success(c, http.StatusOK, "Role fetched", role)
}

func (h *RoleHandler) UpdateRole(c *gin.Context) {
	id := c.Param("id")
	roleID, err := strconv.Atoi(id)
	if err != nil {
		response.BadRequestError(c, "Invalid ID", err)
		return
	}

	var req models.UpdateRoleRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequestError(c, "Invalid payload", err)
		return
	}

	role, err := h.RoleService.UpdateRole(c, int32(roleID), req)
	if err != nil {
		response.InternalServerError(c, "Failed to update role", err)
		return
	}

	response.Success(c, http.StatusAccepted, "Role updated successfully", role)
}

func (h *RoleHandler) DeleteRole(c *gin.Context) {
	id := c.Param("id")
	roleID, err := strconv.Atoi(id)
	if err != nil {
		response.BadRequestError(c, "Invalid ID", err)
		return
	}

	if err := h.RoleService.DeleteRole(c, int32(roleID)); err != nil {
		response.InternalServerError(c, "Could not delete role", err)
		return
	}

	response.Success(c, http.StatusNoContent, "Role deleted", nil)
}

func (h *RoleHandler) ListRoles(c *gin.Context) {
	var req models.ListRoleRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, "Invalid Role Request", err)
	}
	req.SetDefaults()

	roles, err := h.RoleService.ListRoles(c, req)
	if err != nil {
		response.InternalServerError(c, "Failed to list roles", err)
		return
	}

	response.PaginatedSuccess(c, http.StatusOK, "success", response.PaginatedResponse{
		SearchBy: *req.SearchBy,
		OrderBy:  *req.OrderBy,
		OrderDir: *req.OrderDir,
		Page:     *req.Page,
		PerPage:  *req.PerPage,
		Results:  roles,
	})
}

// func (h *RoleHandler) AssignPermissionToRole(c *gin.Context) {
// 	roleID, err := strconv.Atoi(c.Param("id"))
// 	if err != nil {
// 		response.BadRequestError(c, "Invalid role ID", err)
// 		return
// 	}

// 	var req struct {
// 		PermissionID int32 `json:"permission_id" binding:"required"`
// 	}

// 	if err := c.ShouldBindJSON(&req); err != nil {
// 		response.BadRequestError(c, "Invalid payload", err)
// 		return
// 	}

// 	err = h.RoleService.AssignPermissionToRole(c, int32(roleID), req.PermissionID)
// 	if err != nil {
// 		response.InternalServerError(c, "Failed to assign permission", err)
// 		return
// 	}

// 	response.Success(c, "Permission assigned", nil)
// }

// func (h *RoleHandler) RemovePermissionFromRole(c *gin.Context) {
// 	roleID, err := strconv.Atoi(c.Param("id"))
// 	if err != nil {
// 		response.BadRequestError(c, "Invalid role ID", err)
// 		return
// 	}

// 	permID, err := strconv.Atoi(c.Param("permission_id"))
// 	if err != nil {
// 		response.BadRequestError(c, "Invalid permission ID", err)
// 		return
// 	}

// 	err = h.RoleService.RemovePermissionFromRole(c, int32(roleID), int32(permID))
// 	if err != nil {
// 		response.InternalServerError(c, "Failed to remove permission", err)
// 		return
// 	}

// 	response.Success(c, "Permission removed", nil)
// }

// func (h *RoleHandler) ListPermissionsByRole(c *gin.Context) {
// 	roleID, err := strconv.Atoi(c.Param("id"))
// 	if err != nil {
// 		response.BadRequestError(c, "Invalid role ID", err)
// 		return
// 	}

// 	perms, err := h.RoleService.ListPermissionsByRole(c, int32(roleID))
// 	if err != nil {
// 		response.InternalServerError(c, "Failed to fetch permissions", err)
// 		return
// 	}

// 	response.Success(c, "Permissions fetched", perms)
// }

// func parseInt(s string) int32 {
// 	v, _ := strconv.Atoi(s)
// 	return int32(v)
// }
