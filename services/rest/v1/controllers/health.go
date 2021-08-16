package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/unionofblackbean/api/common"
	"github.com/unionofblackbean/api/common/responses"
	"net/http"
)

func Health(ctx *gin.Context) {
	switch ctx.Request.Method {
	case http.MethodGet:
		responses.SendJsonResponse(ctx, http.StatusOK, "services are running", &gin.H{
			"time": common.NowUTCUnix(),
		})

	default:
		responses.SendMethodNotAllowed(ctx)
	}
}
