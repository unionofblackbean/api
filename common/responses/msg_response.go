package responses

import "github.com/gin-gonic/gin"

func SendMsgResponse(ctx *gin.Context,
	code int, msg string,
) {
	SendJsonResponse(ctx, code, msg, struct{}{})
}
