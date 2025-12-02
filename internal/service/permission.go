package service

import (
	"context"
	v1 "go-nunu/api/v1"
	"go-nunu/internal/model"
	"go-nunu/internal/repository"
)

type PermissionService interface {
	GetPermissionList(ctx context.Context) ([]model.Permission, error)
	CreatePermission(ctx context.Context, req v1.CreatePermissionRequest) (*model.Permission, error)
}

func NewPermissionService(
	service *Service,
	permissionRepository repository.PermissionRepository,
) PermissionService {
	return &permissionService{
		Service:              service,
		permissionRepository: permissionRepository,
	}
}

type permissionService struct {
	*Service
	permissionRepository repository.PermissionRepository
}

func (s *permissionService) GetPermissionList(ctx context.Context) ([]model.Permission, error) {
	return s.permissionRepository.GetPermissionList(ctx)
}

func (s *permissionService) CreatePermission(ctx context.Context, req v1.CreatePermissionRequest) (*model.Permission, error) {
	permission := &model.Permission{
		Name:   req.Name,
		Key:    req.Key,
		// Status: req.Status,
	}
	return s.permissionRepository.CreatePermission(ctx, permission)
}
