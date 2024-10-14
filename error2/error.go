package error2

import "bitbucket.org/alibaba-international/go-pkg/errorx"

var (
	ErrRecordCanceled       = errorx.C(1000, "record already canceled")
	ErrRecordNotCompleted   = errorx.C(1001, "record not completed")
	ErrTransactionNotExists = errorx.C(1002, "transaction does not exist")
	ErrEnvNotFound          = errorx.C(1003, "environment not found")
	ErrInternal             = errorx.C(errorx.CodeInternal, "internal server error")
)
