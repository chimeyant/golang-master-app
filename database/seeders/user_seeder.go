package seeders

import (
	"github.com/google/uuid"
	"github.com/goravel/framework/facades"
	"goravel/app/helpers"
	"goravel/app/models"
)

type UserSeeder struct {
}

// Signature The name and signature of the seeder.
func (s *UserSeeder) Signature() string {
	return "UserSeeder"
}

// Run executes the seeder logic.
func (s *UserSeeder) Run() error {
	hashPassword, _ := helpers.HashPassword("rahasia")
	user := models.User{
		Uuid:     uuid.NewString(),
		Name:     "Administrator",
		Email:    "admin@ants.loc",
		Password: hashPassword,
		Authent:  "superadmin",
		Avatar:   "/images/user.png",
		Status:   1,
	}
	return facades.Orm().Query().Create(&user)
}
