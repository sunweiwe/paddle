package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sunweiwe/paddle/pkg/server/route"
)

func (api *API) RegisterRoutes(engine *gin.Engine) {
	coreGroup := engine.Group("/api/core/v1/users")
	var coreRoutes = route.Routes{
		{
			Method:      http.MethodGet,
			HandlerFunc: api.List,
		},
	}
	route.RegisterRoutes(coreGroup, coreRoutes)
}
