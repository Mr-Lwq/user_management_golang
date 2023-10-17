package utils

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

func GenerateSessionToken(userName string) (string, error) {
	jwtKey := []byte("this is a test key")
	//_, err := rand.Read(jwtKey)
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["userId"] = userName                          // 您的用户ID或其他信息
	claims["exp"] = time.Now().Add(time.Hour * 2).Unix() // 设置令牌的过期时间
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func IsTokenValid(tokenString string) (bool, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte("this is a test key"), nil
	})
	if err != nil {
		return false, err
	}
	if !token.Valid {
		return false, nil
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return false, err
	}
	exp := int64(claims["exp"].(float64))
	currentTimestamp := time.Now().Unix()
	return currentTimestamp <= exp, nil
}

func GetUserIdFromToken(tokenString string) (string, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte("this is a test key"), nil
	})
	if err != nil {
		return "", err
	}
	if !token.Valid {
		return "", fmt.Errorf("token is not valid")
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return "", fmt.Errorf("unable to get claims")
	}
	userId, ok := claims["userId"].(string)
	if !ok {
		return "", fmt.Errorf("userId not found in token claims")
	}
	return userId, nil
}
