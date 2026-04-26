package controller

import (
	"net/http"
	"time"

	"ginBackend/internal/model"
	"ginBackend/internal/service"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService service.UserService
}

func NewUserController(userService service.UserService) *UserController {
	return &UserController{userService: userService}
}

// Register a new user
func (ctrl *UserController) Register(ctx *gin.Context) {
	var req model.RegisterRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		respondError(ctx, http.StatusBadRequest, "發生錯誤", "提示：電子郵件, 密碼為必填")
		return
	}

	res, err := ctrl.userService.Register(&req)
	if err != nil {
		respondError(ctx, http.StatusConflict, err.Error(), nil)
		return
	}

	respondSuccess(ctx, http.StatusOK, "註冊成功", res)
}

// Login. Return a JWT token on success
func (ctrl *UserController) Login(ctx *gin.Context) {
	var req model.LoginRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		respondError(ctx, http.StatusBadRequest, "發生錯誤", "提示：電子郵件, 密碼為必填")
		return
	}

	res, err := ctrl.userService.Login(&req)
	if err != nil {
		respondError(ctx, http.StatusUnauthorized, err.Error(), nil)
		return
	}

	respondSuccess(ctx, http.StatusOK, "登入成功", res)
}

// Change the user's password. JWT token Required.
func (ctrl *UserController) ChangePassword(ctx *gin.Context) {
	var req model.ChangePasswordRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		respondError(ctx, http.StatusBadRequest, "發生錯誤", "提示：電子郵件, 舊密碼, 新密碼為必填")
		return
	}

	if err := ctrl.userService.ChangePassword(&req); err != nil {
		respondError(ctx, http.StatusBadRequest, err.Error(), nil)
		return
	}

	respondSuccess(ctx, http.StatusOK, "密碼已更新", nil)
}

func respondSuccess(ctx *gin.Context, code int, message string, data interface{}) {
	ctx.JSON(code, model.APIResponse{
		Success:   true,
		Message:   message,
		Data:      data,
		Error:     nil,
		Code:      code,
		Timestamp: time.Now().UTC(),
	})
}

func respondError(ctx *gin.Context, code int, message string, detail interface{}) {
	var errBody interface{}
	if detail != nil {
		errBody = gin.H{"detail": detail}
	}
	ctx.JSON(code, model.APIResponse{
		Success:   false,
		Message:   message,
		Data:      nil,
		Error:     errBody,
		Code:      code,
		Timestamp: time.Now().UTC(),
	})
}
