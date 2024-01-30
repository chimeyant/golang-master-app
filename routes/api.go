package routes

import (
	"goravel/app/http/controllers/auth"
	"goravel/app/http/controllers/masterdata"
	"goravel/app/http/controllers/utility"
	"goravel/app/http/middleware"
	"goravel/app/services"

	"github.com/goravel/framework/contracts/route"
	"github.com/goravel/framework/facades"
)

func Api() {
	facades.Route().Static("static", "./public")

	//begin declare service
	appInfoService := services.NewAppInfoService()
	userService := services.NewUserService()
	userAuthService := services.NewAuthService()

	//begin declare controller
	appInfoController := masterdata.NewAppInfoController(appInfoService)
	userAuthController := auth.NewUserAuthController(userAuthService, userService)
	userController := utility.NewUserController(userService)
	userMenus := auth.NewMenusController()

	//begin add route

	facades.Route().Prefix("api/v2").Group(func(router route.Router) {
		//base route
		router.Get("app-info", appInfoController.ShowInfo)

		//Route Auth
		router.Prefix("auth").Group(func(router route.Router) {
			router.Post("register", userAuthController.Register)
			router.Post("token", userAuthController.Login)
			router.Middleware(middleware.Auth()).Post("logout", userAuthController.Logout)
			router.Middleware(middleware.Auth()).Get("user-info", userAuthController.UserInfo)
			router.Middleware(middleware.Auth()).Get("menus", userMenus.Index)
		})

		//Route Master Data
		router.Middleware(middleware.Auth()).Prefix("master-data").Group(func(router route.Router) {
			router.Get("app-info", appInfoController.Index)
			router.Post("app-info", appInfoController.Store)
		})

		router.Get("users", userController.Index)

	})

}
