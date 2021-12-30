package entmysql

import (
	"github.com/Khmer495/go-templete/internal/pkg/infrastracture/ent"
	"github.com/Khmer495/go-templete/internal/pkg/util/cerror"
	"golang.org/x/xerrors"
)

func newIsNotFoundError(err error, msg string, clientMsg string) error {
	notFoundErr := &ent.NotFoundError{}
	if xerrors.As(err, &notFoundErr) {
		return cerror.Wrap(notFoundErr, cerror.ErrorLevel, cerror.NotFoundErrorCode, msg, clientMsg)
	}
	return nil
}

func newIsNotSingularError(err error, msg string, clientMsg string) error {
	notSingularErr := &ent.NotSingularError{}
	if xerrors.As(err, &notSingularErr) {
		return cerror.Wrap(notSingularErr, cerror.ErrorLevel, cerror.InterServerErrorCode, msg, clientMsg)
	}
	return nil
}

func newTxRollbackError(err error, msg string) error {
	return cerror.Wrap(err, cerror.ErrorLevel, cerror.InterServerErrorCode, msg+"/tx.Rollback", cerror.InterServerErrorCode.ToString())
}

func newTxCommitError(err error) error {
	return cerror.WrapInternalServerError(err, "tx.Commit")
}

// func parseFindError(err error, msg string) error {
// 	if _err := newIsNotFoundError(err); _err != nil {
// 		return _err
// 	}
// 	return cerror.WrapInternalServerError(err, msg)
// }

func parseQueryOnlyByUlid(err error, msg string, clientMsg string) error {
	if _err := newIsNotFoundError(err, msg, clientMsg); _err != nil {
		return _err
	}
	if _err := newIsNotSingularError(err, "同一のUlidが複数存在します。Ulidフィールドにユニーク制約が付与されていることを確認してください。", cerror.InterServerErrorCode.ToString()); _err != nil {
		return _err
	}
	return cerror.WrapInternalServerError(err, msg)
}

func txRollbackAndParseError(tx *ent.Tx, err error, msg string) error {
	rerr := tx.Rollback()
	if rerr != nil {
		return newTxRollbackError(err, msg)
	}
	return xerrors.Errorf(msg+": %w", err)
}
