package main

import (
    "github.com/dgrijalva/jwt-go"
    "github.com/lty5240/blue/service-user/proto"
    "log"
)

var (
    key = []byte("HBoUZ#YEkJ1q7Ni$hhePOQIa710QAp7y")
)

type Claims struct {
    User *proto.User
    jwt.StandardClaims
}

type AuthAble interface {
    Decode(token string) (*Claims, error)
    Encode(user *proto.User) (string, error)
}

type TokenService struct {
    repository Repository
}

func (*TokenService) Decode(token string) (*Claims, error) {
    tokenType, err := jwt.ParseWithClaims(string(key), &Claims{},
        func(token *jwt.Token) (interface{}, error) {
            return key, nil
        })
    if tokenType == nil {
        log.Panic("token解析结果为空")
    }
    if claims, ok := tokenType.Claims.(*Claims); ok && tokenType.Valid {
        return claims, nil
    } else {
        return nil, err
    }
}

func (service *TokenService) Encode(user *proto.User) (string, error) {
    claims := Claims{
        User: user,
        StandardClaims: jwt.StandardClaims{
            ExpiresAt: 1296000,
            Issuer:    "cn.lintyone.blue",
        },
    }
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    return token.SignedString(key)
}
