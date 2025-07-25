package service

import (
	"context"

	userpb "kratos/api/user/v1"
	"kratos/internal/biz"
	"kratos/internal/pkg"
)

type UserService struct {
	userpb.UnimplementedUserServiceServer
	repo      biz.UserRepo
	dtmServer string // DTM服务地址1
}

func NewUserService(repo biz.UserRepo, dtmServer string) *UserService {
	return &UserService{repo: repo, dtmServer: dtmServer}
}

// 用户注册（Saga事务分支）
func (s *UserService) Register(ctx context.Context, req *userpb.RegisterRequest) (*userpb.RegisterReply, error) {
	id, err := s.repo.Register(&biz.User{
		UserName:     req.UserName,
		UserNickname: req.UserNickname,
		Department:   req.Department,
		Mobile:       req.Mobile,
		Email:        req.Email,
		Password:     req.Password,
		Gender:       req.Gender,
		Role:         req.Role,
		Comment:      req.Comment,
	}, "")
	if err != nil {
		return nil, err
	}
	return &userpb.RegisterReply{Id: id}, nil
}

// 用户注册补偿（Saga事务分支补偿）
func (s *UserService) RegisterCompensate(ctx context.Context, req *userpb.RegisterCompensateRequest) (*userpb.RegisterCompensateReply, error) {
	// 直接删除用户
	err := s.repo.RegisterCompensate(req.Id, req.DtmGid)
	if err != nil {
		return &userpb.RegisterCompensateReply{Success: false}, err
	}
	return &userpb.RegisterCompensateReply{Success: true}, nil
}

// 用户登录
func (s *UserService) Login(ctx context.Context, req *userpb.LoginRequest) (*userpb.LoginReply, error) {
	user, err := s.repo.Login(req.UserName, req.Password)
	if err != nil {
		return nil, err
	}
	token, err := pkg.GenerateToken(user.Id, user.UserName)
	if err != nil {
		return nil, err
	}
	return &userpb.LoginReply{
		Id:       user.Id,
		UserName: user.UserName,
		Token:    token,
	}, nil
}

// 查询单个用户
func (s *UserService) GetUser(ctx context.Context, req *userpb.GetUserRequest) (*userpb.GetUserReply, error) {
	user, err := s.repo.GetUser(req.Id)
	if err != nil {
		return nil, err
	}
	return &userpb.GetUserReply{User: toProtoUser(user)}, nil
}

// 用户列表
func (s *UserService) ListUser(ctx context.Context, req *userpb.ListUserRequest) (*userpb.ListUserReply, error) {
	users, total, err := s.repo.ListUser(req.Page, req.PageSize)
	if err != nil {
		return nil, err
	}
	var pbUsers []*userpb.User
	for _, u := range users {
		pbUsers = append(pbUsers, toProtoUser(u))
	}
	return &userpb.ListUserReply{Users: pbUsers, Total: total}, nil
}

// 用户更新
func (s *UserService) UpdateUser(ctx context.Context, req *userpb.UpdateUserRequest) (*userpb.UpdateUserReply, error) {
	user := &biz.User{
		Id:           req.Id,
		UserNickname: req.UserNickname,
		Department:   req.Department,
		Mobile:       req.Mobile,
		Email:        req.Email,
		Gender:       req.Gender,
		Role:         req.Role,
		UserStatus:   req.UserStatus,
		Comment:      req.Comment,
	}
	err := s.repo.UpdateUser(user)
	return &userpb.UpdateUserReply{Success: err == nil}, err
}

// 用户删除
func (s *UserService) DeleteUser(ctx context.Context, req *userpb.DeleteUserRequest) (*userpb.DeleteUserReply, error) {
	err := s.repo.DeleteUser(req.Id)
	return &userpb.DeleteUserReply{Success: err == nil}, err
}

// 工具函数
func toProtoUser(u *biz.User) *userpb.User {
	if u == nil {
		return nil
	}
	return &userpb.User{
		Id:           u.Id,
		UserName:     u.UserName,
		UserNickname: u.UserNickname,
		Department:   u.Department,
		Mobile:       u.Mobile,
		Email:        u.Email,
		Password:     "", // 不返回密码
		Gender:       u.Gender,
		Role:         u.Role,
		UserStatus:   u.UserStatus,
		Comment:      u.Comment,
	}
}
