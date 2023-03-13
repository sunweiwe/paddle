package user

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"
	"github.com/sunweiwe/paddle/core/controller/user"
	"github.com/sunweiwe/paddle/lib/q"
	"github.com/sunweiwe/paddle/pkg/server/response"
	"github.com/sunweiwe/paddle/pkg/server/rpcerror"
)

type API struct {
	userCtr user.Controller
	store   sessions.Store
}

func NewAPI(ctr user.Controller, store sessions.Store) *API {
	return &API{
		userCtr: ctr,
		store:   store,
	}
}

func (api *API) List(c *gin.Context) {
	queryName := c.Query("filter")
	keywords := q.KeyWords{}
	if queryName != "" {
		keywords["filter"] = queryName
	}

	var userTypes []int
	queryUserTypes := c.QueryArray("userType")
	for _, s := range queryUserTypes {
		userType, err := strconv.Atoi(s)
		if err != nil {
			response.AbortWithRPCError(c,
				rpcerror.ParameterError.WithErrorMessageFormat("invalid user type: %s,err: %v", s, err))
			return
		}
		userTypes = append(userTypes, userType)
	}

	if len(userTypes) == 0 {
		userTypes = append(userTypes, 0)
	}
	keywords["userType"] = userTypes

	query := q.New(keywords).WithPagination(c)
	total, users, err := api.userCtr.List(c, query)
	if err != nil {
		response.AbortWithRPCError(c,
			rpcerror.InternalError.WithErrorMessageFormat("Failed to get users: "+"err = %v", err))
		return
	}

	response.SuccessWithData(c,
		response.DataWithCount{
			Total: total,
			Data:  users,
		})
}
