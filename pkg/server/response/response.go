package response

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	requestid "github.com/sunweiwe/paddle/core/middleware/requestId"
	"github.com/sunweiwe/paddle/pkg/server/rpcerror"
)

type DataWithCount struct {
	Total int64       `json:"total"`
	Data  interface{} `json:"data"`
}

type Response struct {
	ErrorCode    string      `json:"errorCode,omitempty"`
	ErrorMessage string      `json:"errorMessage,omitempty"`
	Data         interface{} `json:"data,omitempty"`
	RequestID    string      `json:"requestID,omitempty"`
}

func Abort(c *gin.Context, httpCode int, errorCode, errorMessage string) {
	rid, err := requestid.FromContext(c)
	if err != nil {
		log.Fatal(c, "error to get requestID from context, err: %v", err)
	}

	c.JSON(httpCode, &Response{
		ErrorCode:    errorCode,
		ErrorMessage: errorMessage,
		RequestID:    rid,
	})
	c.Abort()
}

func SuccessWithData(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, NewResponseWithData(data))
}

func AbortWithRPCError(c *gin.Context, rpcError rpcerror.RPCError) {
	Abort(c, rpcError.HTTPCode, string(rpcError.ErrorCode), rpcError.ErrorMessage)
}

func NewResponseWithData(data interface{}) *Response {
	return &Response{
		Data: data,
	}
}
