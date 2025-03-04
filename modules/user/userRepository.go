package user

import (
	"context"
	"github.com/reborn-hu/gsharp"
	"go.uber.org/zap"
)

type (
	IUserManageRepository interface {
		AutoMigrate() (err error)
		Add(entity *UserManageEntity) (res bool, row int64, err error)
		Update(primaryKey uint64, entity map[string]any) (res bool, row int64, err error)
		QueryList() []UserManageEntity
		QueryPageList(page int, size int) (total int64, list []UserManageEntity)
	}

	userManageRepository struct {
		context context.Context
		db      *gsharp.DbClient
		redis   gsharp.IRedis
	}
)

func UserManageRepository(ctx context.Context) IUserManageRepository {
	return &userManageRepository{
		context: ctx,
		db:      gsharp.GetDbClient(gsharp.BaseWrite),
		redis:   gsharp.GetRedisClient(ctx, "default"),
	}
}

func (repo *userManageRepository) AutoMigrate() (err error) {
	if err := repo.db.WithContext(repo.context).AutoMigrate(&UserManageEntity{}); err != nil {
		zap.L().Error(err.Error())
		return err
	}
	return nil
}

func (repo *userManageRepository) Add(entity *UserManageEntity) (res bool, row int64, err error) {
	result := repo.db.WithContext(repo.context).Where(UserManageEntity{Audience: entity.Audience}).FirstOrCreate(entity)
	if result.Error != nil {
		zap.L().Error(result.Error.Error())
		return false, result.RowsAffected, result.Error
	}
	return true, result.RowsAffected, nil
}

func (repo *userManageRepository) Update(primaryKey uint64, entity map[string]any) (res bool, row int64, err error) {
	result := repo.db.WithContext(repo.context).Table(UserManageEntity{}.TableName()).Where("id = ? ", primaryKey).Updates(entity)
	if result.Error != nil {
		zap.L().Error(result.Error.Error())
		return false, result.RowsAffected, result.Error
	}
	return true, result.RowsAffected, nil
}

func (repo *userManageRepository) QueryList() []UserManageEntity {
	var authorizeManage []UserManageEntity
	result := repo.db.WithContext(repo.context).Find(&authorizeManage)
	if result.Error != nil {
		return nil
	}
	return authorizeManage
}

func (repo *userManageRepository) QueryPageList(page int, size int) (total int64, list []UserManageEntity) {
	var authorizeManage []UserManageEntity
	result := repo.db.WithContext(repo.context).Model(&UserManageEntity{}).Count(&total).Scopes(gsharp.PaginateScope(page, size)).Find(&authorizeManage)
	if result.Error != nil {
		return 0, nil
	}
	return total, authorizeManage
}
