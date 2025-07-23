package service

import (
	v1 "IOT_Hummingbird_back_end/api/user/v1"
	"IOT_Hummingbird_back_end/internal/biz"
	"IOT_Hummingbird_back_end/internal/data"
	context "context"
)

type UserService struct {
	v1.UnimplementedUserServer
	uc *biz.UserUsecase
}

func NewUserService(uc *biz.UserUsecase) *UserService {
	return &UserService{uc: uc}
}

func (s *UserService) Register(ctx context.Context, req *v1.RegisterRequest) (*v1.RegisterReply, error) {
	user := &biz.WlUser{
		UserName: req.UserName,
		Password: req.Password,
		Mobile:   req.Mobile,
		Email:    req.Email,
	}
	newUser, err := s.uc.Register(ctx, user)
	if err != nil {
		return &v1.RegisterReply{Code: 1, Message: err.Error()}, nil
	}
	return &v1.RegisterReply{Code: 0, Message: "success", User: &v1.WlUser{
		Id:           newUser.Id,
		UserName:     newUser.UserName,
		UserNickname: newUser.UserNickname,
		Department:   newUser.Department,
		Mobile:       newUser.Mobile,
		Email:        newUser.Email,
		Password:     newUser.Password,
		Gender:       newUser.Gender,
		Role:         newUser.Role,
		UserStatus:   newUser.UserStatus,
		Comment:      newUser.Comment,
	}}, nil
}

func (s *UserService) Login(ctx context.Context, req *v1.LoginRequest) (*v1.LoginReply, error) {
	user, err := s.uc.Login(ctx, req.UserName, req.Password)
	if err != nil {
		return &v1.LoginReply{Code: 1, Message: err.Error()}, nil
	}
	token, err := data.GenerateJWT(user.Id, user.UserName)
	if err != nil {
		return &v1.LoginReply{Code: 2, Message: "生成Token失败"}, nil
	}
	return &v1.LoginReply{Code: 0, Message: "success", Token: token, User: &v1.WlUser{
		Id:           user.Id,
		UserName:     user.UserName,
		UserNickname: user.UserNickname,
		Department:   user.Department,
		Mobile:       user.Mobile,
		Email:        user.Email,
		Password:     user.Password,
		Gender:       user.Gender,
		Role:         user.Role,
		UserStatus:   user.UserStatus,
		Comment:      user.Comment,
	}}, nil
}
