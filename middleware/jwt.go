package middleware

import (
	"forum/common"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type Jwt interface {
	GetUid() string
}
type jwtService struct {
}

var JwtService = new(jwtService)

type CustomClaims struct {
	jwt.StandardClaims
}

const (
	TokenType = "bearer"
)

type TokenOutPut struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
	TokenType   string `json:"token_type"`
}

// CreateToken 生成 Token
func (jwtService *jwtService) CreateToken(GuardName string, user Jwt) (tokenData TokenOutPut, err error, token *jwt.Token) {
	token = jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		CustomClaims{
			StandardClaims: jwt.StandardClaims{
				ExpiresAt: time.Now().Unix() + common.TTL.Microseconds(),
				Id:        user.GetUid(),
				Issuer:    GuardName,
				NotBefore: time.Now().Unix() - 1000,
			},
		},
	)

	tokenStr, err := token.SignedString([]byte(common.Secret))
	if err != nil {
		panic(err)
	}
	tokenData = TokenOutPut{
		tokenStr,
		int(common.TTL),
		TokenType,
	}
	return
}
