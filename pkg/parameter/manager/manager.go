package manager

import (
	"gorm.io/gorm"

	userManager "github.com/sunweiwe/paddle/pkg/user/manager"
)

type Manager struct {
	UserManager userManager.Manager
}

func CreateManager(db *gorm.DB) *Manager {
	return &Manager{
		UserManager: userManager.New(db),
	}
}
