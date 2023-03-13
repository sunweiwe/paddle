package user

import (
	"time"

	"github.com/sunweiwe/paddle/pkg/user/model"
)

type User struct {
	ID        uint      `json:"id,omitempty"`
	Name      string    `json:"name,omitempty"`
	FullName  string    `json:"fullName,omitempty"`
	Email     string    `json:"email,omitempty"`
	Admin     bool      `json:"Admin"`
	Permitted bool      `json:"permitted"`
	Phone     string    `json:"phone,omitempty"`
	UpdatedAt time.Time `json:"updatedAt,omitempty"`
	CreatedAt time.Time `json:"createdAt,omitempty"`
}

func ofUser(u *model.User) *User {
	return &User{
		ID:        u.ID,
		Name:      u.Name,
		FullName:  u.FullName,
		Email:     u.Email,
		Admin:     u.Admin,
		Permitted: u.Permitted,
		Phone:     u.Phone,
		UpdatedAt: u.UpdatedAt,
		CreatedAt: u.CreatedAt,
	}
}

func ofUsers(users []*model.User) []*User {
	resp := make([]*User, 0, len(users))
	for _, u := range users {
		resp = append(resp, ofUser(u))
	}
	return resp
}
