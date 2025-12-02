package v1

import "go-nunu/internal/model"

type GetPermissionListResponseData struct {
	// UserId   string `json:"userId"`
	// Nickname string `json:"nickname" example:"alan"`
	// Email    string `json:"email" example:"alan"`
	List []model.Permission `json:"list"`
}
type GetPermissionListResponse struct {
	Response
	Data GetPermissionListResponseData `json:"data"`
}

type CreatePermissionRequest struct {
	Name   string `json:"name"`
	Key    string `json:"key"`
	Status int    `json:"status"`
	// Permissions []Permission `gorm:"many2many:sys_role_permissions;" json:"permissions"`
}
type CreatePermissionResponse struct {
	Response
	Data model.Permission `json:"data"`
}
