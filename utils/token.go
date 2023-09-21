package utils

import (
	"crypto/rand"
	"github.com/dgrijalva/jwt-go"
	"time"
)

func GenerateSessionToken(userName string) (string, error) {
	jwtKey := make([]byte, 32)
	_, err := rand.Read(jwtKey)
	// 创建一个JWT令牌
	token := jwt.New(jwt.SigningMethodHS256)

	// 设置令牌的负载，可以包含用户ID或其他有关用户的信息
	claims := token.Claims.(jwt.MapClaims)
	claims["userId"] = userName                           // 您的用户ID或其他信息
	claims["exp"] = time.Now().Add(time.Hour * 48).Unix() // 设置令牌的过期时间

	// 使用密钥签署令牌
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// 检查 JWT 令牌是否可用（未过期）
func IsTokenValid(tokenString string) bool {
	// 解析 JWT 令牌
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// 指定密钥用于验证签名
		return []byte("your_secret_key_here"), nil
	})

	if err != nil {
		// 处理解析错误
		return false // 假定解析错误的令牌是过期的
	}

	// 检查令牌是否有效
	if !token.Valid {
		return false // 令牌无效，假定无效令牌是过期的
	}

	// 获取过期时间
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return false // 无法获取声明，假定过期
	}

	exp := int64(claims["exp"].(float64))
	currentTimestamp := time.Now().Unix()

	// 检查过期时间
	return currentTimestamp <= exp
}
