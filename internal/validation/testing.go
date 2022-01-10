package validation

import (
	"testing"
)

type TestUserData struct {
	Status          string `json:"status" validate:"required"`
	Email           string `json:"email" validate:"required,email"`
	FirstName       string `json:"first_name" validate:"required,min=2,max=50"`
	LastName        string `json:"last_name" validate:"required,min=2,max=50"`
	Password        string `json:"password" validate:"required,min=8"`
	ConfirmPassword string `json:"confirm_password" validate:"required,min=8,eqfield=Password"`
}

func TestValidationData(t *testing.T) *TestUserData {
	t.Helper()
	return &TestUserData{
		Status:          "active",
		Email:          "user@info.com",
		FirstName:      "User",
		LastName:       "Admin",
		Password:       "123456789",
		ConfirmPassword: "123456789",
	}
}


