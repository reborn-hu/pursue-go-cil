package user

import (
	"github.com/gin-gonic/gin"
	"github.com/reborn-hu/gsharp"
	"go.uber.org/zap"
)

type (
	IUserManageService interface {
		AutoMigrate() error
		Create(request *UserManageReq) bool
		Update(request *UserManageDetails) bool
		QueryList() []UserManageEntity
		QueryPageList(page int, size int) *gsharp.PageList
	}

	userManageService struct {
		context              *gin.Context
		userManageRepository IUserManageRepository
	}
)

func UserManageService(ctx *gin.Context) IUserManageService {
	return &userManageService{
		context:              ctx,
		userManageRepository: UserManageRepository(ctx),
	}
}

func (svc *userManageService) AutoMigrate() error {
	if err := svc.userManageRepository.AutoMigrate(); err != nil {
		return err
	}
	return nil
}

func (svc *userManageService) Create(req *UserManageReq) bool {
	entity := req.CreateEntity("sys")
	res, _, err := svc.userManageRepository.Add(entity)
	if err != nil {
		zap.L().Error(err.Error())
		return false
	} else {
		return res
	}
}

func (svc *userManageService) Update(req *UserManageDetails) bool {
	entity := req.UpdateEntity("sys")
	res, _, err := svc.userManageRepository.Update(req.ToId(), entity)
	if err != nil {
		zap.L().Error(err.Error())
		return false
	} else {
		return res
	}
}

func (svc *userManageService) QueryList() []UserManageEntity {
	return svc.userManageRepository.QueryList()
}

func (svc *userManageService) QueryPageList(page int, size int) *gsharp.PageList {
	total, items := svc.userManageRepository.QueryPageList(page, size)
	return gsharp.CreatePageList(total, items)
}
