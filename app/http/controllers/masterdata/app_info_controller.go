package masterdata

import (
	"github.com/goravel/framework/contracts/http"
	"goravel/app/http/requests"
	"goravel/app/lang"
	"goravel/app/services"
)

type AppInfoController struct {
	//Dependent services
	AppInfoService *services.AppInfoService
}

func NewAppInfoController(appInfoService *services.AppInfoService) *AppInfoController {
	return &AppInfoController{
		//Inject services
		AppInfoService: appInfoService,
	}
}

func (r *AppInfoController) Index(ctx http.Context) http.Response {
	result := r.AppInfoService.Fetch()

	return ctx.Response().Success().Json(result)
}

func (r *AppInfoController) Show(ctx http.Context) http.Response {
	result := r.AppInfoService.ShowInfo()

	return ctx.Response().Json(http.StatusOK, result)
}

func (r *AppInfoController) ShowInfo(ctx http.Context) http.Response {
	result := r.AppInfoService.ShowInfo()

	return ctx.Response().Json(http.StatusOK, result)
}

func (r *AppInfoController) Store(ctx http.Context) http.Response {
	var req requests.AppInfoStoreRequest

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

	r.AppInfoService.Store(req)

	return ctx.Response().Json(http.StatusOK, http.Json{
		"code":    200,
		"success": true,
		"message": lang.MsgStoreSuccess,
	})
}
