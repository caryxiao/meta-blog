package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Code    int         `json:"code"`
	Msg     string      `json:"msg"`
	Data    interface{} `json:"data"`
	TraceId string      `json:"trace_id"`
}

func (r *Response) Success(c *gin.Context, data interface{}) {
	traceId, _ := c.Get("traceId")
	response := Response{
		Code:    0,
		Msg:     "success",
		Data:    data,
		TraceId: traceId.(string),
	}
	c.JSON(http.StatusOK, response)
}

func (r *Response) Fail(c *gin.Context, code int, msg string) {
	traceId, _ := c.Get("traceId")
	response := Response{
		Code:    code,
		Msg:     msg,
		Data:    nil,
		TraceId: traceId.(string),
	}
	c.JSON(http.StatusOK, response)
}
