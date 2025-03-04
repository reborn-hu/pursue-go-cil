package user

import (
	"github.com/gin-gonic/gin"
	"github.com/reborn-hu/gsharp"
	"strconv"
)

// AutoMigrateHandler
// @summary 创建数据表
// @Description AutoMigrateHandler
// @Tags AutoMigrateHandler
// @Accept json
// @Produce json
// @Success 200 {object} object "回调对象"
// @Router /api/user/autoMigrate [POST]
func AutoMigrateHandler(ctx *gin.Context) {
	var svc = UserManageService(ctx)
	if err := svc.AutoMigrate(); err != nil {
		gsharp.CallbackError(ctx, gsharp.CodeSeverError)
	}
	gsharp.CallbackSuccess(ctx, true)
}

// CreateUserManageHandler
// @summary 授权应用管理
// @Description CreateUserManageHandler
// @Tags CreateUserManageHandler
// @Accept json
// @Produce json
// @Success 200 {object} object "回调对象"
// @Param Authorize body UserManageReq true "授权请求对象"
// @Router /api/user/create [POST]
func CreateUserManageHandler(ctx *gin.Context) {
	//user := core.GetTenant(ctx)
	//fmt.Print(user)
	var form = new(UserManageReq)
	if err := ctx.ShouldBind(&form); err != nil {
		gsharp.CallbackError(ctx, gsharp.CodeOperationFail)
		return
	}

	svc := UserManageService(ctx)
	result := svc.Create(form)

	gsharp.CallbackSuccess(ctx, result)
}

// UpdateUserManageHandler
// @summary 授权应用管理
// @Description UpdateUserManageHandler
// @Tags UpdateUserManageHandler
// @Accept json
// @Produce json
// @Success 200 {object} object "回调对象"
// @Param Authorize body UserManageDetails true "授权请求对象"
// @Router /api/user/update [POST]
func UpdateUserManageHandler(ctx *gin.Context) {
	var form = new(UserManageDetails)
	if err := ctx.ShouldBind(&form); err != nil {
		gsharp.CallbackError(ctx, gsharp.CodeOperationFail)
		return
	}

	svc := UserManageService(ctx)
	result := svc.Update(form)

	gsharp.CallbackSuccess(ctx, result)
}

// QueryUserManageListHandler
// @summary 查询授权应用列表
// @Description QueryUserManageListHandler
// @Tags QueryUserManageListHandler
// @Accept json
// @Produce json
// @Success 200 {object} object "回调对象"
// @Router /api/user/query [GET]
func QueryUserManageListHandler(ctx *gin.Context) {
	svc := UserManageService(ctx)
	result := svc.QueryList()
	gsharp.CallbackSuccess(ctx, result)
}

// QueryUserManagePageHandler
// @summary 查询授权应用分页列表
// @Description QueryUserManagePageHandler
// @Tags QueryUserManagePageHandler
// @Accept json
// @Produce json
// @Success 200 {object} object "回调对象"
// @Router /api/user/query/:page/:size [GET]
func QueryUserManagePageHandler(ctx *gin.Context) {
	page, err := strconv.Atoi(ctx.Param("page"))
	size, err := strconv.Atoi(ctx.Param("size"))
	if err != nil {
		gsharp.CallbackError(ctx, gsharp.CodeOperationFail)
		return
	}
	svc := UserManageService(ctx)
	result := svc.QueryPageList(page, size)
	gsharp.CallbackSuccess(ctx, result)
}
