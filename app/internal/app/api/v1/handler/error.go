package handler

import (
	"net/http"

	"github.com/Khmer495/go-templete/internal/app/api/v1/openapi"
	"github.com/Khmer495/go-templete/internal/pkg/util/cerror"
	"github.com/labstack/echo/v4"
	"golang.org/x/xerrors"
)

func NewAlreadyExistError(ctx echo.Context, msg string) error {
	return ctx.JSON(
		http.StatusBadRequest,
		openapi.Error{
			Code:    cerror.AlreadyExistsErrorCode.ToString(),
			Message: msg,
		},
	)
}

func NewInvalidArgumentError(ctx echo.Context, msg string) error {
	return ctx.JSON(
		http.StatusBadRequest,
		openapi.Error{
			Code:    "VALIDATION_ERROR",
			Message: "パラメータが不正です。" + msg,
		},
	)
}

func NewUnAuthenticatedError(ctx echo.Context, msg string) error {
	return ctx.JSON(
		http.StatusUnauthorized,
		openapi.Error{
			Code:    "NOT_REGISTERED_USER",
			Message: "登録されていないユーザーです。",
		},
	)
}

func NewInternalServerError(ctx echo.Context) error {
	return ctx.JSON(
		http.StatusInternalServerError,
		openapi.Error{
			Code:    "UNKNOWN_ERROR",
			Message: "不明なエラーが発生しました。",
		},
	)
}

func NewNotRegisteredUserError(ctx echo.Context) error {
	return ctx.JSON(
		http.StatusUnauthorized,
		openapi.Error{
			Code:    "NOT_REGISTERED_USER",
			Message: "登録されていないユーザーです。",
		},
	)
}

func NewInvalidTokenError(ctx echo.Context) error {
	return ctx.JSON(
		http.StatusUnauthorized,
		openapi.Error{
			Code:    "INVALID_TOKEN",
			Message: "不正なトークンです。",
		},
	)
}

func NewInvalidRequestFormatError(ctx echo.Context, err error) error {
	eerr := &echo.HTTPError{}
	msg := "不正なリクエストの形式です。"
	if xerrors.As(err, &eerr) {
		msg += eerr.Internal.Error()
	}
	return ctx.JSON(
		http.StatusBadRequest,
		openapi.Error{
			Code:    "INVALID_REQUEST_FORMAT",
			Message: msg,
		},
	)
}

func NewAlreadyRegisteredUserError(ctx echo.Context) error {
	return ctx.JSON(
		http.StatusBadRequest,
		openapi.Error{
			Code:    "ALREADY_REGISTERED_USER_ERROR",
			Message: "既に登録されているユーザーです。",
		},
	)
}

func NewNotFoundError(ctx echo.Context) error {
	return ctx.JSON(
		http.StatusNotFound,
		openapi.Error{
			Code:    "NOT_FOUND",
			Message: "指定のリソースが存在しません。",
		},
	)
}

func NewAlreadyAppliedError(ctx echo.Context) error {
	return ctx.JSON(
		http.StatusBadRequest,
		openapi.Error{
			Code:    "ALREADY_APPLIED",
			Message: "既に参加申請済みのチームです。",
		},
	)
}

func NewAlreadyOperatedApplicantError(ctx echo.Context) error {
	return ctx.JSON(
		http.StatusBadRequest,
		openapi.Error{
			Code:    "ALREADY_OPERATED_APPLICANT",
			Message: "既に参加申請が承認、もしくは拒否されたユーザーです。",
		},
	)
}
