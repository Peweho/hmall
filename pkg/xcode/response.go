package xcode

import (
	"hmall/pkg/xcode/types"
)

func ErrHandler(err error) (int, any) {
	code := CodeFromError(err)

	return code.Code(), types.Status{
		Message: code.Message(),
	}
}
