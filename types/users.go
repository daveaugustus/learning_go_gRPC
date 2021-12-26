package types

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

// TempUser - the temp user for creating a new user
type TempUser struct {
	FirstName       string `json:"first_name" validate:"required,gte=4"`
	LastName        string `json:"last_name" validate:"required,gte=4"`
	Email           string `json:"email" validate:"required,contains=@"`
	Password        string `json:"-" validate:"required,gte=8"`
	ConfirmPassword string `json:"-" validate:"required,gte=8"`
}

// User - the user in the system
type User struct {
	ID        int64  `json:"id" xorm:"'id' pk autoincr" schema:"id"`
	FirstName string `json:"first_name" xorm:"first_name" schema:"first_name" validate:"required,gte=4"`
	LastName  string `json:"last_name" xorm:"last_name" schema:"last_name" validate:"required,gte=4"`
	Email     string `json:"email" xorm:"email" schema:"email" validate:"required,contains=@"`
	Password  string `json:"-" xorm:"password" schema:"password" validate:"required,gte=8"`
	Visible   bool   `json:"visible" xorm:"visible" schema:"visible"`
}

// TableName - the table when using xorm
func (u *User) TableName() string {
	return "users"
}

// NewUser - creates a new uesr from a temp user
func NewUser(nu *TempUser) (user *User, err error) {
	if nu.Password != nu.ConfirmPassword {
		err = fmt.Errorf("password and confirm password do not match")
		return
	}
	user = &User{
		// ID - is auto generated
		Email:     nu.Email,
		FirstName: nu.FirstName,
		LastName:  nu.LastName,
		Visible:   true,
	}

	if err = user.SetPassword(nu.Password); err != nil {
		return
	}

	return
}

// SetPassword - use bcrypt to set the password hash
func (u *User) SetPassword(password string) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	u.Password = string(hash)
	return nil
}

// Authenticate - authenticates a password set against the stored hash
func (u *User) Authenticate(password string) error {
	if !u.Visible {
		return fmt.Errorf("user is inactive")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password)); err != nil {
		return err
	}
	return nil
}
