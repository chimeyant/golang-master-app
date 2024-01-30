package requests

import (
	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/contracts/validation"
)

type UserLoginRequest struct {
	Email    string `form:"email" json:"email"`
	Password string `form:"password" json:"password"`
}

func (r *UserLoginRequest) Authorize(ctx http.Context) error {
	return nil
}

func (r *UserLoginRequest) Rules(ctx http.Context) map[string]string {
	return map[string]string{
		"email":    "required",
		"password": "required",
	}
}

func (r *UserLoginRequest) Messages(ctx http.Context) map[string]string {
	return map[string]string{
		"email.required":    "Email tidak boleh kosong",
		"password.required": "Kata sandi tidak boleh kosong",
	}
}

func (r *UserLoginRequest) Attributes(ctx http.Context) map[string]string {
	return map[string]string{}
}

func (r *UserLoginRequest) PrepareForValidation(ctx http.Context, data validation.Data) error {
	return nil
}
