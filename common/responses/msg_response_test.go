package responses

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSendMsgResponse(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	SendMsgResponse(c, http.StatusOK, "test message")

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, `{"code":200,"msg":"test message","data":{}}`, w.Body.String())
}
