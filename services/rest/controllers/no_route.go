package controllers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/leungyauming/api/common/responses"
	"net/http"
)

func NoRoute(ctx *gin.Context) {
	responses.SendErrorResponse(ctx,
		http.StatusNotFound,
		errors.New("unknown endpoint"))
}
