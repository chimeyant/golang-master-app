package auth

import (
	"goravel/app/http/requests"
	"goravel/app/lang"
	"goravel/app/models"
	"goravel/app/services"
	http2 "net/http"

	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/facades"
)

type UserAuthController struct {
	//Dependent services
	AuthService *services.AuthService
	UserService *services.UserService
}

type UserInfoType struct {
	Name       string `json:"name"`
	Email      string `json:"email"`
	Authent    string `json:"authent"`
	AvatarPath string `json:"avatar_path"`
}

func NewUserAuthController(authService *services.AuthService, userService *services.UserService) *UserAuthController {
	return &UserAuthController{
		//Inject services
		AuthService: authService,
		UserService: userService,
	}
}

func (r *UserAuthController) Login(ctx http.Context) http.Response {
	var req requests.UserLoginRequest

	errors, err := ctx.Request().ValidateRequest(&req)

	if err != nil {
		return ctx.Response().Json(http2.StatusInternalServerError, http.Json{
			"code":    500,
			"success": false,
			"message": lang.MsgError,
		})
	}

	if errors != nil {
		return ctx.Response().Json(http.StatusUnprocessableEntity, http.Json{
			"code":    422,
			"success": false,
			"message": lang.MsgError,
			"errors":  errors.All(),
		})
	}

	result, user, userdisplay := r.AuthService.LoginProcess(req)

	if !result {
		return ctx.Response().Json(http.StatusUnauthorized, http.Json{
			"code":    401,
			"success": false,
			"message": lang.MsgLoginError,
		})
	}

	token, _ := facades.Auth().Login(ctx, user)

	payload, _ := facades.Auth().Parse(ctx, token)

	return ctx.Response().Json(http.StatusOK, http.Json{
		"code":    200,
		"success": true,
		"response": http.Json{
			"token": http.Json{
				"type":     "Bearer",
				"token":    token,
				"expireAt": payload.ExpireAt,
			},
			"user": userdisplay,
		},
	})
}

func (r *UserAuthController) Logout(ctx http.Context) http.Response {
	var user models.User

	err := facades.Auth().User(ctx, &user)

	if err != nil {
		facades.Log().Error(err)
	}

	return ctx.Response().Json(http.StatusOK, http.Json{
		"code":    200,
		"success": true,
		"message": "Berhasil",
		"user":    user,
	})
}

func (r *UserAuthController) UserInfo(ctx http.Context) http.Response {

	var model models.User

	err := facades.Auth().User(ctx, &model)

	if err != nil {
		facades.Log().Error(err)
	}

	user := UserInfoType{
		Name:       model.Name,
		Email:      model.Email,
		Authent:    model.Authent,
		AvatarPath: model.Avatar,
	}

	return ctx.Response().Json(http.StatusOK, user)
}

func (r *UserAuthController) Register(ctx http.Context) http.Response {
	var req requests.UserRegisterRequest

	errors, err := ctx.Request().ValidateRequest(&req)

	if err != nil {
		return ctx.Response().Status(500).Json(http.Json{
			"code":    500,
			"success": false,
			"message": lang.MsgError,
		})
	}

	if errors != nil {
		return ctx.Response().Status(422).Json(http.Json{
			"code":    422,
			"success": false,
			"message": lang.MsgError,
			"errors":  errors.All(),
		})
	}

	checkUser := r.UserService.CheckEmail(req.Email)

	if !checkUser {
		return ctx.Response().Json(422, http.Json{
			"code":    422,
			"success": false,
			"message": lang.MsgEmailRegistered,
		})
	}

	r.UserService.Register(req)

	return ctx.Response().Status(200).Json(http.Json{
		"code":    200,
		"success": true,
		"message": lang.MsgRegisterUserSuccess,
	})
}
