package repository

import (
	"context"
	"go-nunu/internal/model"
)

type RoleRepository interface {
	GetRole(ctx context.Context, id int64) (*model.Role, error)
	GetRoleList(ctx context.Context) ([]model.Role, error)
	CreateRole(ctx context.Context, role *model.Role) (*model.Role, error)
	UpdateRole(ctx context.Context, role *model.Role) (*model.Role, error)
}

func NewRoleRepository(
	repository *Repository,
) RoleRepository {
	return &roleRepository{
		Repository: repository,
	}
}

type roleRepository struct {
	*Repository
}

func (r *roleRepository) GetRole(ctx context.Context, id int64) (*model.Role, error) {
	var role model.Role

	return &role, nil
}

func (r *roleRepository) GetRoleList(ctx context.Context) ([]model.Role, error) {
	var roles []model.Role
	err := r.db.WithContext(ctx).Preload("Permissions").Find(&roles).Error
	if err != nil {
		return nil, err
	}
	return roles, nil
}

// CreateRole 创建角色
// 参数：需要把要创建的角色数据 (role) 传进来
// 返回：error (如果创建失败)
func (r *roleRepository) CreateRole(ctx context.Context, role *model.Role) (*model.Role, error) {
	// err := r.db.WithContext(ctx).Create(role).Error
	err := r.DB(ctx).
		Model(&role).
		Association("Permissions").
		Replace(role.Permissions) // 传入的是完整的 Permission 实体数组 (可能为空)
	return role, err
}

func (r *roleRepository) UpdateRole(ctx context.Context, role *model.Role) (*model.Role, error) {
	err := r.db.WithContext(ctx).Save(role).Error
	return role, err
}
