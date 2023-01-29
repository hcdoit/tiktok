package pack

import (
	"errors"
	"github.com/hcdoit/tiktok/pkg/errno"
)

func BuildStatus(err error) (int32, string) {
	if err == nil {
		return errno.Success.ErrCode, errno.Success.ErrMsg
	}
	e := errno.ErrNo{}
	if errors.As(err, &e) {
		return e.ErrCode, e.ErrMsg
	}
	s := errno.ServiceErr.WithMessage(err.Error())
	return s.ErrCode, s.ErrMsg
}
