package main

import (
	"context"
	"github.com/hcdoit/tiktok/cmd/user/service"
	"github.com/hcdoit/tiktok/cmd/user/utils"
	"github.com/hcdoit/tiktok/kitex_gen/user"
	"github.com/hcdoit/tiktok/pkg/errno"
)

// UserServiceImpl implements the last service interface defined in the IDL.
type UserServiceImpl struct{}

// Register implements the UserServiceImpl interface.
func (s *UserServiceImpl) Register(ctx context.Context, req *user.UserRegisterRequest) (resp *user.UserRegisterResponse, err error) {
	// TODO: Your code here...
	resp = new(user.UserRegisterResponse)
	//判断参数长度
	if err = req.IsValid(); err != nil {
		resp.StatusCode, resp.StatusMsg = utils.BuildStatus(err)
		return resp, nil
	}
	//注册用户
	err = service.NewRegisterService(ctx).Register(req)
	if err != nil {
		resp.StatusCode, resp.StatusMsg = utils.BuildStatus(err)
	}
	//注册后登录获得uid
	uid, token, err := service.NewLoginService(ctx).Login(&user.UserLoginRequest{Username: req.Username, Password: req.Password})
	if err != nil {
		resp.StatusCode, resp.StatusMsg = utils.BuildStatus(err)
		return resp, nil
	}
	resp.UserId = uid
	resp.Token = token
	resp.StatusCode, resp.StatusMsg = utils.BuildStatus(errno.Success)

	return resp, nil
}

// Login implements the UserServiceImpl interface.
func (s *UserServiceImpl) Login(ctx context.Context, req *user.UserLoginRequest) (resp *user.UserLoginResponse, err error) {
	// TODO: Your code here...
	resp = new(user.UserLoginResponse)
	//判断参数长度
	if err = req.IsValid(); err != nil {
		resp.StatusCode, resp.StatusMsg = utils.BuildStatus(err)
		return resp, nil
	}
	//登录
	uid, token, err := service.NewLoginService(ctx).Login(&user.UserLoginRequest{Username: req.Username, Password: req.Password})
	if err != nil {
		resp.StatusCode, resp.StatusMsg = utils.BuildStatus(err)
		return resp, nil
	}
	resp.UserId = uid
	resp.Token = token
	resp.StatusCode, resp.StatusMsg = utils.BuildStatus(errno.Success)
	return resp, nil
}

// GetUser implements the UserServiceImpl interface.
func (s *UserServiceImpl) GetUser(ctx context.Context, req *user.GetUserRequest) (resp *user.GetUserResponse, err error) {
	// TODO: Your code here...
	resp = new(user.GetUserResponse)
	claim, err := utils.Jwt.ParseToken(req.Token)
	if err != nil {
		resp.StatusCode, resp.StatusMsg = utils.BuildStatus(err)
		return resp, nil
	}
	if err = req.IsValid(); err != nil {
		resp.StatusCode, resp.StatusMsg = utils.BuildStatus(err)
		return resp, nil
	}

	user, err := service.NewGetUserService(ctx).GetUser(req, claim.Id)
	if err = req.IsValid(); err != nil {
		resp.StatusCode, resp.StatusMsg = utils.BuildStatus(err)
		return resp, nil
	}
	resp.StatusCode, resp.StatusMsg = utils.BuildStatus(errno.Success)
	resp.User = user

	return resp, nil
}
