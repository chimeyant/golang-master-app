package requests

import (
	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/contracts/validation"
)

type UserRegisterRequest struct {
	Name     string `form:"name" json:"name"`
	Email    string `form:"email" json:"email"`
	Password string `form:"password" json:"password"`
}

func (r *UserRegisterRequest) Authorize(ctx http.Context) error {
	return nil
}

func (r *UserRegisterRequest) Rules(ctx http.Context) map[string]string {
	return map[string]string{
		"name":     "required",
		"email":    "required|email",
		"password": "required",
	}
}

func (r *UserRegisterRequest) Messages(ctx http.Context) map[string]string {
	return map[string]string{
		"name.required":     "Nama pengguna tidak boleh kosong...!",
		"email.required":    "Email tidak boleh kosong..!",
		"email.email":       "Email tidak benar..!",
		"password.required": "Kata sandi tidak boleh kosong...!",
	}
}

func (r *UserRegisterRequest) Attributes(ctx http.Context) map[string]string {
	return map[string]string{}
}

func (r *UserRegisterRequest) PrepareForValidation(ctx http.Context, data validation.Data) error {
	return nil
}
