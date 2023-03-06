package echo_middleware

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

func Log() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			start := time.Now()

			err := next(c)
			if err != nil {
				c.Error(err)
			}

			req := c.Request()
			rsp := c.Response()

			format := `{"request_id": "%s", "request": "%s %s", "status": %d, "error": "%s", "size": %d, "latency": "%s", "remote_ip": "%s", "user_agent": "%s"}`

			requestId := req.Header.Get(echo.HeaderXRequestID)
			if requestId == "" {
				requestId = rsp.Header().Get(echo.HeaderXRequestID)
			}

			var e string
			if err != nil {
				e = err.Error()
			}

			if rsp.Status >= http.StatusMultipleChoices {
				c.Logger().Errorf(
					format,
					requestId,
					req.Method, req.RequestURI,
					rsp.Status,
					e,
					rsp.Size,
					time.Since(start).String(),
					c.RealIP(),
					req.UserAgent(),
				)
			} else {
				c.Logger().Infof(
					format,
					requestId,
					req.Method, req.RequestURI,
					rsp.Status,
					e,
					rsp.Size,
					time.Since(start).String(),
					c.RealIP(),
					req.UserAgent(),
				)
			}

			return nil
		}
	}
}
