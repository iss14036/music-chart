package responsewrapper

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func Response(message, status, data interface{}) map[string]interface{} {
	return map[string]interface{}{
		"message": message,
		"status":  status,
		"data":    data,
	}
}

func response(message string, httpStatus int, data interface{}) map[string]interface{} {
	return Response(message, http.StatusText(httpStatus), data)
}

func OK(c echo.Context, message string, data interface{}) error {
	body := response(message, http.StatusOK, data)
	return c.JSON(http.StatusOK, body)
}

func Created(c echo.Context, message string, data interface{}) error {
	body := response(message, http.StatusCreated, data)
	return c.JSON(http.StatusCreated, body)
}

func BadRequest(c echo.Context, err error, data interface{}) error {
	body := response(err.Error(), http.StatusBadRequest, data)
	return c.JSON(http.StatusBadRequest, body)
}

func InternalServerError(c echo.Context, err error, data interface{}) error {
	body := response(err.Error(), http.StatusInternalServerError, data)
	return c.JSON(http.StatusInternalServerError, body)
}

func Unauthorized(c echo.Context, err error, data interface{}) error {
	body := response(err.Error(), http.StatusUnauthorized, data)
	return c.JSON(http.StatusUnauthorized, body)
}
