package service

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

func UseToken(token string) bool {
	tokenCacheMutex.RLock()
	tokenInfo, exists := tokenCache[token]
	tokenCacheMutex.RUnlock()

	if !exists {
		fmt.Println("Token not found.")
		return false
	}

	if tokenInfo.Expiry.Before(time.Now()) {
		fmt.Println("Token expired.")
		return false
	}

	// 惰性更新：仅当 token 即将在短时间内过期时更新
	if time.Until(tokenInfo.Expiry) < 30*time.Minute {
		newExpiry := time.Now().Add(1 * time.Hour)
		tokenCacheMutex.Lock()
		tokenCache[token] = TokenInfo{Expiry: newExpiry}
		tokenCacheMutex.Unlock()
		fmt.Println("Token expiry extended.")
	}

	return true
}

func CleanupExpiredTokens() {
	for {
		<-time.After(1 * time.Hour) // 每小时运行一次

		tokenCacheMutex.Lock()
		for token, info := range tokenCache {
			if info.Expiry.Before(time.Now()) {
				delete(tokenCache, token)
				fmt.Printf("Expired token %s has been removed.\n", token)
			}
		}
		tokenCacheMutex.Unlock()
	}
}
