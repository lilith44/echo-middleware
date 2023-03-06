package echo_middleware

import "github.com/labstack/echo/v4"

type CheckPermissionFunc func(c echo.Context) bool

func CheckPermission(check CheckPermissionFunc) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			if !check(c) {
				return echo.ErrForbidden
			}
			return next(c)
		}
	}
}
