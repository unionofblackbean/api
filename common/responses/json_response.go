package responses

import "github.com/gin-gonic/gin"

type jsonResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"msg"`
	Data    interface{} `json:"data"`
}

func SendJsonResponse(ctx *gin.Context,
	code int, msg string, data interface{},
) {
	ctx.JSON(code, &jsonResponse{
		Code:    code,
		Message: msg,
		Data:    data,
	})
}
