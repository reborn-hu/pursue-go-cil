package modules

import (
	"github.com/reborn-hu/gsharp"
	"pursue-go-cil/modules/user"
)

func GetRouter() (router map[string]gsharp.RouterOptions) {

	routerMap := make(map[string]gsharp.RouterOptions)

	routerMap["api/user/autoMigrate"] = gsharp.RouterOptions{Method: gsharp.Post, Handler: user.AutoMigrateHandler, Authorize: false}
	routerMap["api/user/create"] = gsharp.RouterOptions{Method: gsharp.Post, Handler: user.CreateUserManageHandler, Authorize: false}
	routerMap["api/user/update"] = gsharp.RouterOptions{Method: gsharp.Post, Handler: user.UpdateUserManageHandler, Authorize: false}
	routerMap["api/user/query"] = gsharp.RouterOptions{Method: gsharp.Get, Handler: user.QueryUserManageListHandler, Authorize: false}
	routerMap["api/user/query/:page/:size"] = gsharp.RouterOptions{Method: gsharp.Get, Handler: user.QueryUserManagePageHandler, Authorize: false}

	return routerMap
}
