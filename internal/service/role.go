package service

import (
	"context"
	v1 "go-nunu/api/v1"
	"go-nunu/internal/model"
	"go-nunu/internal/repository"
)

type RoleService interface {
	GetRole(ctx context.Context, id int64) (*model.Role, error)
	CreateRole(ctx context.Context, req v1.CreateRoleRequest) (*model.Role, error)
	GetRoleList(ctx context.Context) ([]model.Role, error)
	UpdateRolePermissions(ctx context.Context, roleId int64, permissionIds []int) error
}

func NewRoleService(
	service *Service,
	roleRepository repository.RoleRepository,
) RoleService {
	return &roleService{
		Service:        service,
		roleRepository: roleRepository,
	}
}

type roleService struct {
	*Service
	roleRepository repository.RoleRepository
}

func (s *roleService) GetRole(ctx context.Context, id int64) (*model.Role, error) {
	return s.roleRepository.GetRole(ctx, id)
}

func (s *roleService) CreateRole(ctx context.Context, req v1.CreateRoleRequest) (*model.Role, error) {
	permissions := make([]model.Permission, len(req.PermissionIds))
	for i, id := range req.PermissionIds {
		permissions[i] = model.Permission{ID: id}
	}
	role := &model.Role{
		Name:        req.Name,
		Key:         req.Key,
		Status:      req.Status,
		Permissions: permissions,
	}
	return s.roleRepository.CreateRole(ctx, role)
}

func (s *roleService) GetRoleList(ctx context.Context) ([]model.Role, error) {
	return s.roleRepository.GetRoleList(ctx)
}

func (s *roleService) UpdateRolePermissions(ctx context.Context, roleId int64, permissionIds []int) error {
	// Fetch the role
	role, err := s.roleRepository.GetRole(ctx, roleId)
	if err != nil {
		return err
	}

	// Update permissions
	permissions := make([]model.Permission, len(permissionIds))
	for i, id := range permissionIds {
		permissions[i] = model.Permission{ID: id}
	}
	role.Permissions = permissions

	// Here you would typically call a repository method to save the updated role.
	// For simplicity, we'll assume the roleRepository has an UpdateRole method.
	_, err = s.roleRepository.UpdateRole(ctx, role) // This should be an update method in a real scenario
	return err
}
