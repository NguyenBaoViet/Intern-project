package utils

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func Middleware_logger(param gin.LogFormatterParams) string {

	// your custom format
	return fmt.Sprintf("ClientIP: %s - Time: [%s], Method=%s, Path=%s, Protocol=%s, Statuscode=%d, Latency=%s, User-agent=%s, err=%s, x-request-id=%s\n",
		param.ClientIP,
		param.TimeStamp.Format(time.RFC1123),
		param.Method,
		param.Path,
		param.Request.Proto,
		param.StatusCode,
		param.Latency,
		param.Request.UserAgent(),
		param.ErrorMessage,
		param.Request.Header.Get("x-request-id"),
	)
}
