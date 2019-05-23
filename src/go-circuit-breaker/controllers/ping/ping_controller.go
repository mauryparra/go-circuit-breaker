package ping

import (
	"github.com/gin-gonic/gin"
)

// Ping answers with string "pong"
func Ping(context *gin.Context) {
	context.String(200, "Pong Middle")
}
