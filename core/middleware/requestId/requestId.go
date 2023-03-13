package requestid

import (
	"context"

	"github.com/sunweiwe/paddle/core/errors"
)

const HeaderXRequestID = "X-Request-ID"

func FromContext(ctx context.Context) (string, error) {
	rid, ok := ctx.Value(HeaderXRequestID).(string)
	if !ok {
		return "", errors.ErrFailedToGetRequestID
	}
	return rid, nil
}
