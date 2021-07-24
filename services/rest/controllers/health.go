package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/leungyauming/api/common/responses"
	"net/http"
	"time"
)

func Health(ctx *gin.Context) {
	switch ctx.Request.Method {
	case http.MethodGet:
		responses.SendJsonResponse(ctx, http.StatusOK, "services are running", &gin.H{
			"time": time.Now().Unix(),
		})

	default:
		responses.SendMethodNotAllowed(ctx)
	}
}
