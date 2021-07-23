package responses

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SendTooManyRequestsResponse(ctx *gin.Context) {
	SendErrorResponse(ctx, http.StatusTooManyRequests, errors.New("too many requests"))
}
