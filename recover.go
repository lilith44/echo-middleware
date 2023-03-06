package echo_middleware

import (
	"fmt"
	"runtime/debug"

	"github.com/labstack/echo/v4"
)

func Recover() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) (err error) {
			defer func() {
				if r := recover(); r != nil {
					c.Logger().Info(fmt.Sprintf("[Panic Recover]\n%s", debug.Stack()))

					err = fmt.Errorf("%s", r)
				}
			}()

			return next(c)
		}
	}
}
