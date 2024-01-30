package requests

import (
	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/contracts/validation"
)

type AppInfoStoreRequest struct {
	AppName    string `form:"app_name" json:"app_name"`
	AppVer     string `form:"app_ver" json:"app_ver"`
	AppDesc    string `form:"app_desc" json:"app_desc"`
	AppLogo    string `form:"app_logo" json:"app_logo"`
	AppTheme   string `form:"app_theme" json:"app_theme"`
	AppColor   string `form:"app_color" json:"app_color"`
	AppBg      string `form:"app_bg" json:"app_bg"`
	AppBgNav   string `form:"app_bg_nav" json:"app_bg_nav"`
	AppUrl     string `form:"app_url" json:"app_url"`
	AppCompany string `form:"app_company" json:"app_company"`
	AppSlogan  string `form:"app_slogan" json:"app_slogan"`
	AppAddress string `form:"app_address" json:"app_address"`
	AppWa      string `form:"app_wa" json:"app_wa"`
	AppTw      string `form:"app_tw" json:"app_tw"`
	AppIg      string `form:"app_ig" json:"app_ig"`
}

func (r *AppInfoStoreRequest) Authorize(ctx http.Context) error {
	return nil
}

func (r *AppInfoStoreRequest) Rules(ctx http.Context) map[string]string {
	return map[string]string{
		"app_name":    "required",
		"app_ver":     "required",
		"app_desc":    "required",
		"app_theme":   "required",
		"app_color":   "required",
		"app_company": "required",
		"app_slogan":  "required",
	}
}

func (r *AppInfoStoreRequest) Messages(ctx http.Context) map[string]string {
	return map[string]string{
		"app_name.required":    "Nama aplikasi tidak boleh kosong",
		"app_ver.required":     "Versi aplikasi tidak boleh kosong",
		"app_desc.required":    "Deskripsi aplikasi tidak boleh kosong",
		"app_color.required":   "Pilihan warna tidak boleh kosong",
		"app_theme.required":   "Pilihan theme tidak boleh kosong",
		"app_company.required": "Nama perusahaan tidak boleh kosong",
		"app_slogan.required":  "Slogan/Visi dan Misi Perusahaan tidak boleh kosong",
	}
}

func (r *AppInfoStoreRequest) Attributes(ctx http.Context) map[string]string {
	return map[string]string{}
}

func (r *AppInfoStoreRequest) PrepareForValidation(ctx http.Context, data validation.Data) error {
	return nil
}
