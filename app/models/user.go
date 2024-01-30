package models

import (
	"github.com/goravel/framework/database/orm"
)

type User struct {
	orm.Model
	Uuid          string `json:"id"`
	Name          string `json:"name"`
	Email         string
	Password      string
	Authent       string
	Avatar        string
	GoogleId      string
	RememberToken string
	Status        int
	orm.SoftDeletes
}

func (r *User) TableName() string {
	return "users"
}
