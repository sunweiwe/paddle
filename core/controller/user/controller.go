package user

import (
	"context"

	"github.com/sunweiwe/paddle/lib/q"
	"github.com/sunweiwe/paddle/pkg/parameter"
	"github.com/sunweiwe/paddle/pkg/user/manager"
)

type Controller interface {
	List(ctx context.Context, query *q.Query) (int64, []*User, error)
}

type controller struct {
	user manager.Manager
}

func NewController(param *parameter.Parameter) Controller {
	return &controller{
		user: param.UserManager,
	}
}

func (c *controller) List(ctx context.Context, query *q.Query) (int64, []*User, error) {

	total, users, err := c.user.List(ctx, query)
	if err != nil {
		return 0, nil, err
	}

	return total, ofUsers(users), nil
}
