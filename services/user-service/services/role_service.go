package services

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/shehbaazsk/go-commerce-microservices/pkg/converters"
	db "github.com/shehbaazsk/go-commerce-microservices/user-service/db/queries"
	"github.com/shehbaazsk/go-commerce-microservices/user-service/models"
)

func mapRoleToResponse(role db.Role) models.RoleResponse {
	return models.RoleResponse{
		Name:        role.Name,
		Description: converters.TextOrNil(role.Description),
		IsActive:    converters.BoolOrNil(role.IsActive),
		CreatedAt:   role.CreatedAt.Time,
		UpdatedAt:   role.UpdatedAt.Time,
	}
}

type RoleService interface {
	CreateRole(ctx context.Context, req models.CreateRoleRequest) (models.RoleResponse, error)
	GetRoleByID(ctx context.Context, id int32) (models.RoleResponse, error)
	ListRoles(ctx context.Context, req models.ListRoleRequest) ([]models.RoleResponse, error)
	UpdateRole(ctx context.Context, id int32, req models.UpdateRoleRequest) (models.RoleResponse, error)
	DeleteRole(ctx context.Context, id int32) error

	// AssignPermission(ctx context.Context, roleID, permissionID int32) error
	// RemovePermission(ctx context.Context, roleID, permissionID int32) error
	// ListPermissionsByRole(ctx context.Context, roleID int32) ([]db.Permission, error)
}

type roleService struct {
	q *db.Queries
}

func NewRoleService(dbPool *pgxpool.Pool) RoleService {
	q := db.New(dbPool)
	return &roleService{q: q}
}

func (s *roleService) CreateRole(ctx context.Context, req models.CreateRoleRequest) (models.RoleResponse, error) {
	createRoleParams := db.CreateRoleParams{
		Name:        req.Name,
		Description: converters.StringToPgText(req.Description),
		Column3:     req.IsActive,
	}
	role, err := s.q.CreateRole(ctx, createRoleParams)
	if err != nil {
		return models.RoleResponse{}, err
	}

	return mapRoleToResponse(role), nil
}

func (s *roleService) GetRoleByID(ctx context.Context, id int32) (models.RoleResponse, error) {
	role, err := s.q.GetRoleByID(ctx, id)
	if err != nil {
		return models.RoleResponse{}, err
	}

	return mapRoleToResponse(role), nil
}

func (s *roleService) ListRoles(ctx context.Context, req models.ListRoleRequest) ([]models.RoleResponse, error) {
	listRoleParams := db.ListRolesPaginatedSearchParams{
		SearchBy: converters.StringToPgText(req.SearchBy),
		OrderBy:  *req.OrderBy,
		OrderDir: *req.OrderDir,
		Limit:    *req.PerPage,
		Offset:   req.Offset(),
	}
	roles, err := s.q.ListRolesPaginatedSearch(ctx, listRoleParams)
	if err != nil {
		return []models.RoleResponse{}, err
	}
	roleRes := []models.RoleResponse{}
	for _, r := range roles {
		roleRes = append(roleRes, models.RoleResponse{
			Name:        r.Name,
			Description: converters.TextOrNil(r.Description),
			IsActive:    converters.BoolOrNil(r.IsActive),
			CreatedAt:   r.CreatedAt.Time,
			UpdatedAt:   r.UpdatedAt.Time,
		})
	}
	return roleRes, nil

}

func (s *roleService) UpdateRole(ctx context.Context, id int32, req models.UpdateRoleRequest) (models.RoleResponse, error) {
	updateRoleParams := db.UpdateRoleParams{
		ID:          id,
		Name:        converters.StringToPgText(req.Name),
		Description: converters.StringToPgText(req.Description),
		IsActive:    converters.BoolToPgBool(req.IsActive),
	}
	role, err := s.q.UpdateRole(ctx, updateRoleParams)
	if err != nil {
		return models.RoleResponse{}, err
	}
	return mapRoleToResponse(role), nil
}

func (s *roleService) DeleteRole(ctx context.Context, id int32) error {
	return s.q.DeleteRole(ctx, id)
}

// func (s *roleService) AssignPermission(ctx context.Context, roleID, permissionID int32) error {
// 	return q.AssignPermissionToRole(ctx, db.AssignPermissionToRoleParams{
// 		RoleID:       roleID,
// 		PermissionID: permissionID,
// 	})
// }

// func (s *roleService) RemovePermission(ctx context.Context, roleID, permissionID int32) error {
// 	return q.RemovePermissionFromRole(ctx, db.RemovePermissionFromRoleParams{
// 		RoleID:       roleID,
// 		PermissionID: permissionID,
// 	})
// }

// func (s *roleService) ListPermissionsByRole(ctx context.Context, roleID int32) ([]db.Permission, error) {
// 	return q.ListPermissionsByRole(ctx, roleID)
// }
