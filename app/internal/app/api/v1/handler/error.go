package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"golang.org/x/xerrors"
)

func NewInvalidArgumentError(ctx echo.Context, msg string) error {
	return ctx.JSON(http.StatusBadRequest, msg)
}

func NewInternalServerError(ctx echo.Context) error {
	return ctx.JSON(
		http.StatusInternalServerError,
		"不明なエラーが発生しました。",
	)
}

func NewNotRegisteredUserError(ctx echo.Context) error {
	return ctx.JSON(
		http.StatusUnauthorized,
		"登録されていないユーザーです。",
	)
}

func NewInvalidTokenError(ctx echo.Context) error {
	return ctx.JSON(
		http.StatusUnauthorized,
		"不正なトークンです。",
	)
}

func NewInvalidRequestFormatError(ctx echo.Context, err error) error {
	eerr := &echo.HTTPError{}
	msg := "不正なリクエストの形式です。"
	if xerrors.As(err, &eerr) {
		msg += eerr.Internal.Error()
	}
	return ctx.JSON(http.StatusBadRequest, msg)
}

func NewAlreadyRegisteredUserError(ctx echo.Context) error {
	return ctx.JSON(
		http.StatusBadRequest,
		"既に登録されているユーザーです。",
	)
}

func NewNotFoundError(ctx echo.Context) error {
	return ctx.JSON(
		http.StatusNotFound,
		"指定のリソースが存在しません。",
	)
}

func NewAlreadyJoinedError(ctx echo.Context) error {
	return ctx.JSON(
		http.StatusBadRequest,
		"既に参加済みのチームです。",
	)
}
