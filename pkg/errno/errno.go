package errno

import (
	"errors"
	"fmt"
	"github.com/hcdoit/tiktok/kitex_gen/user"
)

type ErrNo struct {
	ErrCode int32
	ErrMsg  string
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
	Success                = NewErrNo(int32(user.ErrCode_SuccessCode), "Success")
	ServiceErr             = NewErrNo(int32(user.ErrCode_ServiceErrCode), "Service is unable to start successfully")
	ParamErr               = NewErrNo(int32(user.ErrCode_ParamErrCode), "Wrong Parameter has been given")
	UserAlreadyExistErr    = NewErrNo(int32(user.ErrCode_UserAlreadyExistErrCode), "User already exists")
	AuthorizationFailedErr = NewErrNo(int32(user.ErrCode_AuthorizationFailedErrCode), "Authorization failed")
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
