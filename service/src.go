package service

import (
	"errors"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"strings"
	"sync"
	"time"
	"unicode/utf8"
	"user_management_golang/core"
	"user_management_golang/dao"
)

type Server struct {
	Db dao.ORM
}

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

// UserRegister
func (s *Server) UserRegister(account core.Account) (bool, error) {
	var db = s.Db
	var err error

	_, err = s.Db.Search(account)
	if err == nil {
		return false, fmt.Errorf("the user id '%s' already exists. Please change it and register it again", account.UserId)
	}
	//defer db.Close()
	passwordLength := utf8.RuneCountInString(account.Password)
	usernameLength := utf8.RuneCountInString(account.Username)
	if usernameLength >= 8 && usernameLength <= 16 && passwordLength >= 8 && passwordLength <= 16 {
		password := []byte(account.Password) // 用户的密码
		hashedPassword, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
		if err != nil {
			return false, err
		}
		account.Password = string(hashedPassword)
		if err = db.Insert(account); err != nil {
			return false, err
		} else {
			return true, nil
		}
	}

	return false, fmt.Errorf("username and password must be a string of 8 to 16 characters")
}

// Login
func (s *Server) Login(account core.Account) (string, error) {
	var db = s.Db
	result, err := db.Search(account)

	if err != nil {
		return "", fmt.Errorf("error searching for user: %v", err)
	}

	user, ok := result.(*core.Account)
	if !ok {
		return "", fmt.Errorf("expected result type: *core.Account, got: %T", result)
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(account.Password)); err != nil {
		return "", errors.New("authentication failed")
	}

	if user.Status != "activate" {
		return "", errors.New("account is unusable")
	}

	copCacheMutex.RLock()
	cop, exists := copCache[user.UserId]
	if exists && cop >= PCU {
		return "", fmt.Errorf("maximum concurrent online limit reached")
	}
	copCacheMutex.RUnlock()

	sessionToken, err := GenerateSessionToken(user.Username)
	if err != nil {
		return "", fmt.Errorf("error generating session token: %v", err)
	}

	copCache[user.UserId] = cop + 1

	expiryTime := time.Now().Add(2 * time.Hour).Unix()
	tokenCache[sessionToken] = TokenInfo{Expiry: time.Unix(expiryTime, 0)}

	return sessionToken, nil
}

// LogoutByToken
func (s *Server) LogoutByToken(token string) error {
	tokenCacheMutex.RLock()
	_, tokenExists := tokenCache[token]
	tokenCacheMutex.RUnlock()

	if !tokenExists {
		return fmt.Errorf("invalid token")
	}

	tokenCacheMutex.Lock()
	delete(tokenCache, token)
	tokenCacheMutex.Unlock()

	userId, _ := GetUserIdFromToken(token)
	copCache[userId] -= 1

	return nil
}

// LogoutByCredentials
func (s *Server) LogoutByCredentials(account core.Account) error {
	result, err := s.Db.Search(account)
	if err != nil {
		return fmt.Errorf("error searching for user: %v", err)
	}

	user, ok := result.(*core.Account)
	if !ok {
		return fmt.Errorf("expected result type: *core.Account, got: %T", result)
	}

	if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(account.Password)); err != nil {
		return errors.New("authentication failed")
	}

	copCacheMutex.RLock()
	cop, exists := copCache[user.UserId]
	if !exists || cop <= 0 {
		return errors.New("no device is currently online")
	}
	copCacheMutex.RUnlock()

	copCache[user.UserId] -= 1

	return nil
}

// SearchRole
func (s *Server) SearchRole(account *core.Account) (string, error) {
	var db = s.Db
	result, err := db.Search(account)
	if err != nil {
		return "", fmt.Errorf("username error, not found")
	}
	user, ok := result.(*core.Account)
	if !ok {
		return "", fmt.Errorf("expected result type: *core.Account, got: %T", result)
	} else {
		*account = *user
		return strings.Join(user.Roles, ","), nil
	}

}

// SearchGroup
func (s *Server) SearchGroup(account *core.Account) (string, error) {
	var db = s.Db
	result, err := db.Search(account)
	if err != nil {
		return "", fmt.Errorf("server internal error")
	} else {
		user, ok := result.(*core.Account)
		if !ok {
			return "", fmt.Errorf("expected result type: *core.Account, got: %T", result)
		} else {
			*account = *user
			return strings.Join(user.UserGroups, ","), nil
		}
	}
}

// SearchPermission
func (s *Server) SearchPermission(account *core.Account) (string, error) {
	var db = s.Db
	result, err := db.Search(account)
	if err != nil {
		return "", fmt.Errorf("username error, not found")
	}
	user, ok := result.(*core.Account)
	if !ok {
		return "", fmt.Errorf("expected result type: *core.Account, got: %T", result)
	} else {
		*account = *user
		return strings.Join(user.Permissions, ","), nil
	}
}

// Edit
func (s *Server) Edit(account *core.Account) error {
	var db = s.Db
	result, err := db.Search(account)
	if err != nil {
		return fmt.Errorf("server internal error")
	} else {
		user, ok := result.(*core.Account)
		if !ok {
			return fmt.Errorf("expected result type: *core.Account, got: %T", result)
		} else {
			user.Username = account.Username
			user.Email = account.Email
			user.Phone = account.Phone
			user.FullName = account.FullName
			user.ProfilePicture = account.ProfilePicture
			err = db.Update(user)
			if err != nil {
				return err
			}
			return nil
		}
	}
}

// CreateRole
func (s *Server) CreateRole(account *core.Account) {

}

// VerifyToken
func (s *Server) VerifyToken(token string) error {
	if UseToken(token) {
		return nil
	} else {
		return fmt.Errorf("the token is expired")
	}
}

// RetrieveTokenForUser
func (s *Server) RetrieveTokenForUser(account core.Account) ([]string, error) {
	var db = s.Db
	var tokens []string
	result, err := db.Search(account)

	if err != nil {
		return nil, fmt.Errorf("error searching for user: %v", err)
	}

	user, ok := result.(*core.Account)
	if !ok {
		return nil, fmt.Errorf("expected result type: *core.Account, got: %T", result)
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(account.Password)); err != nil {
		return nil, errors.New("authentication failed")
	}

	if user.Status != "activate" {
		return nil, errors.New("account is unusable")
	}

	tokenCacheMutex.Lock()
	for token, info := range tokenCache {
		if info.Expiry.Before(time.Now()) {
			delete(tokenCache, token)
		}
		userId, _ := GetUserIdFromToken(token)
		if userId == user.UserId {
			tokens = append(tokens, token)
		}
	}
	tokenCacheMutex.Unlock()

	return tokens, err
}
