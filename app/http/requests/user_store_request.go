package requests

import (
	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/contracts/validation"
)

type UserStoreRequest struct {
	Name     string `form:"name" json:"name"`
	Email    string `form:"email" json:"email"`
	Password string `form:"password" json:"password"`
	Authent  string `form:"authent" json:"authent"`
	Status   string `form:"status" json:"status"`
}

func (r *UserStoreRequest) Authorize(ctx http.Context) error {
	return nil
}

func (r *UserStoreRequest) Rules(ctx http.Context) map[string]string {
	return map[string]string{
		"name":     "required",
		"email":    "required|email",
		"password": "required",
		"authent":  "required",
		"status":   "required",
	}
}

func (r *UserStoreRequest) Messages(ctx http.Context) map[string]string {
	return map[string]string{
		"name.required":  "Nama pengguna tidak boleh kosong..!",
		"email.required": "Email tidak boleh kosong..!",
		"email.email":    "Email tidak benar..!",
	}
}

func (r *UserStoreRequest) Attributes(ctx http.Context) map[string]string {
	return map[string]string{}
}

func (r *UserStoreRequest) PrepareForValidation(ctx http.Context, data validation.Data) error {
	return nil
}
