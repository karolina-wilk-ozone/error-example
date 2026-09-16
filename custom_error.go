package main

import (
	"errors"
	"log"
	"net/http"

	"github.com/labstack/echo/v5"
)

type AppError struct {
	Code     int
	Public   string
	Internal error
}

type ErrorResponse struct {
	Error string `json:"error"`
}

func (e *AppError) Error() string   { return e.Public }
func (e *AppError) Unwrap() error   { return e.Internal }
func (e *AppError) StatusCode() int { return e.Code }

func BadRequest(public string, internal error) *AppError {
	return &AppError{
		Code:     http.StatusBadRequest,
		Public:   public,
		Internal: internal,
	}
}

func Internal(internal error) *AppError {
	return &AppError{http.StatusInternalServerError, "internal server error", internal}
}

func errorHandler(c *echo.Context, err error) {
	var appErr *AppError
	if !errors.As(err, &appErr) {
		appErr = Internal(err)
	}
	if appErr.Internal != nil {
		reqID := c.Request().Header.Get(echo.HeaderXRequestID)
		log.Printf("request_id=%s path=%s: %v", reqID, c.Path(), appErr.Internal)
	}
	if r, _ := echo.UnwrapResponse(c.Response()); r != nil && r.Committed {
		return
	}
	if cErr := c.JSON(appErr.Code, ErrorResponse{Error: appErr.Public}); cErr != nil {
		log.Printf("failed writing error response: %v", cErr)
	}
}
