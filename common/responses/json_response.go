package responses

import (
	"github.com/gin-gonic/gin"
	"time"
)

type jsonResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"msg"`
	Time    int64       `json:"time"`
	Data    interface{} `json:"data"`
}

func SendJsonResponse(ctx *gin.Context,
	code int, msg string, data interface{},
) {
	ctx.JSON(code, &jsonResponse{
		Code:    code,
		Message: msg,
		Time:    time.Now().UTC().Unix(),
		Data:    data,
	})
}
