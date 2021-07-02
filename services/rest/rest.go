package rest

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func Start(addr string, port uint16) error {
	r := gin.Default()

	return r.Run(fmt.Sprintf("%s:%d", addr, port))
}
