package models

import (
	"context"

	"time"
)

type User struct {
	ID              uint      `json:"id"`
	FirstName       string    `json:"firstName"`
	LastName        string    `json:"lastName"`
	Password        string    `json:"password"`
	AvatarLink      string    `json:"avatarLink"`
	Gender          bool      `json:"gender"`
	Age             int       `json:"age"`
	PhoneNumber     string    `json:"phoneNumber"`
	CityOfResidence string    `json:"cityOfResidence"`
	Description     string    `json:"description"`
	Email           string    `json:"email"`
	IsEmailVerified bool      `json:"isEmailVerified"`
	Username        string    `json:"username"`
	IsPrivate       bool      `json:"isPrivate"`
	ReceivePush     bool      `json:"receivePush"`
	RoleID          uint      `json:"roleId"`
	CreatedAt       time.Time `json:"createdAt"`
}

type UserRequest struct {
	Email       string `json:"email"`
	Password    string `json:"password"`
	PhoneNumber string `json:"phoneNumber"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Password struct {
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirmPassword"`
}

type UserRepository interface {
	GetUserByEmail(c context.Context, email string) (User, error)
	GetUserByID(c context.Context, userID int) (User, error)
	GetProfile(c context.Context, userID int) (User, error)

	CreateUser(c context.Context, user User) (int, error)
}
