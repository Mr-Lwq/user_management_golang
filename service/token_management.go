package service

import (
	"fmt"
	"sync"
	"time"
)

var PCU int

type TokenInfo struct {
	Expiry time.Time
}

var (
	// Token 缓存
	tokenCache = make(map[string]TokenInfo)
	// 同步锁
	tokenCacheMutex = &sync.RWMutex{}
	// COP 缓存
	copCache      = make(map[string]int)
	copCacheMutex = &sync.RWMutex{}
)

// UseToken 当使用token的时候惰性更新过期时间
func UseToken(token string) bool {
	tokenInfo, exists := tokenCache[token]

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
		newExpiry := time.Now().Add(2 * time.Hour)
		tokenCacheMutex.Lock()
		tokenCache[token] = TokenInfo{Expiry: newExpiry}
		tokenCacheMutex.Unlock()
		fmt.Println("Token expiry extended.")
	}

	return true
}

// TokenStacking token 存储入栈
func TokenStacking(token string) {
	expiryTime := time.Now().Add(2 * time.Hour).Unix()

	tokenCacheMutex.RLock()
	tokenCache[token] = TokenInfo{Expiry: time.Unix(expiryTime, 0)}
	tokenCacheMutex.RUnlock()

	userId, _ := GetUserIdFromToken(token)
	cop, _ := copCache[userId]

	copCacheMutex.RLock()
	copCache[userId] = cop + 1
	copCacheMutex.RUnlock()
}

// TokenOutOfStack token 出栈
func TokenOutOfStack(token string) {
	tokenCacheMutex.RLock()
	delete(tokenCache, token)
	tokenCacheMutex.RUnlock()

	userId, _ := GetUserIdFromToken(token)

	copCacheMutex.RLock()
	copCache[userId] -= 1
	copCacheMutex.RUnlock()

}

// CleanupExpiredTokens 定时清洗无效token
func CleanupExpiredTokens() {
	for {
		<-time.After(1 * time.Hour) // 每小时运行一次

		for token, info := range tokenCache {
			if info.Expiry.Before(time.Now()) {
				TokenOutOfStack(token)
				fmt.Printf("Expired token %s has been removed.\n", token)
			}
		}
	}
}
