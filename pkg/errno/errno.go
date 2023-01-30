package errno

import (
	"errors"
	"fmt"
)

type ErrNo struct {
	ErrCode int32  `json:"status_code"`
	ErrMsg  string `json:"status_msg"`
}

func (e ErrNo) Error() string {
	return fmt.Sprintf("err_code=%d, err_msg=%s", e.ErrCode, e.ErrMsg)
}

func NewErrNo(code int32, msg string) ErrNo {
	return ErrNo{
		ErrCode: code,
		ErrMsg:  msg,
	}
}

func (e ErrNo) WithMessage(msg string) ErrNo {
	e.ErrMsg = msg
	return e
}

var (
	Success                = NewErrNo(int32(0), "Success")
	ServiceErr             = NewErrNo(int32(1), "Service is unable to start successfully")
	ParamErr               = NewErrNo(int32(2), "Wrong Parameter has been given")
	UserAlreadyExistErr    = NewErrNo(int32(3), "User already exists")
	AuthorizationFailedErr = NewErrNo(int32(4), "Authorization failed")
	ResourceNotFound       = NewErrNo(int32(5), "Resource not found")
)

func ConvertErr(err error) ErrNo {
	Err := ErrNo{}
	if errors.As(err, &Err) {
		return Err
	}
	s := ServiceErr
	s.ErrMsg = err.Error()
	return s
}
