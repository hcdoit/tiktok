package main

import (
	"context"
	"github.com/hcdoit/tiktok/cmd/user/service"
	"github.com/hcdoit/tiktok/cmd/user/utils"
	"github.com/hcdoit/tiktok/kitex_gen/user"
	"github.com/hcdoit/tiktok/pkg/errno"
	"github.com/hcdoit/tiktok/pkg/jwt"
)

// UserServiceImpl implements the last service interface defined in the IDL.
type UserServiceImpl struct{}

// Register implements the UserServiceImpl interface.
func (s *UserServiceImpl) Register(ctx context.Context, req *user.UserRegisterRequest) (resp *user.UserRegisterResponse, err error) {
	resp = new(user.UserRegisterResponse)

	// 校验参数
	if err = req.IsValid(); err != nil {
		resp.StatusCode, resp.StatusMsg = utils.BuildStatus(err)
		return resp, nil
	}

	// 调用service层
	err = service.NewRegisterService(ctx).Register(req)
	if err != nil {
		resp.StatusCode, resp.StatusMsg = utils.BuildStatus(err)
		return resp, nil
	}

	// 注册后登录获取Token
	uid, token, err := service.NewLoginService(ctx).Login(&user.UserLoginRequest{Username: req.Username, Password: req.Password})
	if err != nil {
		resp.StatusCode, resp.StatusMsg = utils.BuildStatus(err)
		return resp, nil
	}

	//包装正常响应
	resp.UserId = uid
	resp.Token = token
	resp.StatusCode, resp.StatusMsg = utils.BuildStatus(errno.Success)

	return resp, nil
}

// Login implements the UserServiceImpl interface.
func (s *UserServiceImpl) Login(ctx context.Context, req *user.UserLoginRequest) (resp *user.UserLoginResponse, err error) {
	resp = new(user.UserLoginResponse)

	// 校验参数
	if err = req.IsValid(); err != nil {
		resp.StatusCode, resp.StatusMsg = utils.BuildStatus(err)
		return resp, nil
	}

	// 调用service层
	uid, token, err := service.NewLoginService(ctx).Login(&user.UserLoginRequest{Username: req.Username, Password: req.Password})
	if err != nil {
		resp.StatusCode, resp.StatusMsg = utils.BuildStatus(err)
		return resp, nil
	}

	// 包装正常响应
	resp.UserId = uid
	resp.Token = token
	resp.StatusCode, resp.StatusMsg = utils.BuildStatus(errno.Success)
	return resp, nil
}

// GetUser implements the UserServiceImpl interface.
func (s *UserServiceImpl) GetUser(ctx context.Context, req *user.GetUserRequest) (resp *user.GetUserResponse, err error) {
	resp = new(user.GetUserResponse)

	// 校验参数
	if err = req.IsValid(); err != nil {
		resp.StatusCode, resp.StatusMsg = utils.BuildStatus(err)
		return resp, nil
	}

	// 解析Token获取id，若空默认为0
	myID := int64(0)
	if len(req.Token) != 0 {
		claim, err := jwt.ParseToken(req.Token)
		if err != nil {
			resp.StatusCode, resp.StatusMsg = utils.BuildStatus(err)
			return resp, nil
		}
		myID = claim.Id
	}

	// 调用service层
	user, err := service.NewGetUserService(ctx).GetUser(req, myID)
	if err != nil {
		resp.StatusCode, resp.StatusMsg = utils.BuildStatus(err)
		return resp, nil
	}

	// 包装正常响应
	resp.StatusCode, resp.StatusMsg = utils.BuildStatus(errno.Success)
	resp.User = user

	return resp, nil
}
