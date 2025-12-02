package repository

import (
    "context"
	"go-nunu/internal/model"
)

type PermissionRepository interface {
	GetPermissionList(ctx context.Context) ([]model.Permission, error)
	CreatePermission(ctx context.Context, permission *model.Permission) (*model.Permission, error)
}

func NewPermissionRepository(
	repository *Repository,
) PermissionRepository {
	return &permissionRepository{
		Repository: repository,
	}
}

type permissionRepository struct {
	*Repository
}

func (r *permissionRepository) GetPermissionList(ctx context.Context) ([]model.Permission, error) {
	var permissionList []model.Permission
	err := r.db.WithContext(ctx).Find(&permissionList).Error
	if err != nil {
		return nil, err
	}
	return permissionList, nil
}

func (r *permissionRepository) CreatePermission(ctx context.Context, permission *model.Permission) (*model.Permission, error) {
	err := r.db.WithContext(ctx).Create(permission).Error
	return permission, err
}