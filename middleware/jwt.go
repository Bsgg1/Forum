package middleware

import (
	"encoding/json"
	"errors"
	"fmt"
	"forum/common"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"time"
)

type Claims struct {
	UserName string `json:"username"`
	UserId   uint   `json:"userid"`
	jwt.StandardClaims
}
type JWT struct {
	SigningKey []byte
}

func NewJWT() *JWT {
	return &JWT{SigningKey: []byte(common.Secret)}
}

const TokenExpireDuration = time.Hour * 24 //设置过期时间

var Secret = []byte("secret")

// GenToken 生成jwt
func (j *JWT) GenToken(username string, userid uint) (string, error) {
	c := Claims{
		UserName: username,
		UserId:   userid,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(TokenExpireDuration).Unix(),
			Issuer:    "ybs",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	return token.SignedString(Secret)
}

// ParseToken 解析token
func (j *JWT) ParseToken(t string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(t, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return j.SigningKey, nil
	})
	if err != nil {
		if v, ok := err.(*jwt.ValidationError); ok {
			if v.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, errors.New("That's not even a token")
			} else if v.Errors&jwt.ValidationErrorExpired != 0 {
				// Token is expired
				return nil, errors.New("Token is expired")
			} else if v.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, errors.New("Token not active yet")
			} else {
				return nil, errors.New("Couldn't handle this token:")
			}
		}
	}
	if token != nil {
		if claims, ok := token.Claims.(*Claims); ok && token.Valid {
			return claims, nil
		}
		return nil, errors.New("Couldn't handle this token:")

	} else {
		return nil, errors.New("Couldn't handle this token:")

	}
}

// RefreshToken 更新token
func (j *JWT) RefreshToken(t string) (string, error) {
	jwt.TimeFunc = func() time.Time {
		return time.Unix(0, 0)
	}
	token, err := jwt.ParseWithClaims(t, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return j.SigningKey, nil
	})
	if err != nil {
		return "", err
	}
	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		jwt.TimeFunc = time.Now
		claims.StandardClaims.ExpiresAt = time.Now().Add(1 * time.Hour).Unix()
		return j.GenToken(claims.UserName, claims.UserId)
	}
	return "", errors.New("Couldn't handle this token:")
}
func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.PostForm("token")
		fmt.Println("token")
		if token == "" {
			c.JSON(0, common.Message{
				Code: -1,
				Msg:  "非法访问",
			})
			return
		}
		j := NewJWT()
		claims, err := j.ParseToken(token)
		if err != nil {
			c.JSON(0, common.Message{
				Code: -1,
				Msg:  "非法访问",
			})
			return
		}

		marshal, err := json.Marshal(claims)
		if err != nil {
			c.JSON(0, common.Message{
				Code: -1,
				Msg:  err.Error(),
			})
			return
		}
		c.Set("claims", marshal)
		c.Next()
	}
}
