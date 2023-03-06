package echo_middleware

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/lilith44/easy"
)

type HttpErrorHandlerOptions struct {
	MiddlewareErrorCode int
	DefaultErrorCode    int
	DefaultErrorMessage string
}

func HttpErrorHandler(options *HttpErrorHandlerOptions) func(err error, c echo.Context) {
	return func(err error, c echo.Context) {
		if c.Response().Committed {
			return
		}

		switch e := err.(type) {
		case *echo.HTTPError:
			_ = c.JSON(e.Code, easy.Fail(options.MiddlewareErrorCode, e.Message))
		case *echo.BindingError:
			_ = c.JSON(e.Code, easy.Fail(options.MiddlewareErrorCode, e.Message))
		case *validator.ValidationErrors, validator.ValidationErrors:
			_ = c.JSON(http.StatusBadRequest, easy.Fail(options.MiddlewareErrorCode, e.Error()))
		case easy.HttpError:
			_ = c.JSON(e.StatusCode(), easy.Fail(e.Code(), e.Error()))
		default:
			_ = c.JSON(http.StatusInternalServerError, easy.Fail(options.DefaultErrorCode, options.DefaultErrorMessage))
		}
	}
}
