package responses

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSendTooManyRequestsResponse(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	SendTooManyRequestsResponse(c)

	assert.Equal(t, http.StatusTooManyRequests, w.Code)
	assert.Equal(t, `{"code":429,"msg":"too many requests","data":{}}`, w.Body.String())
}
