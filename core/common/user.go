package common

import (
	"context"

	"github.com/sunweiwe/paddle/core/errors"
	"github.com/sunweiwe/paddle/pkg/auth/user"
)

const (
	contextUserKey = "contextUser"

	AuthorizationHeaderKey = "Authorization"
	TokenHeaderValuePrefix = "Bearer"
)

func UserFromContext(ctx context.Context) (user.User, error) {
	u, ok := ctx.Value(contextUserKey).(user.User)
	if !ok {
		return nil, errors.ErrFailedToGetUser
	}
	return u, nil
}
