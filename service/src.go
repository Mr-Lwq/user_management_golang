package service

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"strings"
	"unicode/utf8"
	"user_management_golang/dao"
	"user_management_golang/core"
	"user_management_golang/utils"
)

type Server struct {
	Db dao.ORM
}

// UserRegister
func (s *Server) UserRegister(account core.Account) (bool, error) {
	var db = s.Db
	var err error

	_, err = s.Db.Search(account)
	if err == nil {
		return false, fmt.Errorf("the user id '%s' already exists. Please change it and register it again", account.UserId)
	} else {
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
		} else {
			return false, fmt.Errorf("username and password must be a string of 8 to 16 characters")
		}
	}
}

// Login
func (s *Server) Login(account core.Account) (string, error) {
	var db = s.Db
	result, err := db.Search(account)
	if err != nil {
		return "", fmt.Errorf("username error, not found")
	} else {
		user, ok := result.(*core.Account)
		if !ok {
			return "", fmt.Errorf("expected result type: *core.Account, got: %T", result)
		} else {
			if user.Status == "online" {
				valid, _ := utils.IsTokenValid(user.SessionToken)
				if !valid {
					sessionToken, _ := utils.GenerateSessionToken(user.Username)
					user.SessionToken = sessionToken
					err = db.Update(user)
					if err != nil {
						return "", err
					}
					return user.SessionToken, nil
				}
				return "", fmt.Errorf("the account has been logged in, you need to log in again, please logout first")
			}
			err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(account.Password))
			if err != nil {
				fmt.Printf("password comparison failed: %v\n", err)
				return "", fmt.Errorf("password does not match")
			} else {
				sessionToken, _ := utils.GenerateSessionToken(user.Username)
				user.SessionToken = sessionToken
				user.Status = "online"
				err = db.Update(user)
				if err != nil {
					return "", err
				}
				return user.SessionToken, nil
			}
		}
	}
}

// Logout
func (s *Server) Logout(account core.Account) error {
	var db = s.Db
	result, _ := db.Search(account)
	user, ok := result.(*core.Account)
	if !ok {
		return fmt.Errorf("expected result type: *core.Account, got: %T", result)
	} else {
		if user.Status == "offline" {
			return fmt.Errorf("the account is offline, please log in first if you need to log out")
		} else {
			sessionToken, _ := utils.GenerateSessionToken(user.UserId)
			user.SessionToken = sessionToken
			user.Status = "offline"
			err := db.Update(user)
			if err != nil {
				return err
			}
			return nil
		}
	}
}

// LogoutByCredentials
func (s *Server) LogoutByCredentials(account core.Account) error {
	var db = s.Db
	result, err := db.Search(account)
	if err != nil {
		return fmt.Errorf("username error, not found")
	} else {
		user, ok := result.(*core.Account)
		if !ok {
			return fmt.Errorf("expected result type: *core.Account, got: %T", result)
		} else {
			err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(account.Password))
			if err != nil {
				fmt.Printf("password comparison failed: %v\n", err)
				return fmt.Errorf("password does not match")
			} else {
				if user.Status == "offline" {
					return fmt.Errorf("the account is offline, please log in first if you need to log out")
				}
				sessionToken, _ := utils.GenerateSessionToken(user.UserId)
				user.SessionToken = sessionToken
				user.Status = "offline"
				err := db.Update(user)
				if err != nil {
					return err
				}
				return nil
			}
		}
	}
}

// SearchRole
func (s *Server) SearchRole(account *core.Account) (string, error) {
	var db = s.Db
	result, err := db.Search(account)
	if err != nil {
		return "", fmt.Errorf("username error, not found")
	} else {
		user, ok := result.(*core.Account)
		if !ok {
			return "", fmt.Errorf("expected result type: *core.Account, got: %T", result)
		} else {
			*account = *user
			return strings.Join(user.Roles, ","), nil
		}
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
//func (s *Server) SearchPermission(account *core.Account) (string, error){
//	var db = s.Db
//	result, err := db.Search(account)
//	if err != nil {
//		return "", fmt.Errorf("username error, not found")
//	} else {
//		user, ok := result.(*core.Account)
//		if !ok {
//			return "", fmt.Errorf("expected result type: *core.Account, got: %T", result)
//		} else {
//			return strings.Join(, ","), nil
//		}
//	}
//}

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
func (s *Server) VerifyToken(account *core.Account, token string) error {
	var db = s.Db
	result, err := db.Search(account)
	if err != nil {
		return fmt.Errorf("server internal error")
	} else {
		user, ok := result.(*core.Account)
		if !ok {
			return fmt.Errorf("expected result type: *core.Account, got: %T", result)
		} else {
			if user.SessionToken == token {
				return nil
			} else {
				return fmt.Errorf("the token is expired")
			}
		}
	}
}
