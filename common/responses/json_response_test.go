package responses

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSendJsonResponse(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	SendJsonResponse(c, http.StatusOK, "test message", map[string]interface{}{
		"test_key": "test_value",
	})

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, `{"code":200,"msg":"test message","data":{"test_key":"test_value"}}`, w.Body.String())
}
