package requests

import (
	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/contracts/validation"
)

type UserUpdateRequest struct {
	Name     string `form:"name" json:"name"`
	Email    string `form:"email" json:"email"`
	Password string `form:"password" json:"password"`
	Authent  string `form:"authent" json:"authent"`
	Status   string `form:"status" json:"status"`
}

func (r *UserUpdateRequest) Authorize(ctx http.Context) error {
	return nil
}

func (r *UserUpdateRequest) Rules(ctx http.Context) map[string]string {
	return map[string]string{
		"name":     "required",
		"email":    "required|email",
		"password": "required",
		"authent":  "required",
		"status":   "required",
	}
}

func (r *UserUpdateRequest) Messages(ctx http.Context) map[string]string {
	return map[string]string{}
}

func (r *UserUpdateRequest) Attributes(ctx http.Context) map[string]string {
	return map[string]string{}
}

func (r *UserUpdateRequest) PrepareForValidation(ctx http.Context, data validation.Data) error {
	return nil
}
