package user

import (
	"github.com/reborn-hu/gsharp"
	"strconv"
	"time"
)

// UserManageEntity 数据库实体
type UserManageEntity struct {
	gsharp.PrimaryOptions `gorm:"embedded"`
	Name                  string `json:"name" gorm:"column:name;comment:授信应用名称"`
	Audience              string `json:"audience" gorm:"column:audience;comment:授信应用Code或域名"`
	Key                   string `json:"key" gorm:"column:key;comment:应用Key"`
	Secret                string `json:"secret" gorm:"column:secret;comment:应用秘钥"`
	PrivateKey            string `json:"privateKey"  gorm:"column:private_key;comment:私钥"`
	AccessExpires         int    `json:"accessExpires" gorm:"column:access_expires;comment:accessToken有效时间"`
	RefreshExpires        int    `json:"refreshExpires" gorm:"column:refresh_expires;comment:refreshToken有效时间"`
	gsharp.OperateOptions `gorm:"embedded"`
}

// UserManageReq 请求对象
type UserManageReq struct {
	Name           string `json:"name" binding:"required"`     // 授信应用名称
	Audience       string `json:"audience" binding:"required"` // 授信应用Code或域名
	Key            string `json:"key" binding:"required"`      // 应用Key
	AccessExpires  int    `json:"accessExpires"`               // accessToken有效时间
	RefreshExpires int    `json:"refreshExpires"`              // refreshToken有效时间
}

// UserManage 列表对象
type UserManage struct {
	Id       string `json:"id"`       // 唯一业务ID
	Name     string `json:"name"`     // 授信应用名称
	Audience string `json:"audience"` // 授信应用Code或域名
	gsharp.OperateOptions
}

// UserManageDetails 详情
type UserManageDetails struct {
	UserManage
	Key            string `json:"key"`
	Secret         string `json:"secret"`
	PrivateKey     string `json:"privateKey"`
	AccessExpires  int    `json:"accessExpires"`
	RefreshExpires int    `json:"refreshExpires"`
}

// TableName 获取表名
func (UserManageEntity) TableName() string {
	return "authorize_manage"
}

func (model *UserManage) ToId() uint64 {
	val, err := strconv.ParseUint(model.Id, 10, 64)
	if err != nil {
		panic("非法类型字符串")
	}
	return val
}

func (entity UserManageEntity) DetailsEntity() *UserManageDetails {
	dto := new(UserManageDetails)
	err := gsharp.Mapper(&entity, &dto)
	if err != nil {
		return nil
	}
	return dto
}

func (req *UserManageReq) CreateEntity(createUser string) *UserManageEntity {
	entity := new(UserManageEntity)
	entity.Del = false
	entity.CreateUser = createUser
	entity.CreateDate = gsharp.DataTime(time.Now())
	entity.ModifyDate = gsharp.DataTime(time.Now())
	entity.AccessExpires = 7200
	entity.RefreshExpires = 8000

	err := gsharp.Mapper(&req, &entity)
	if err != nil {
		return nil
	}

	rsaKey := gsharp.CreateCryptology().GenRsaKey()
	entity.Secret = rsaKey.PublicKey
	entity.PrivateKey = rsaKey.PrivateKey

	return entity
}

func (req *UserManageDetails) UpdateEntity(modifyUser string) map[string]any {
	entityMap := make(map[string]any)
	entityMap["key"] = req.Key
	entityMap["access_expires"] = req.AccessExpires
	entityMap["refresh_expires"] = req.RefreshExpires
	entityMap["modify_user"] = modifyUser
	entityMap["modify_date"] = gsharp.DataTime(time.Now())

	return entityMap
}
