package dao

import (
	"context"
	"fmt"

	"github.com/sunweiwe/paddle/lib/q"
	"github.com/sunweiwe/paddle/pkg/user/model"
	"gorm.io/gorm"
)

type DAO interface {
	List(ctx context.Context, query *q.Query) (int64, []*model.User, error)
}

type dao struct{ db *gorm.DB }

func NewDAO(db *gorm.DB) DAO {
	return &dao{db: db}
}

func (d *dao) List(ctx context.Context, query *q.Query) (int64, []*model.User, error) {
	var users []*model.User
	tx := d.db.Table("tb_user")

	if query != nil {
		for k, v := range query.Keywords {
			switch k {
			case "filter":
				tx = tx.Where("name like ?", fmt.Sprintf("%%%v%%", v))
			case "userType":
				tx = tx.Where("user_type in ?", v)
			}
		}
	}

	var total int64
	tx.Count(&total)

	if query != nil {
		tx = tx.Limit(query.Limit()).Offset(query.Offset())
	}
	ret := tx.Scan(&users)

	err := ret.Error
	if err != nil {
		return 0, nil, err
	}

	if ret.RowsAffected == 0 {
		return 0, make([]*model.User, 0), nil
	}

	return total, users, nil
}
