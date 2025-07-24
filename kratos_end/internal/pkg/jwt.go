package pkg

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
)

var JwtKey = []byte("your_secret_key") // 建议放配置文件

type CustomClaims struct {
	UserID   int32  `json:"user_id"`
	UserName string `json:"user_name"`
	jwt.RegisteredClaims
}

// 生成token
func GenerateToken(userID int32, userName string) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &CustomClaims{
		UserID:   userID,
		UserName: userName,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(JwtKey)
}

// 解析token
func ParseToken(tokenStr string) (*CustomClaims, error) {
	claims := &CustomClaims{}
	token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
		return JwtKey, nil
	})
	if err != nil || !token.Valid {
		return nil, err
	}
	return claims, nil
}
