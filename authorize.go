package echo_middleware

import "github.com/labstack/echo/v4"

type AuthorizeFunc func(c echo.Context) bool

func Authorize(authorize AuthorizeFunc) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			if !authorize(c) {
				return echo.ErrUnauthorized
			}
			return next(c)
		}
	}
}
