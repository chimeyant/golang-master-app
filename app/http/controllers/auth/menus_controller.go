package auth

import (
	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/facades"
	"goravel/app/models"
)

type MenusController struct {
	//Dependent services
}

type SubMenuProps struct {
	Type  string `json:"type"`
	Icon  string `json:"icon"`
	Route string `json:"route"`
}

type SubMenuItems struct {
	Title string       `json:"title"`
	Icon  string       `json:"icon"`
	Props SubMenuProps `json:"props"`
}

type MenuProps struct {
	Type     string         `json:"type"`
	Icon     string         `json:"icon"`
	Route    string         `json:"route"`
	Submenus []SubMenuItems `json:"submenus"`
}

type MenuItem struct {
	Title string    `json:"title"`
	Props MenuProps `json:"props"`
	Route string    `json:"route"`
}

func NewMenusController() *MenusController {
	return &MenusController{
		//Inject services
	}
}

func (r *MenusController) Index(ctx http.Context) http.Response {
	var user models.User

	err := facades.Auth().User(ctx, &user)

	if err != nil {
		facades.Log().Error(err)
	}

	var authent string = user.Authent

	if authent == "superadmin" {
		menus := []MenuItem{
			{
				Title: "Dashboard",
				Props: MenuProps{
					Type:  "item",
					Icon:  "mdi-view-dashboard",
					Route: "/auth/dashboard",
				},
				Route: "/auth/dashboard",
			}, {
				Title: "Aplikasi",
				Props: MenuProps{
					Type: "subheader",
					Icon: "mdi-monitor-dashboard",
				},
			},
			{
				Title: "Data Master",
				Props: MenuProps{
					Type: "group",
					Icon: "mdi-database-settings",
					Submenus: []SubMenuItems{
						{
							Title: "Konfigurasi Aplikasi",
							Icon:  "",
							Props: SubMenuProps{
								Type:  "item",
								Route: "/auth/master-app-info",
							},
						},
					},
				},
			}, {
				Title: "Utilitas",
				Props: MenuProps{
					Type: "subheader",
					Icon: "mdi-tools",
				},
			}, {
				Title: "Manajemen Pengguna",
				Props: MenuProps{
					Type:  "item",
					Icon:  "mdi-account-group",
					Route: "/auth/utility-user-manajemen",
				},
			}, {
				Title: "Profil Pengguna",
				Props: MenuProps{
					Type:  "item",
					Icon:  "mdi-account-details",
					Route: "/auth/utility-user-profile",
				},
			},
			{
				Title: "Ubah Kata Sandi",
				Props: MenuProps{
					Type:  "item",
					Icon:  "mdi-key",
					Route: "/auth/utility-user-change-pwd",
				},
			},
		}
		return ctx.Response().Status(200).Json(menus)
	}

	return nil

}
