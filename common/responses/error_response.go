package responses

import "github.com/gin-gonic/gin"

func SendErrorResponse(ctx *gin.Context, code int, err error) {
	SendJsonResponse(ctx, code, err.Error(), struct{}{})
}
