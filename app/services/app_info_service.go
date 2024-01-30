package services

import (
	"github.com/google/uuid"
	"github.com/goravel/framework/facades"
	"goravel/app/http/requests"
	"goravel/app/models"
)

type AppInfoService struct {
}

type Theme struct {
	Mode  string `json:"mode"`
	Color string `json:"color"`
}

type ResultResponse struct {
	Uuid       string `json:"id"`
	AppName    string `json:"app_name"`
	AppVer     string `json:"app_ver"`
	AppDesc    string `json:"app_desc"`
	AppColor   string `json:"app_color"`
	AppTheme   Theme  `json:"app_theme"`
	AppCompany string `json:"app_company"`
	AppSlogan  string `json:"app_slogan"`
	AppAddress string `json:"app_address"`
}

func NewAppInfoService() *AppInfoService {
	return &AppInfoService{
		// Initialize any dependencies if needed
	}
}

func (s *AppInfoService) Fetch() []ResultResponse {
	var model []models.AppInfo

	var results []ResultResponse

	err := facades.Orm().Query().Model(&model).Order("id asc").Get(&results)

	if err != nil {
		facades.Log().Error(err)
	}

	return results
}

func (s *AppInfoService) Show(uuid string) models.AppInfo {
	var model models.AppInfo

	err := facades.Orm().Query().Where("uuid", uuid).First(&model)

	if err != nil {
		facades.Log().Error(err)
	}

	return model

}

func (s *AppInfoService) ShowInfo() ResultResponse {
	var model models.AppInfo

	err := facades.Orm().Query().Order("id asc").First(&model)

	if err != nil {
		facades.Log().Error(err)
	}

	theme := Theme{
		Mode:  "dark",
		Color: model.AppTheme,
	}

	result := ResultResponse{
		Uuid:       model.Uuid,
		AppName:    model.AppName,
		AppVer:     model.AppVer,
		AppDesc:    model.AppDesc,
		AppColor:   model.AppColor,
		AppTheme:   theme,
		AppCompany: model.AppCompany,
		AppSlogan:  model.AppSlogan,
		AppAddress: model.AppAddress,
	}

	return result
}

func (s *AppInfoService) Store(payload requests.AppInfoStoreRequest) {
	var model models.AppInfo

	model.Uuid = uuid.NewString()
	model.AppName = payload.AppName
	model.AppVer = payload.AppVer
	model.AppDesc = payload.AppDesc
	model.AppLogo = payload.AppLogo
	model.AppTheme = payload.AppTheme
	model.AppColor = payload.AppColor
	model.AppBg = payload.AppBg
	model.AppBgNav = payload.AppBgNav
	model.AppUrl = payload.AppUrl
	model.AppCompany = payload.AppCompany
	model.AppSlogan = payload.AppSlogan
	model.AppAddress = payload.AppAddress
	model.AppWa = payload.AppWa
	model.AppTw = payload.AppTw
	model.AppIg = payload.AppIg

	err := facades.Orm().Query().Save(&model)

	if err != nil {
		facades.Log().Error(err)
	}
}

func (s *AppInfoService) update(payload requests.AppInfoUpdateRequest, uuid string) {
	var model models.AppInfo

	err := facades.Orm().Query().Where("uuid", uuid).First(&model)

	if err != nil {
		facades.Log().Error(err)
	}

	model.AppName = payload.AppName
	model.AppVer = payload.AppVer
	model.AppDesc = payload.AppDesc
	model.AppLogo = payload.AppLogo
	model.AppTheme = payload.AppTheme
	model.AppColor = payload.AppColor
	model.AppBg = payload.AppBg
	model.AppBgNav = payload.AppBgNav
	model.AppUrl = payload.AppUrl
	model.AppCompany = payload.AppCompany
	model.AppSlogan = payload.AppSlogan
	model.AppAddress = payload.AppAddress
	model.AppWa = payload.AppWa
	model.AppTw = payload.AppTw
	model.AppIg = payload.AppIg

	err = facades.Orm().Query().Save(&model)

	if err != nil {
		facades.Log().Error(err)
	}
}
