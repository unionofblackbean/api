package responses

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSendMethodNotAllowed(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	SendMethodNotAllowed(c)

	assert.Equal(t, http.StatusMethodNotAllowed, w.Code)
	assert.Equal(t, `{"code":405,"msg":"method not allowed","data":{}}`, w.Body.String())
}
