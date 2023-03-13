package model

import "github.com/sunweiwe/paddle/pkg/server/global"

type User struct {
	global.Model

	Name      string
	FullName  string
	Email     string
	Phone     string
	UserType  uint
	OidcID    string `gorm:"column:oidc_id"`
	OidcType  string `gorm:"column:oidc_type"`
	Admin     bool
	Permitted bool
}

type UserBasic struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func ToUser(user *User) *UserBasic {
	if user == nil {
		return nil
	}

	return &UserBasic{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}
}
