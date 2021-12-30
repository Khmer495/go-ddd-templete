package cerror

import (
	"fmt"

	"golang.org/x/xerrors"
)

type customError struct {
	level     customErrorLevel
	code      customErrorCode
	msg       string
	clientMsg string
	err       error
	frame     xerrors.Frame
}

var nilCustomError = &customError{}

func (e customError) Error() string {
	return fmt.Sprintf("*****StartCustomError*****\nErrorLevel: %s\nErrorCode: %s\nErrorMsg: %s\nClientMsg: %s\n*****EndCustomError*****", e.level, e.code, e.msg, e.clientMsg)
}

func (e customError) Level() customErrorLevel {
	return e.level
}

func (e customError) ClientMsg() string {
	return e.clientMsg
}

func New(level customErrorLevel, code customErrorCode, msg string, clientMsg string) *customError {
	return Wrap(nil, level, code, msg, clientMsg)
}

func Wrap(err error, level customErrorLevel, code customErrorCode, msg string, clientMsg string) *customError {
	return &customError{
		level:     level,
		code:      code,
		msg:       msg,
		clientMsg: clientMsg,
		err:       err,
		frame:     xerrors.Caller(1),
	}
}

func (e *customError) Unwrap() error {
	return e.err
}

func (e *customError) Is(err error) bool {
	var cerr *customError
	return xerrors.As(err, cerr) && e.code == cerr.code
}

func (e *customError) IsLevel(level customErrorLevel) bool {
	return e.level == level
}

func (e *customError) IsCode(code customErrorCode) bool {
	return e.code == code
}

func (e *customError) FormatError(p xerrors.Printer) (next error) {
	p.Print(e.Error())
	e.frame.Format(p)
	return e.err
}

func (e *customError) Format(s fmt.State, v rune) {
	xerrors.FormatError(e, s, v)
}

func IsLevel(err error, level customErrorLevel) bool {
	cerr := nilCustomError
	return xerrors.As(err, &cerr) && cerr.IsLevel(level)
}

func IsCode(err error, code customErrorCode) bool {
	if err == nil {
		return false
	}
	cerr := nilCustomError
	return xerrors.As(err, &cerr) && cerr.IsCode(code)
}

func As(err error) *customError {
	cerr := nilCustomError
	xerrors.As(err, &cerr)
	return cerr
}

var (
	NotImplementedError *customError = New(ErrorLevel, NotImplementedErrorCode, NotImplementedErrorCode.ToString(), NotImplementedErrorCode.ToString())
)

func NewInternalServerError(msg string) *customError {
	return Wrap(nil, ErrorLevel, InterServerErrorCode, msg, InterServerErrorCode.ToString())
}

func WrapInternalServerError(err error, msg string) *customError {
	return Wrap(err, ErrorLevel, InterServerErrorCode, msg, InterServerErrorCode.ToString())
}

func NewInvalidArgumentError(msg string, clientMsg string) *customError {
	return Wrap(nil, ErrorLevel, InvalidArgumentErrorCode, msg, clientMsg)
}

func WrapInvalidArgumentError(err error, msg string, clientMsg string) *customError {
	return Wrap(err, ErrorLevel, InvalidArgumentErrorCode, msg, clientMsg)
}

func NewNotFoundError(msg string, clientMsg string) *customError {
	return Wrap(nil, ErrorLevel, NotFoundErrorCode, msg, clientMsg)
}

func WrapNotFoundError(err error, msg string, clientMsg string) *customError {
	return Wrap(err, ErrorLevel, NotFoundErrorCode, msg, clientMsg)
}

func NewUnauthorizedError(msg string, clientMsg string) *customError {
	return Wrap(nil, ErrorLevel, UnauthorizedErrorCode, msg, clientMsg)
}

func WrapUnauthorizedError(err error, msg string, clientMsg string) *customError {
	return Wrap(err, ErrorLevel, UnauthorizedErrorCode, msg, clientMsg)
}

func NewAlreadyExistError(msg string, clientMsg string) *customError {
	return Wrap(nil, ErrorLevel, AlreadyExistsErrorCode, msg, clientMsg)
}
