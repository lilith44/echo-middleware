package echo_middleware

import "github.com/labstack/echo/v4"

func Authorize(authorize func(c echo.Context) bool) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			if !authorize(c) {
				return echo.ErrUnauthorized
			}
			return next(c)
		}
	}
}
