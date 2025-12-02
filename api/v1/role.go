package v1

import "go-nunu/internal/model"

type GetRoleListResponseData struct {
	// UserId   string `json:"userId"`
	// Nickname string `json:"nickname" example:"alan"`
	// Email    string `json:"email" example:"alan"`
	List []model.Role `json:"list"`
}
type GetRoleListResponse struct {
	Response
	Data GetRoleListResponseData `json:"data"`
}

type CreateRoleRequest struct {
	Name          string `json:"name"`
	Key           string `json:"key"`
	Status        int    `json:"status"`
	PermissionIds []int  `json:"permission_ids"`
}

type UpdateRolePermissionsRequest struct {
	ID            int64 `json:"id"`
	PermissionIds []int `json:"permission_ids"`
}
type CreateRoleResponse struct {
	Response
	Data model.Role `json:"data"`
}
