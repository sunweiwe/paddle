package manager

import (
	"context"

	"github.com/sunweiwe/paddle/lib/q"
	"github.com/sunweiwe/paddle/pkg/user/dao"
	"github.com/sunweiwe/paddle/pkg/user/model"
	"gorm.io/gorm"
)

type Manager interface {
	// Create(ctx context.Context, user *model.User) (*model.User, error)
	List(ctx context.Context, query *q.Query) (int64, []*model.User, error)
	// GetUserByID(ctx context.Context, userID uint) (*model.User, error)
	// UpdateByID(ctx context.Context, id uint, db *model.User) (*model.User, error)
	// DeleteUser(ctx context.Context, id uint) error
}

type manager struct {
	dao dao.DAO
}

func New(db *gorm.DB) Manager {
	return &manager{dao: dao.NewDAO(db)}
}

func (m *manager) List(ctx context.Context, query *q.Query) (int64, []*model.User, error) {
	return m.dao.List(ctx, query)
}
