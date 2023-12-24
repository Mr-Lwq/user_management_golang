package service

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

func GenerateSessionToken(userName string) (string, error) {
	jwtKey := []byte("this is a test key")
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["userId"] = userName       // 您的用户名
	claims["iat"] = time.Now().Unix() // 设置令牌的生成时间

	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func GetUserIdFromToken(tokenString string) (string, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte("this is a test key"), nil
	})
	if err != nil {
		return "", err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return "", fmt.Errorf("unable to get claims")
	}

	userName, ok := claims["userId"].(string)
	if !ok {
		return "", fmt.Errorf("userId not found in token claims")
	}

	return userName, nil
}

func GetTokenCreationTime(tokenString string) (time.Time, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte("this is a test key"), nil
	})
	if err != nil {
		return time.Time{}, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return time.Time{}, fmt.Errorf("unable to get claims")
	}

	iat, ok := claims["iat"].(float64)
	if !ok {
		return time.Time{}, fmt.Errorf("issue time (iat) not found in token claims")
	}

	return time.Unix(int64(iat), 0), nil
}
