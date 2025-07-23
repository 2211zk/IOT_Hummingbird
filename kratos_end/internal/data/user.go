package data

import (
	"IOT_Hummingbird_back_end/internal/biz"
	"context"
	"crypto/md5"
	"encoding/hex"
	"errors"
	"time"

	"fmt"

	"github.com/dtm-labs/dtmcli"
	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
)

var jwtSecret = []byte("your_jwt_secret")

func md5Encrypt(pwd string) string {
	h := md5.New()
	h.Write([]byte(pwd))
	return hex.EncodeToString(h.Sum(nil))
}

func GenerateJWT(userId int32, userName string) (string, error) {
	claims := jwt.MapClaims{
		"user_id":   userId,
		"user_name": userName,
		"exp":       time.Now().Add(time.Hour * 24).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

type userRepo struct {
	data *Data
}

func NewUserRepo(data *Data) biz.UserRepo {
	return &userRepo{data: data}
}

func (r *userRepo) Register(ctx context.Context, user *biz.WlUser) (*biz.WlUser, error) {
	user.Password = md5Encrypt(user.Password)

	// DTM SAGA 分布式事务示例
	dtmServer := "http://localhost:36789/api/dtmsvr"
	gid := dtmcli.MustGenGid(dtmServer)

	// 业务服务的 SAGA 分支接口
	addUserURL := "http://localhost:8000/api/user/AddUser"
	compensateUserURL := "http://localhost:8000/api/user/CompensateUser"

	s := dtmcli.NewSaga(dtmServer, gid).
		Add(addUserURL, compensateUserURL, map[string]interface{}{
			"user_name": user.UserName,
			"password":  user.Password,
			"email":     user.Email,
		})

	err := s.Submit()
	if err != nil {
		return nil, fmt.Errorf("DTM SAGA 事务失败: %w", err)
	}
	return user, nil
}

func (r *userRepo) Login(ctx context.Context, userName, password string) (*biz.WlUser, error) {
	var user biz.WlUser
	if err := r.data.MySQL.Table("wl_user").Where("user_name = ?", userName).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("用户不存在")
		}
		return nil, err
	}
	if user.Password != md5Encrypt(password) {
		return nil, errors.New("密码错误")
	}
	return &user, nil
}
