package helper

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func JSONErrorResponse(ctx echo.Context, status int, message string) error {
	return ctx.JSON(status, map[string]interface{}{
		"success": false,
		"message": message,
	})
}

func JSONSuccessResponse(ctx echo.Context, data interface{}) error {
	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"success": true,
		"data":    data,
	})
}

func GetIDParam(ctx echo.Context) (int, error) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return 0, err
	}
	return id, nil
}
