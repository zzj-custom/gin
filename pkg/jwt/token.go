package jwt

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"github.com/gookit/goutil/jsonutil"
	"github.com/sirupsen/logrus"
	"go-api/internal/consts"
	"go-api/internal/util"
	"time"
)

type Claims struct {
	Id       int
	UserName string `json:"userName"`
	jwt.RegisteredClaims
}

type Token struct {
	UserName  string `json:"userName"`
	SecretKey string `json:"secretKey"`
	AppId     string `json:"appId"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
}

// GenerateToken 生成token
func GenerateToken(id int, userName string, secret []byte) (string, error) {
	c := Claims{
		Id:       id,
		UserName: userName,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(consts.TokenExpiredDuration)),
			Issuer:    "signer",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	return token.SignedString(secret)
}

// ParseToken 解析JWT
func ParseToken(tokenString string, secret []byte) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return secret, nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("token 不合法")
}

func SecretSalt(secret string, salt string) string {
	return util.MD5(util.MD5(secret+salt) + salt)
}

func EncodeUserInfo(token Token) (string, error) {
	tokenBytes, err := jsonutil.Encode(token)
	if err != nil {
		logrus.Infof("token生成失败1，错误：%v\n", err)
		return "", err
	}
	return string(tokenBytes), nil
}

func GetToken(id int, token Token, secret []byte) (string, error) {
	tokenStr, err := EncodeUserInfo(token)
	if err != nil {
		logrus.Infof("token生成失败2，错误：%v\n", err)
		return "", err
	}
	return GenerateToken(id, tokenStr, secret)
}

// AesCtrCrypt 加密
func AesCtrCrypt(plainText []byte, key string) ([]byte, error) {
	//1. 创建cipher.Block接口
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return nil, err
	}
	//2. 创建分组模式，在crypto/cipher包中
	iv := bytes.Repeat([]byte("1"), block.BlockSize())
	stream := cipher.NewCTR(block, iv)
	//3. 加密
	dst := make([]byte, len(plainText))
	stream.XORKeyStream(dst, plainText)

	return dst, nil
}
