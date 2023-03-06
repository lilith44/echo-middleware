package echo_middleware

import (
	"fmt"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func Log(logger *zap.Logger) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			start := time.Now()

			err := next(c)
			if err != nil {
				c.Error(err)
			}

			req := c.Request()
			rsp := c.Response()

			fields := []zapcore.Field{
				zap.Int("status", rsp.Status),
				zap.String("request", fmt.Sprintf("%s %s", req.Method, req.RequestURI)),
				zap.Error(err),
				zap.Int64("size", rsp.Size),
				zap.String("latency", time.Since(start).String()),
				zap.String("remote_ip", c.RealIP()),
				zap.String("user_agent", req.UserAgent()),
			}

			requestId := req.Header.Get(echo.HeaderXRequestID)
			if requestId == "" {
				requestId = rsp.Header().Get(echo.HeaderXRequestID)
			}
			fields = append(fields, zap.String("request_id", requestId))

			if rsp.Status >= http.StatusMultipleChoices {
				logger.Error("", fields...)
			} else {
				logger.Info("", fields...)
			}

			return nil
		}
	}
}
