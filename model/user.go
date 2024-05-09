package model

import (
	"time"
)

type User struct {
	UserId                     int64     `json:"user_id"`
	Username                   string    `json:"username"`
	Email                      string    `json:"email"`
	FullName                   string    `json:"full_name"`
	Password                   string    `json:"password"`
	RoleDescription            string    `json:"role_description"`
	RoleId                     int16     `json:"role_id"`
	IsActive                   bool      `json:"is_active"`
	IsBlocked                  bool      `json:"is_blocked"`
	LastSucceedLoginDate       time.Time `json:"last_succeed_login_date"`
	RefreshTokenExpirationDate time.Time `json:"refresh_token_expiration_date"`
	RefreshToken               string    `json:"refresh_token"`
	CreateBy                   int64     `json:"create_by"`
	CreateDate                 time.Time `json:"create_date"`
	UpdateBy                   int64     `json:"update_by"`
	UpdateDate                 time.Time `json:"update_date"`
}
