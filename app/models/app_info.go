package models

import (
	"github.com/goravel/framework/database/orm"
)

type AppInfo struct {
	orm.Model
	Uuid       string `json:"id"`
	AppName    string `json:"app_name"`
	AppVer     string
	AppDesc    string
	AppLogo    string
	AppTheme   string
	AppColor   string
	AppBg      string
	AppBgNav   string
	AppUrl     string
	AppCompany string
	AppSlogan  string
	AppAddress string
	AppWa      string
	AppTw      string
	AppIg      string
	orm.SoftDeletes
}

func (a *AppInfo) TableName() string {
	return "app_infos"
}
