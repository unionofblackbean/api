package responses

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SendMethodNotAllowed(ctx *gin.Context) {
	SendErrorResponse(ctx, http.StatusMethodNotAllowed, errors.New("method not allowed"))
}
