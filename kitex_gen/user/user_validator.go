package user

import (
	"bytes"
	"fmt"
	"reflect"
	"regexp"
	"strings"
	"time"
)

// unused protection
var (
	_ = fmt.Formatter(nil)
	_ = (*bytes.Buffer)(nil)
	_ = (*strings.Builder)(nil)
	_ = reflect.Type(nil)
	_ = (*regexp.Regexp)(nil)
	_ = time.Nanosecond
)

func (p *UserRegisterRequest) IsValid() error {
	if len(p.Username) < int(1) || len(p.Username) > int(32) {
		return fmt.Errorf("field Username len rule failed, current value: %d", len(p.Username))
	}
	if len(p.Password) < int(1) || len(p.Password) > int(32) {
		return fmt.Errorf("field Password len rule failed, current value: %d", len(p.Password))
	}
	return nil
}

func (p *UserLoginRequest) IsValid() error {
	if len(p.Username) < int(1) || len(p.Username) > int(32) {
		return fmt.Errorf("field Username len rule failed, current value: %d", len(p.Username))
	}
	if len(p.Password) < int(1) || len(p.Password) > int(32) {
		return fmt.Errorf("field Password len rule failed, current value: %d", len(p.Password))
	}
	return nil
}

func (p *GetUserRequest) IsValid() error {
	if p.UserId < 0 {
		return fmt.Errorf("field UserId rule failed, current value: %d", p.UserId)
	}
	return nil
}
