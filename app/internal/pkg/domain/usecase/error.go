package usecase

import (
	"github.com/Khmer495/go-templete/internal/pkg/util/cerror"
)

func parsemodelConstractorError(err error, msg string) error {
	if cerror.IsCode(err, cerror.InvalidArgumentErrorCode) {
		return cerror.WrapInvalidArgumentError(err, msg, cerror.As(err).ClientMsg())
	}
	return cerror.WrapInternalServerError(err, msg)
}
