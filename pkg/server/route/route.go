package route

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Route struct {
	Method      string
	Pattern     string
	HandlerFunc gin.HandlerFunc
}

type Routes []Route

func RegisterRoutes(api *gin.RouterGroup, routes Routes) {
	for _, route := range routes {
		switch route.Method {
		case http.MethodGet:
			api.GET(route.Pattern, route.HandlerFunc)
		case http.MethodPost:
			api.POST(route.Pattern, route.HandlerFunc)
		case http.MethodPut:
			api.PUT(route.Pattern, route.HandlerFunc)
		case http.MethodPatch:
			api.PATCH(route.Pattern, route.HandlerFunc)
		case http.MethodDelete:
			api.DELETE(route.Pattern, route.HandlerFunc)
		}
	}

}
