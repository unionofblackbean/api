package responses

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSendErrorResponse(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	SendErrorResponse(c, http.StatusInternalServerError, errors.New("test error"))

	assert.Equal(t, http.StatusInternalServerError, w.Code)
	assert.Equal(t, `{"code":500,"msg":"test error","data":{}}`, w.Body.String())
}
