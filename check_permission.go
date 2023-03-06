package echo_middleware

import "github.com/labstack/echo/v4"

func CheckPermission(check func(c echo.Context) bool) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			if !check(c) {
				return echo.ErrForbidden
			}
			return next(c)
		}
	}
}
