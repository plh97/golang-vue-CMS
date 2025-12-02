package handler

import (
	"fmt"
	v1 "go-nunu/api/v1"
	"go-nunu/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PermissionHandler struct {
	*Handler
	permissionService service.PermissionService
}

func NewPermissionHandler(
	handler *Handler,
	permissionService service.PermissionService,
) *PermissionHandler {
	return &PermissionHandler{
		Handler:           handler,
		permissionService: permissionService,
	}
}

func (h *PermissionHandler) GetPermissionList(ctx *gin.Context) {
	// Implementation for getting a permission
	permissionList, err := h.permissionService.GetPermissionList(ctx)
	if err != nil {
		v1.HandleError(ctx, http.StatusOK, v1.ErrBadRequest, nil)
		return
	}

	v1.HandleSuccess(ctx, v1.GetPermissionListResponseData{
		List: permissionList,
	})
}

func (h *PermissionHandler) CreatePermission(ctx *gin.Context) {
	var req v1.CreatePermissionRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		v1.HandleError(ctx, http.StatusBadRequest, v1.ErrBadRequest, nil)
		return
	}
	permission, err := h.permissionService.CreatePermission(ctx, req)
	if err != nil {
		fmt.Println("error:", err)
		v1.HandleError(ctx, http.StatusOK, v1.ErrBadRequest, nil)
		return
	}

	v1.HandleSuccess(ctx, permission)
}
