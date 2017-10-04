package ginzap

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// NowFunc returns current time.
var NowFunc = func() time.Time {
	return time.Now().UTC()
}

// New creates new gin middleware object for logging request/response with zap.Logger.
func New(logger *zap.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		requestedAt := NowFunc()
		path := c.Request.URL.Path

		c.Next()

		latency := time.Since(requestedAt)

		if len(c.Errors) > 0 {
			for _, errMsg := range c.Errors.Errors() {
				logger.Error(errMsg)
			}
		} else {
			logger.Info(
				fmt.Sprintf("%s %s", c.Request.Method, path),
				zap.Int("status", c.Writer.Status()),
				zap.String("method", c.Request.Method),
				zap.String("path", path),
				zap.String("ip", c.ClientIP()),
				zap.String("user-agent", c.Request.UserAgent()),
				zap.Time("time", requestedAt),
				zap.Duration("latency", latency),
				zap.Object("headers", zapcore.ObjectMarshalerFunc(func(inner zapcore.ObjectEncoder) error {
					for k, values := range c.Request.Header {
						err := inner.AddArray(k, zapcore.ArrayMarshalerFunc(func(inner zapcore.ArrayEncoder) error {
							for _, v := range values {
								inner.AppendString(v)
							}
							return nil
						}))
						if err != nil {
							return err
						}
					}
					return nil
				})),
			)
		}
	}
}
