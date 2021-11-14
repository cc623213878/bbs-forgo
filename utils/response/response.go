package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func NewResponse(code int, msg string, data interface{}) *Response {
	return &Response{Code: code, Msg: msg, Data: data}
}

func Success(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, NewResponse(0, "success", data))
}

func Error(c *gin.Context, msg string) {
	c.JSON(http.StatusOK, NewResponse(1, msg, ""))
}
